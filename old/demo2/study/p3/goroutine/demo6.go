package main

import (
	"fmt"
	"sync"
)

var data int                // int的零值为0
var memoryAccess sync.Mutex //1

func main() {
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("value is %d\n", data)
	}
	memoryAccess.Unlock()
}

/*Output
可能情况
1: 输出0
2: 不输出

加锁后，不会出现data == 0，但输出1的情况，解决了部分问题
*/
