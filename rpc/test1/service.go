package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (h *HelloService) GetInfo(request string,reply *string) error {
	*reply += request+"this is you need info"
	return nil
}

func main()  {
	err := rpc.RegisterName("HelloService",new(HelloService))
	if err != nil {
		fmt.Println(err)
	}

	listener,e := net.Listen("tcp",":8080")
	if e != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn,err := listener.Accept()
	if err != nil {
		fmt.Println("accept error")
	}

	rpc.ServeConn(conn)
}