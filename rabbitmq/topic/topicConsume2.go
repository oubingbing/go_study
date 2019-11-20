package main

import "go_study/rabbitmq/mq"

func main()  {
	//接收全部消息
	rabbitmq := mq.NewRabbitMqTopic("topic","#")
	rabbitmq.ConsumeTopic()
}
