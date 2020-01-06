package main

// 只声明未赋值的字符串值为""空

import (
	"fmt"
	"reflect"
)

func main() {
	var s string

	if s == "" {
		fmt.Println("s为空")
	}

	// fmt.Println(s)

	// s = "aa"
	fmt.Println(reflect.TypeOf(s))

}
