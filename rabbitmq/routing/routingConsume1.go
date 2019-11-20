package main

import "go_study/rabbitmq/mq"

func main()  {
	rabbitmq := mq.NewRabbitMqRouting("topic","routing_one")
	rabbitmq.ConsumeRouting()
}
