package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			//fmt.Printf("counter %v \n",x)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
			//fmt.Printf("Squarer %v \n",x)
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Printf("printer %v \n",x)
	}
}