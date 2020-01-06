// 通道的堵塞
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(5e9)
		v := <-ch // 接收
		fmt.Println("received", v)
	}()

	fmt.Println("sending")
	ch <- 10 // 发送
	fmt.Println("sent")

}

/* Output:
sending 10
(5 s later):
received 10
sent 10
*/
