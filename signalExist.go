package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	go func() {
		fmt.Println("go run")
	}()

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
