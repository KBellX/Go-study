package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

func main() {
	num = 0

	wNum := 500000
	wg.Add(2 * wNum)

	for i := 1; i <= wNum; i++ {
		num = num + 1
		num = num + 1
		/*
			go add1()
			go add2()
		*/
	}

	// wg.Wait()

	fmt.Println(num)
}

func add1() {
	defer wg.Done()
	num = num + 1
}

func add2() {
	defer wg.Done()
	num = num + 1
}
