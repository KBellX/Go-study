/*
自定义包规范地返回error
*/

package test

import (
	"fmt"
)

/*
导出函数，
捕获(恢复)所有异常
尽可能返回结果和error
*/
func TestMain(i int) (s string, err error) {
	// 捕获,处理异常
	defer func() {
		if r := recover(); r != nil {
			s = "fail"
			err = fmt.Errorf("pkg: %v", r)
		}
	}()

	s = level_1(i)
	err = nil
	return
}

func level_1(i int) string {
	return level_2(i)
}

// 内部多层次调用的函数，若不想一层层返回，可以用panic
func level_2(i int) string {
	if i == 1 {
		panic("error1")
	} else if i == 2 {
		panic("error2")
	} else {
		return "success"
	}
}
