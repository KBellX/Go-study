package main

import (
	"./libs"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"reflect"
)

// 助手函数，方法

/*
	将元素类型为任意的数组, 切片转换成元素类型为interface{}的数组切片
*/

func main() {
	slice1 := []int{1, 3, 5, 7, 9, 11}
	slice2 := []int{1, 3, 5, 7, 8}

	toItf := &libs.ToItf{}

	itf1 := toItf.ArrInt2Itf(slice1)
	itf2 := toItf.ArrInt2Itf(slice2)

	dif := different(itf1, itf2)
	fmt.Println(reflect.TypeOf(dif).String())

	for _, v := range dif {
		fmt.Println(reflect.TypeOf(v).String())
	}
}

func different(a []interface{}, b []interface{}) (dif []interface{}) {
	setA := mapset.NewSetFromSlice(a)
	setB := mapset.NewSetFromSlice(b)

	dif = setA.Difference(setB).ToSlice()

	return
}
