package main

// 可以直接访问map["index"] 不存在的索引

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["a"] = "b"

	d := map1["c"]

	fmt.Println(d)
}
