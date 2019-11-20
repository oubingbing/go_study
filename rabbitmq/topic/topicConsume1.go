package main

import "go_study/rabbitmq/mq"

func main()  {
	//接收指定匹配消息
	rabbitmq := mq.NewRabbitMqTopic("topic","routing.*.one")
	rabbitmq.ConsumeTopic()
}
