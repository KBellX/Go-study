package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func main() {
	kide := mapset.NewSet()
	kide.Add("a")
	kide.Add("b")
	kide.Add("c")
	kide.Add("d")

	address := mapset.NewSet()

	address.Add("a")
	address.Add("b")
	address.Add("c")
	// address.Add("e")

	//一个并集的运算
	allClasses := kide.Union(address)
	fmt.Println(allClasses)

	//两个集合的差集
	fmt.Println(kide.Difference(address)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}
}
