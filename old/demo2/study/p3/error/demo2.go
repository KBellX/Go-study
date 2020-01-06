package main

import "fmt"

func main() {
	defer test2()
	test1()
}

func test1() {
	defer test2()
	panic("test1")
}

func test2() {
	fmt.Println("test2")
}
