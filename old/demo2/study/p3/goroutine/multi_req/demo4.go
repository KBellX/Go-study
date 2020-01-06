package main

import (
	"fmt"
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

func main() {
	// 数据库取到3条数据，现要去别的服务请求extra字段
	books := make([]*Book, 3, 3)
	books[0] = &Book{1, "第一本书", 11, 0}
	books[1] = &Book{2, "第二本书", 12, 0}
	books[2] = &Book{3, "第三本书", 13, 0}

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

	// 输出查看结果
	for _, book := range books {
		fmt.Println(book)
	}

}

func send(ch chan *Res, param int, index int) {
	data := param * param // 别的服务的数据处理

	ch <- &Res{index, data}
}

func (self *Book) String() string {
	return fmt.Sprintf("id:%d; title:%s; param:%d; extra:%d\n", self.id, self.title, self.param, self.extra)
}
