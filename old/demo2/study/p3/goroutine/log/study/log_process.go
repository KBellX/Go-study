/*
日志监控， 读取，处理，写入日志

协程方式封装
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	rc   chan string
	wc   chan string
	path string
	dsn  string
}

func main() {
	lp := &LogProcess{
		rc:   make(chan string),
		wc:   make(chan string),
		path: "/tmp/access.log",
		dsn:  "username=abc...",
	}

	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInFluxDB()

	time.Sleep(1e9)
}

func (l *LogProcess) ReadFromFile() {
	// 读取模块
	l.rc <- "message"
}

func (l *LogProcess) Process() {
	// 处理模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func (l *LogProcess) WriteToInFluxDB() {
	// 写入模块
	fmt.Println(<-l.wc)
}
