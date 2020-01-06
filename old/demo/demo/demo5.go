package main

import (
	"fmt"
)

type s struct {
	name string
	age  int
}

func main() {
	a := "name"

	s1 := &s{"kbell", 10}

	fmt.Println(s1.a)

}
