package main

import (
	"sync"
	"time"
)

var mt sync.Mutex

func main()  {
	for i := 0; i <= 10; i++ {
		go sum(i)
	}
	time.Sleep(5)
}

func sum(count int) int {
	mt.Lock()
	count++
	mt.Unlock()
	return count
}