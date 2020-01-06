/*
1.流程：读取->处理->写入， 因为步骤读取可以从文件，标准输入... , 而写入可以写到文件，各数所以
	将读取和写入写成接口，规范接口方法的参数和返回值， 确保实现接口的类型，能用上公共步骤处理

2. 接口不能定义字段，如何处理？
	只能在具体类型里定义，在这里，各Reader类型并不一定有公共字段

3. 步骤中有公共的属性，抽象成接口后，写在哪里？如rc， ReadFromFile方法和Process方法都有用到
	由外界传，即在调用处声明，传参

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

func Process(rc, wc chan string) {
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
	// var rc, wc chan string // 可理解成传引用
	rc := make(chan string) // 记得初始化
	wc := make(chan string)
	reader := &ReadFromFile{"/tmp/1.log"}
	go reader.Read(rc)

	go Process(rc, wc)

	writer := &WriteToInfluxDB{"dwadwad"}
	go writer.Write(wc)

	time.Sleep(1e9)
}
