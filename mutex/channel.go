package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	go func(c chan int) {
		a := <- c
		fmt.Println(a)
	}(ch)

	ch <- 1

	time.Sleep(time.Second*5)

	fmt.Println("ok")
}
