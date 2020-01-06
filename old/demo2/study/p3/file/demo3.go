/*
12.1
编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数

	思路：按行划分输入，统计
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var charNum, wordNum, lineNum int

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// 敲换行符，则将\n即\n前所有返回给input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		// linux下，行首开始敲S+换行，表示结束
		if input == "S\n" {
			fmt.Println("Here are the counts:\n")
			fmt.Printf("Numbers of characters is %d\n", charNum)
			fmt.Printf("Numbers of words is %d\n", wordNum)
			fmt.Printf("Numbers of lines is %d\n", lineNum)
			os.Exit(0)
		}

		Counters(input)
	}
}

// 统计
func Counters(input string) {
	charNum += len(input) - 1             // linux下减去 \n
	wordNum += len(strings.Fields(input)) // Fields将字符串按空格分隔成数组
	lineNum++
}
