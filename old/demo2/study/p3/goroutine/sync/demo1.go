/*
WaitGroup可理解为安全的并发计数器
Add加
Done减
Wait阻塞
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1) // 按1.2,3,4,5顺序创了5个协程，但最终输出顺序不确定，因为Println函数执行顺序不确定
	}
	wg.Wait()
}
