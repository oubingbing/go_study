package main

import "go_study/rabbitmq/mq"

func main()  {
	rabbitmq := mq.NewRabbitMqSub("sub")
	rabbitmq.ConsumeSub()
}
