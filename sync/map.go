package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var syncMap sync.Map
	syncMap.Store("name","区志彬")
	//mp := make(map[string]int)

	for i := 1; i <= 100; i++ {
		go func() {
			syncMap.Store("name","区志彬")
		}()
		go func() {
			syncMap.Store("name","宝连")
		}()
		go func() {
			fmt.Println(syncMap.Load("name"))
		}()
	}

	time.Sleep(time.Second*5)
}
