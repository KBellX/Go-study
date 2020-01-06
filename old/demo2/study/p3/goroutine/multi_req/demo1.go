// 通过 并发程序计数器 实现main协程与其他协程同步
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建5个协程，输出
	num := 5
	var wg sync.WaitGroup

	wg.Add(num) // 加

	for i := 1; i <= num; i++ {
		go p(&wg, i)
	}

	wg.Wait() // 阻塞
}

func p(wg *sync.WaitGroup, i int) {
	defer wg.Done() // 减
	fmt.Println(i)
}
