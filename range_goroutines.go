package main

import (
	"fmt"
	"sync"
)

func main() {

	channel := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println(1)
		channel <- 1
		wg.Done()
	}()

	go func() {
		fmt.Println(2)
		channel <- 2
		wg.Done()
	}()

	go func() {
		fmt.Println(3)
		channel <- 3
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(channel)
	}()

	for c := range channel {
		fmt.Printf("c %v \n",c)
	}
}