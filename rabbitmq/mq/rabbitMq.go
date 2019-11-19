package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://dev:bing@127.0.0.1:5672/imoocuser"

type RabbitMq struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
	Exchange string
	Key string
	MqUrl string
}

//实例化mq
func NewRabbitMq(queueName string,exchange string,key string) *RabbitMq {
	rabbitmq := RabbitMq{QueueName:queueName,Exchange:exchange,Key:key,MqUrl:MQURL}
	var err error
	rabbitmq.conn,err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOrError(err,"rabbitmq:创建连接失败")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOrError(err,"获取渠道失败")
	return &rabbitmq
}

//关闭连接
func (mq *RabbitMq) Close()  {
	mq.conn.Close()
	mq.channel.Close()
}


//错误处理函数
func (mq *RabbitMq) failOrError(err error,message string) {
	if err != nil{
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

//简单模式
func NewRabbitMqSample(queueName string) *RabbitMq {
	return  NewRabbitMq(queueName,"","")
}

func (mq *RabbitMq)PublicSample(message string)  {
	_,err := mq.channel.QueueDeclare(mq.QueueName, false, false, false, false, nil, )
	if err != nil{
		fmt.Printf("创建队列错误：%v\n",err.Error())
	}
	err = mq.channel.Publish(mq.Exchange,mq.QueueName,false,false,amqp.Publishing{ContentType:"text/plan",Body:[]byte(message)})
	if err != nil{
		fmt.Printf("发送消息失败：%v\n",err.Error())
	}
}

func (mq *RabbitMq)ConsumeSample()  {
	_,err := mq.channel.QueueDeclare(mq.QueueName, false, false, false, false, nil, )
	if err != nil{
		fmt.Printf("创建队列错误：%v\n",err.Error())
	}

	msg,err := mq.channel.Consume(mq.QueueName,"",true,false,false,false,nil)
	if err != nil{
		fmt.Println(err)
	}

	chn := make(chan bool)
	go func() {
		for m := range  msg {
			fmt.Printf("接收到的消息：%v\n",string(m.Body))
			//chn <- true
		}
	}()

	fmt.Printf("等待消息...")
	<- chn
}