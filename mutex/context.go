package main

import (
	"context"
	"fmt"
	"errors"
	"time"
)

func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := false
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <- ctx.Done():
		// 其他RPC调用调用失败
		fmt.Println("被取消了")
		return ctx.Err()
	case e := <- err:
		// 本RPC调用失败，返回错误信息
		fmt.Println("调用失败")
		return e
	case <- result:
		// 本RPC调用成功，不返回错误信息
		fmt.Println("调用成功")
		return nil
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done(): //一直阻塞在这里，需要让main goruntine执行cancel关闭channel
				return // returning not to leak the goroutine
			case dst <- n:
				fmt.Println("test")
				n++
			}
		}
	}()
	return dst
}

func t()  {
	fmt.Println("被main goruntine 通知退出")
}
