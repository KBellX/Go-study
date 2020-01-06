/*
有缓冲的通道，生产者，消费者会堵塞吗
*/
package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 1)

var resource int

func main() {
	resource = 3

	go producer()
	go costomer()

	time.Sleep(4e9)
}

// 生产
func producer() {
	time.Sleep(1e9)
	c <- 1
	fmt.Println("produced")
	/*
		for i := resource; i > 0; i-- {
			fmt.Println("start produce")
			c <- i
		}
	*/
}

// 消费
func costomer() {
	// time.Sleep(1e9)
	fmt.Println(<-c)
	fmt.Println("costomer")

	/*
		for v := range c {
			// time.Sleep(1e9)
			// fmt.Println("start costomer")

			fmt.Println(v)
		}
	*/

}
