package main

import (
	"fmt"
	"strconv"
	"time"
)

type Book struct {
	id    int
	title string
	param int // 请求参数
	extra int // 别的服务返回的数据
}

type Res struct {
	index int
	data  int // 返回数据
}

var books []*Book // 声明切片

func main() {
	num := 10
	getBooks(num)

	t := time.Now()

	normalReq()
	// multiReq()

	fmt.Println("run time: ", time.Since(t))

	// 输出查看结果
	for _, book := range books {
		fmt.Println(book)
	}

}

// 普通请求
func normalReq() {
	for _, book := range books {
		book.extra = dataDeal(book.param)
	}
}

// 并发
func multiReq() {
	ch := make(chan *Res) // 创建无缓冲通道（传*Res结构数据）

	// 并发请求extra
	for index, book := range books {
		go send(ch, book.param, index) // 创建协程请求
	}

	// 组装返回数据
	for i := 0; i < len(books); i++ {
		res := <-ch                       // 阻塞等待返回数据
		books[res.index].extra = res.data // 通过index正确匹配数据
	}

}

func send(ch chan *Res, param int, index int) {
	data := dataDeal(param)
	ch <- &Res{index, data}
}

// 模拟B服务的处理
func dataDeal(param int) int {
	time.Sleep(1e9)
	return param * param
}

// 格式化输出Book
func (self *Book) String() string {
	return fmt.Sprintf("id:%d; title:%s; param:%d; extra:%d\n", self.id, self.title, self.param, self.extra)
}

func getBooks(num int) {
	for i := 1; i <= num; i++ {
		book := &Book{}
		book.id = i
		book.title = "第" + strconv.Itoa(i) + "本书"
		book.param = i + 10
		books = append(books, book)
	}
}
