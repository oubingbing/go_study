package main

import (
	"fmt"
	"go_study/rabbitmq/mq"
)

func main()  {
	rabbimq := mq.NewRabbitMqSample("sample")
	for i:=0; i<=10;i++  {
		rabbimq.PublicSample(fmt.Sprintf("你有新的快递:%v",i))
	}
	fmt.Printf("public：消息发送完毕")
}
