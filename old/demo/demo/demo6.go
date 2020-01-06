package main

import (
	"fmt"
)

func main() {
	// var m map[string]string	// 直接赋值报错，此时m为nil
	m := make(map[string]string)

	m["a"] = "n"

	fmt.Println(m)
}
