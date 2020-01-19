package kgo

import (
	"fmt"
	"net/http"
)

type ControllerRegister struct {
	// map：装url与处理方法映射
}

func NewControllerRegister() *ControllerRegister {
	c := &ControllerRegister{}
	return c
}

func (c *ControllerRegister) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("在这里根据分发请求到控制器")
}

// add
