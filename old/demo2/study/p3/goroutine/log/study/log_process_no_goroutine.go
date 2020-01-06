// 非协程方式 封装
package main

import (
	"fmt"
	"strings"
)

type LogProcess struct {
	rd   string // 从文件读取到的数据
	wd   string // 经处理写入的数据, 其实rd和wd可以用同一个字段，因为同步~
	path string
	dsh  string
}

func (l *LogProcess) ReadFromFile() {
	l.rd = "message"
}

func (l *LogProcess) Process() {
	l.wd = strings.ToUpper(l.rd)
}

func (l *LogProcess) WriteToInfluxDB() {
	fmt.Println(l.wd)
}

func main() {
	lp := &LogProcess{
		path: "/tmp/1.log",
		dsh:  "dadwad",
	}
	lp.ReadFromFile()
	lp.Process()
	lp.WriteToInfluxDB()
}
