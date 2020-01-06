// 带缓冲通道, 实现异步
package main

import (
	"fmt"
	"time"
)

// 接收
func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	ch := make(chan int, 1) // 1buf

	ch <- 10
	go f1(ch)

	time.Sleep(2e9) // 为显示效果
}
