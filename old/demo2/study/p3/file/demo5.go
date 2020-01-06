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
	"github.com/golang-collections/collections/stack"
	"strconv"
)

const Max = 100

const Q = " Or q to exit!"

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := stack.New()

	for {
		length := s.Len()
		if length <= 1 {
			fmt.Println("Please enter a number" + Q)
		} else {
			fmt.Println(length - 1)
			fmt.Println("Please enter a number or  a operator" + Q + " (" + strconv.Itoa(length-1) + " operator need to be provided)")
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

func Calc(s *stack.Stack, input string) {
	var result, tmp1, tmp2 int

	if num, err := strconv.Atoi(input); err == nil {
		s.Push(num)
	} else {
		switch input {
		case "+":
			tmp1 = s.Pop().(int)
			tmp2 = s.Pop().(int)
			result = tmp1 + tmp2
			break
		case "-":
			tmp1 = s.Pop().(int)
			tmp2 = s.Pop().(int)
			result = tmp2 - tmp1
			break
		case "*":
			tmp1 = s.Pop().(int)
			tmp2 = s.Pop().(int)
			result = tmp1 * tmp2
			break
		case "/":
			tmp1 = s.Pop().(int)
			tmp2 = s.Pop().(int)
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
