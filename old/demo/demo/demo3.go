package main

// 结构体的未赋值字符串字段为空""

import (
	"fmt"
)

type s struct {
	str string
	num int
}

func main() {
	s1 := &s{
		num: 1,
	}

	if s1.str == "" {
		fmt.Println("s1.str为空")
	}
}
