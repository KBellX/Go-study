package main

import (
	"fmt"
)

var data int // int的零值为0

func main() {
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("value is %d\n", data)
	}
}

/*Output
可能情况
1: 输出0
2: 输出1
3: 不输出
*/
