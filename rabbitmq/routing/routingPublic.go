package main

import "go_study/rabbitmq/mq"

func main()  {
	rabbitmq1 := mq.NewRabbitMqRouting("routing","routing_one")
	rabbitmq1.PublicRouting("routing_one")

	rabbitmq2 := mq.NewRabbitMqRouting("routing","routing_two")
	rabbitmq2.PublicRouting("routing_two")
}
