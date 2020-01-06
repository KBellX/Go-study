// 死锁，main协程阻塞了，不会执行go f1(ch)
package main

import (
	"fmt"
)

// 接收
func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	ch := make(chan int)

	ch <- 10
	go f1(ch)
}
