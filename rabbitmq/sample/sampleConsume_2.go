package main

import (
	"fmt"
	"go_study/rabbitmq/mq"
)

func main()  {
	rabbimq := mq.NewRabbitMqSample("sample")
	rabbimq.ConsumeSample()
	fmt.Printf("consume2:消息接收完毕")
}
