// fmt.String接口实现实例
package main

import (
	"fmt"
	"os"
)

type MyFileMode uint32

func main() {
	num := 0777

	a := os.FileMode(num)
	b := MyFileMode(num)

	fmt.Println(a)
	fmt.Println(b)
}

func (self MyFileMode) String() string {
	fmt.Println("调用了")
	if self == 1 {
		return "hello"
	} else {
		return "world"
	}
}
