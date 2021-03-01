package main

import (
	"os"
	"fmt"
	"runtime/trace"
	"sync"
)

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go test1(&wg)
	go test2(&wg)

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")

	wg.Wait()
}

func test1(wg *sync.WaitGroup)  {
	fmt.Println("test1")
	wg.Done()
}

func test2(wg *sync.WaitGroup)  {
	fmt.Println("test2")
	wg.Done()
}