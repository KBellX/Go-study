package main

/*
	使用pool包临时缓存对象，减少内存分配
*/

import (
	"fmt"
	"sync"
)

var p *sync.Pool
var end chan bool

// 工具类
type Util struct {
}

func (u *Util) Work() {
	fmt.Println("I am working")
}

func main() {
	end = make(chan bool)

	// 初始化pool,设置空池时的返回值
	p = &sync.Pool{
		New: func() interface{} {
			return Util{}
		},
	}

	go letUtilWork()
	go letUtilWork()

	for i := 2; i > 0; i-- {
		<-end
	}

}

func letUtilWork() {
	defer func() {
		end <- true
	}()
	// 不用pool: 直接实例化，等GC回收
	// u := Util{}

	// 使用pool：
	// 从pool里拿出来用
	u := p.Get().(Util)
	// 用完返回去
	defer p.Put(u)
	u.Work()
}
