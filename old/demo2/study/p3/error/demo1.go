package main

import (
	"errors"
	"fmt"
)

func main() {
	result := test(1)
	// fmt.Println(result)
	fmt.Printf("error: %v\n", result)
}

func test(i int) error {
	if i == 1 {
		return errors.New("test error")
	} else {
		return nil
	}
}
