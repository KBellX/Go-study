/*
快排
小到大
*/

package main

import (
	"fmt"
)

var arr = [...]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

func main() {
	fmt.Println(arr)

	quickSort()

	fmt.Println(arr)

}

func quickSort() {
	storeIndex := 0
	length := len(arr)
	pivot := arr[length-1]

	for i := 0; i < length; i++ {
		if arr[i] <= pivot {
			arr[i], arr[storeIndex] = swap(arr[i], arr[storeIndex])
			storeIndex++
		}
	}

	arr[storeIndex], arr[length-1] = swap(arr[storeIndex], pivot)
}

// 交换
func swap(i, j int) (int, int) {
	return j, i
}
