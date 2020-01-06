// 通过通道实现同步
package main

import (
	"fmt"
)

func main() {
	num := 5
	ch := make(chan int) // 创建阻塞无缓冲的通道(传输数据类型为int)

	for i := 1; i <= num; i++ {
		// 创建协程执行send方法
		go send(ch, i)
	}

	for i := 1; i <= num; i++ {
		fmt.Println(<-ch) // 阻塞，直到接收到。通过通道接收处理结果
	}

}

func send(ch chan int, i int) {
	result := i * i // 处理
	ch <- result    // 将处理结果通过通道发送给main协程
}
