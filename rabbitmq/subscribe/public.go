package main

import (
	"fmt"
	"go_study/rabbitmq/mq"
)

func main()  {
	rabbitmq := mq.NewRabbitMqSub("sub")
	rabbitmq.PublicSample("我发布消息啦")
	fmt.Println("消息发送完毕")
}