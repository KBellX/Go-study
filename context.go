package main

/*
	使用context控制协程
*/

import (
	"context"
	"errors"
	"sync"
)

func rpc(ctx context.Context) error {
	// rpc请求通过channel返回结果
	err := make(chan error, 1)
	result := make(chan string, 1)

	// 执行rpc请求
	isSuccess := true
	if isSuccess {
		result <- "成功啦"
	} else {
		err <- errors.New("请求失败")
	}

	select {
	case <-ctx.Done():
		// 别的rpc请求挂了
		return ctx.Err()
	case e := <-err:
		// 我这个rpc请求挂了
		return e
	case <-result:
		// 正常
		return nil
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := rpc(ctx)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()

}
