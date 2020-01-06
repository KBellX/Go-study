/*
	自定义包 处理异常的规范
*/
package main

import (
	"./test"
	"fmt"
)

func main() {
	i := 2
	result, err := test.TestMain(i)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
