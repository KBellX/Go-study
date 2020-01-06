// 非协程方式 优化成接口
package main

import (
	"fmt"
	"strings"
)

type Reader interface {
	Read() string
}

type Writer interface {
	Write(wd string)
}

func Process(rd string) string {
	return strings.ToUpper(rd)
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	dsh string
}

func (r *ReadFromFile) Read() string {
	return "message"
}

func (w *WriteToInfluxDB) Write(wd string) {
	fmt.Println(wd)
}

func main() {
	reader := &ReadFromFile{"/tmp/1.log"}
	rd := reader.Read()
	wd := Process(rd)
	writer := &WriteToInfluxDB{"username=abc&password=111..."}
	writer.Write(wd)
}
