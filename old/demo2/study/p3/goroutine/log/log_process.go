package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	TypeHandleLine = 0
	TypeErrNum     = 1
)

/*
	+ 因为有多个协程同时生产(同时生产HandleLine, ErrNum) 所以通过通道(而不是直接修改全局变量)接收, 然后消费，避免不准确
	+ 为了避免生产HandleLine，ErrNum的协程阻塞，用带缓冲的通道。但消费仍是阻塞的
*/
var TypeMonitorChannel = make(chan int, 200)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	rc     chan []byte
	wc     chan *Message
	reader Reader
	writer Writer
}

// 程序运行情况
type SystemInfo struct {
	Tps          float64       `json:"tps"` // 吞吐量(单位时间处理行数)
	HandleLine   int           `json:"handle_line"`
	ErrNum       int           `json:"err_num"`
	RunTime      time.Duration `json:"run_time"`
	ReadChanLen  int           `json:"read_chan_len"`
	WriteChanLen int           `json:"write_chan_len"`
}

type Monitor struct {
	startTime time.Time
	data      SystemInfo
	tpsSli    []int
}

type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

// 开启http服务，阻塞，向外暴露程序运行状态
func (m *Monitor) start(lp *LogProcess) {
	// 消费TypeMonitorChannel数据
	// 创协程的跑，否则会阻塞
	go func() {
		for n := range TypeMonitorChannel {
			switch n {
			case TypeHandleLine:
				m.data.HandleLine += 1
			case TypeErrNum:
				m.data.ErrNum += 1
			}
		}
	}()

	// 定时器每隔5秒记录一次处理行数
	// 创协程的跑，否则会阻塞
	go func() {
		t := time.NewTicker(5 * time.Second)
		for {
			<-t.C // 阻塞，直到通过通道t.C收到
			m.tpsSli = append(m.tpsSli, m.data.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()

	// 给/monitor路由注册处理方法
	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {

		m.data.RunTime = time.Now().Sub(m.startTime)
		m.data.ReadChanLen = len(lp.rc)
		m.data.WriteChanLen = len(lp.wc)

		// 相隔处理行数相减 / 处理时间
		if len(m.tpsSli) >= 2 {
			m.data.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
		}

		ret, _ := json.MarshalIndent(m.data, "", "\t")
		io.WriteString(w, string(ret))
	})

	http.ListenAndServe(":9193", nil)
}

func (lp *LogProcess) Process() {
	// 192.168.211.1 - - [27/Feb/2019:14:26:07 +0800] http \"GET /index.php HTTP/1.1\" 200 2024 \"-\" \"cba\" - 1.001 2.02
	// 172.0.0.12 - - [27/Feb/2019:21:42:10 +0000] http "GET /baz HTTP/1.0" 200 935 "-" "KeepAliveClient" "-" - 0.721

	rule := `([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`
	// rule := `([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"([^"]+)\"\s+\"([^"]+)\"\s+([\d\.]+)\s+([\d\.]+)`
	// rule := `([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"([^"]+)\"\s+\"([^"]+)\"\s+([\d\.]+)\s+([\d\.]+)`
	rep := regexp.MustCompile(rule)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range lp.rc { // 循环从rc取数据, 不用写<-
		ret := rep.FindStringSubmatch(string(v))
		if len(ret) < 14 {
			TypeMonitorChannel <- TypeErrNum
			log.Println("FindStringSubmatch fail:", string(v))
			continue
		}

		message := &Message{}

		// 字符串时间 转为 time包的Time类型 时间
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChannel <- TypeErrNum
			log.Println("time parseInLocation fail:", ret[4], err.Error())
			continue
		}
		message.TimeLocal = t

		byteSent, _ := strconv.Atoi(ret[8])
		message.BytesSent = byteSent

		retSli := strings.Split(ret[6], " ")
		u, err := url.Parse(retSli[1])
		if err != nil {
			TypeMonitorChannel <- TypeErrNum
			log.Println("url parse fail:", err)
			continue
		}
		message.Path = u.Path
		message.Method = retSli[0]
		message.Scheme = ret[5]
		message.Status = ret[7]

		message.UpstreamTime, _ = strconv.ParseFloat(ret[12], 64)
		message.RequestTime, _ = strconv.ParseFloat(ret[13], 64)

		lp.wc <- message
	}
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	dsn string
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("os.Open err: %s\n", err.Error()))
	}

	// 移动指针到文件末尾
	f.Seek(0, 2)

	// 循环读取，往rc里塞
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			// 没新内容停止500毫秒后扫
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s\n", err.Error()))
		}

		rc <- line[:len(line)-1] // 去掉换行符
	}
}

func (w *WriteToInfluxDB) Write(wc chan *Message) {
	infSli := strings.Split(w.dsn, "@")

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     infSli[0],
		Username: infSli[1],
		Password: infSli[2],
	})
	if err != nil {
		TypeMonitorChannel <- TypeErrNum
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  infSli[3],
		Precision: infSli[4],
	})
	if err != nil {
		TypeMonitorChannel <- TypeErrNum
		log.Fatal(err)
	}

	for v := range wc {
		// fmt.Println(v)
		// Create a point and add to batch
		tags := map[string]string{
			"Path":   v.Path,
			"Method": v.Method,
			"Scheme": v.Scheme,
			"Status": v.Status,
		}
		fields := map[string]interface{}{
			"BytesSent":    v.BytesSent,
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
		}

		pt, err := client.NewPoint("log", tags, fields, v.TimeLocal)
		if err != nil {
			TypeMonitorChannel <- TypeErrNum
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := c.Write(bp); err != nil {
			TypeMonitorChannel <- TypeErrNum
			log.Fatal(err)
		}

		// Close client resources
		if err := c.Close(); err != nil {
			TypeMonitorChannel <- TypeErrNum
			log.Fatal(err)
			// 记录处理行数
		}

		// 记录处理行数
		TypeMonitorChannel <- TypeHandleLine
		fmt.Println("success")

	}
}

func main() {
	var readFile, influxDBDsn string
	flag.StringVar(&readFile, "readFile", "./access.log", "read file path")
	flag.StringVar(&influxDBDsn, "influxDsn", "http://localhost:8086@@@mydb@s", "influxdb dsn")
	flag.Parse()

	reader := &ReadFromFile{readFile}
	writer := &WriteToInfluxDB{"http://localhost:8086@@@mydb@s"}

	lp := &LogProcess{
		rc:     make(chan []byte, 200), // 为减少生产者阻塞，设置buf
		wc:     make(chan *Message, 200),
		reader: reader,
		writer: writer,
	}

	// 为提高效率
	// 一个 读取 生产者 对应 两个 处理 消费者
	// 一个处理生产者 对应 两个 写入 消费者
	go lp.reader.Read(lp.rc)
	for i := 1; i <= 2; i++ {
		go lp.Process()
	}
	// 写入模块处理最慢，多开几个协程去消费
	for i := 1; i <= 4; i++ {
		go lp.writer.Write(lp.wc)
	}

	// 监控程序运行状况
	m := &Monitor{
		startTime: time.Now(),
		data:      SystemInfo{},
	}
	m.start(lp)
	// 	time.Sleep(180e9)
}
