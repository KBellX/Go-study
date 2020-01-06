/*

 */

package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

// 再包一层，使外面调用更方便，对外面更透明
type LogProcess struct {
	reader Reader
	writer Writer
}

func (lp *LogProcess) work() {
	rc := make(chan string)
	wc := make(chan string)
	go lp.reader.Read(rc)
	go lp.Process(rc, wc)
	go lp.writer.Write(wc)
}

func (lp *LogProcess) Process(rc, wc chan string) {
	data := <-rc
	wc <- strings.ToUpper(data)
}

// 实现接口的一个类型
type ReadFromFile struct {
	path string // 类型里定义该类型字段, 这个字段与Reader无关，因为各实现Reader接口的类型所需的字段并非一定相同
}

type WriteToInfluxDB struct {
	dsn string
}

func (self *ReadFromFile) Read(rc chan string) {
	rc <- "message"
}

func (self *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

func main() {
	// 后面调用只需指定(若没就按接口约束自定义实现)reader，writer，即可复用Process处理
	reader := &ReadFromFile{"/tmp/1.log"}
	writer := &WriteToInfluxDB{"dwadwad"}

	lp := &LogProcess{
		reader: reader,
		writer: writer,
	}

	lp.work()

	time.Sleep(1e9)
}
