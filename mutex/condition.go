package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	//为什么recvCond是读锁呢，因为lock.RLocker()返回的是(*rlocker)(rw)，rlocker是RWMutex类型，其有两个指针方法Lock和Unlock，
	//里面调用的分别是RLock和RUnlock，所以使用lock.RLocker返回的互斥锁，就是一个读锁
	recvCond := sync.NewCond(lock.RLocker())

	//mailbox = 2

	go func() {
		time.Sleep(time.Second*2)
		lock.Lock()
		for mailbox == 1 {
			fmt.Println("写信阻塞中...")
			sendCond.Wait()
			fmt.Println("被唤醒")
		}
		mailbox = 1
		fmt.Printf("写信成功：%v\n",mailbox)
		lock.Unlock()
		recvCond.Signal()
	}()

	go func() {
		time.Sleep(time.Millisecond)
		lock.RLock()
		for mailbox == 0 {
			fmt.Println("取信阻塞中...")
			recvCond.Wait()
			fmt.Println("被唤醒...")
		}
		fmt.Printf("取信件成功： %v\n",mailbox)
		mailbox = 0
		lock.RUnlock()
		sendCond.Signal()
	}()

	time.Sleep(time.Second*5)
	fmt.Println("结束")
}
