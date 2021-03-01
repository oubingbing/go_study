package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 50

var m sync.RWMutex

//使用渠道进行控制变量的访问
func main()  {

	chn := make(chan int,2)

	go func() {
		chn <- 1
		Deposit(100)
		fmt.Printf("first %v\n",Balance())
		<- chn
	}()

	go func() {
		chn <- 1
		Deposit(200)
		fmt.Printf("second %v\n",Balance())
		<- chn
	}()

	time.Sleep(10)
	close(chn)

	fmt.Printf("总金额：%v\n",Balance())

}

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

