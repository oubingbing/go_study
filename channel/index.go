package main

import (
	"fmt"
	"sync"
)

func main()  {
	chn := make(chan int,5)
	var wg sync.WaitGroup
	wg.Add(5)

	for i:=0;i<5;i++{
		chn<-i
		wg.Done()
	}

	go func() {
		wg.Wait()
		close(chn)
	}()

	for v := range chn {
		fmt.Println(v)
	}
}
