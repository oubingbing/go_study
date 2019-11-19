package main

import (
	"fmt"
	"go_study/rabbitmq/mq"
)

func main()  {
	rb := mq.NewRabbitMqSample("sample")
	rb.ConsumeSample()
	fmt.Printf("消息接收完毕")
}
