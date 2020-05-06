package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main()  {
	client,err := rpc.Dial("tcp","localhost:8080")
	if err != nil {
		fmt.Println("rpc dial error")
	}

	var reply string
	client.Call("HelloService.GetInfo","你好，我需要一些资料",&reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
