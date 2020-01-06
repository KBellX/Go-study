/*
12.2
编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束。请使用您在练习11.3中开发的 stack 包。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
	"strconv"
)

const Max = 100

const Q = " Or q to exit!"

type stack struct {
	arr   [Max]int
	index int // 栈顶元素的索引
}

// 入栈
func (self *stack) Push(v int) {
	if self.index > Max {
		panic("Full stack")
	}
	self.arr[self.index] = v
	self.index++
}

// 出栈
func (self *stack) Pop() int {
	if self.index <= 0 {
		panic("Empty stack")
	}
	self.index--
	return self.arr[self.index]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := &stack{index: 0}

	for {
		if s.index <= 1 {
			fmt.Println("Please enter a number" + Q)
		} else {
			// fmt.Println("Please enter a number or  a operator" + Q + " (" + strconv.Itoa(s.index-1) + " operator need to be provided)")
			fmt.Println("Please enter a number or  a operator" + Q + " (" + string(s.index-1) + " operator need to be provided)")
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			panic(err)
		}
		// input = string([]byte(input)[0 : len(input)-1]) // 去掉最后的\n，string转[]byte获取第一个，再转回string
		input = input[0 : len(input)-1]

		if input == "q" {
			fmt.Println("Exit")
			os.Exit(0)
		}

		Calc(s, input)
	}
}

func Calc(s *stack, input string) {
	var result, tmp1, tmp2 int

	if num, err := strconv.Atoi(input); err == nil {
		s.Push(num)
	} else {
		switch input {
		case "+":
			tmp1 = s.Pop()
			tmp2 = s.Pop()
			result = tmp1 + tmp2
			break
		case "-":
			tmp1 = s.Pop()
			tmp2 = s.Pop()
			result = tmp2 - tmp1
			break
		case "*":
			tmp1 = s.Pop()
			tmp2 = s.Pop()
			result = tmp1 * tmp2
			break
		case "/":
			tmp1 = s.Pop()
			tmp2 = s.Pop()
			result = tmp2 / tmp1
			break
		default:
			fmt.Println("Input error")
			return
			break
		}
		s.Push(result)
		fmt.Printf("%d %s %d = %d\n", tmp1, input, tmp2, result)
	}

}
