package main

import (
	"fmt"
	"time"
)

func main() {
	go a() // 创建了协程goroutine，执行a函数(同时)
	time.Sleep(1e9)
	fmt.Println("main end")
	time.Sleep(2e9) // main返回，所有协程都会销毁，所以要等2秒让所有协程操作完成
}

func a() {
	fmt.Println("a begin")
	time.Sleep(2e9)
	fmt.Println("a end")
}

/* Outout
a begin
main end
a end
*/
