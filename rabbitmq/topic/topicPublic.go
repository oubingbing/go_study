package main

import "go_study/rabbitmq/mq"

func main()  {
	rabbitmq1 := mq.NewRabbitMqTopic("topic","routing.topic.one")
	rabbitmq1.PublicTopic("routing_one")

	rabbitmq2 := mq.NewRabbitMqTopic("topic","routing_two")
	rabbitmq2.PublicTopic("routing_two")
}
