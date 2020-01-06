package main

/*
	单引号使用
*/

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader *bufio.Reader
)

func main() {
	reader = bufio.NewReader(os.Stdin)
	// ReadString接收byte类型(即uint8)参数，即实际为数字，'\n' 被解析为ascii码10
	input, err := reader.ReadString('\n')
	// input, err := reader.ReadString(10)	// 与上效果一致
	if err == nil {
		fmt.Println(input)
	}
}
