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

//订阅模式
func NewRabbitMqSub(exchange string) *RabbitMq {
	return  NewRabbitMq("",exchange,"")
}

//路由模式
func NewRabbitMqRouting(exchange string,routingKey string) *RabbitMq {
	return  NewRabbitMq("",exchange,routingKey)
}

//话题模式
func NewRabbitMqTopic(exchange string,routingKey string) *RabbitMq {
	return  NewRabbitMq("",exchange,routingKey)
}

//简单模式发布消息
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

//订阅模式发布消息
func (mq *RabbitMq) PublicSub(message string)  {
	//需自行交换机，并且类型是fanou
	err := mq.channel.ExchangeDeclare(mq.Exchange,"fanout",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//发布到指定的交换机
	err = mq.channel.Publish(mq.Exchange,"",false,false,amqp.Publishing{ContentType:"text/plan",Body:[]byte(message)})
	mq.failOrError(err,"创建交换机失败")
}

//路由模式发布消息
func (mq *RabbitMq) PublicRouting(message string)  {
	//需自行交换机，并且类型是direct
	err := mq.channel.ExchangeDeclare(mq.Exchange,"direct",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//发布到指定的交换机,以及指定的key
	err = mq.channel.Publish(mq.Exchange,mq.Key,false,false,amqp.Publishing{ContentType:"text/plan",Body:[]byte(message)})
	mq.failOrError(err,"创建交换机失败")
}

//话题模式发布消息
func (mq *RabbitMq) PublicTopic(message string)  {
	//需自行交换机，并且类型是direct
	err := mq.channel.ExchangeDeclare(mq.Exchange,"topic",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//发布到指定的交换机,以及指定的key
	err = mq.channel.Publish(mq.Exchange,mq.Key,false,false,amqp.Publishing{ContentType:"text/plan",Body:[]byte(message)})
	mq.failOrError(err,"创建交换机失败")
}

//订阅模式消费消息
func (mq *RabbitMq)ConsumeSub()  {
	err := mq.channel.ExchangeDeclare(mq.Exchange,"fanout",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//队列名为空,排他性
	var q amqp.Queue
	q,err = mq.channel.QueueDeclare("", false, false, true, false, nil, )
	if err != nil{
		fmt.Printf("创建队列错误：%v\n",err.Error())
	}
	//队列绑定交换机
	err = mq.channel.QueueBind(q.Name,"",mq.Exchange,false,nil)
	mq.failOrError(err,"队列绑定交换机失败")


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

	fmt.Printf("等待消息...\n")
	<- chn
}

//路由模式消费消息
func (mq *RabbitMq)ConsumeRouting()  {
	err := mq.channel.ExchangeDeclare(mq.Exchange,"direct",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//队列名为空,排他性
	var q amqp.Queue
	q,err = mq.channel.QueueDeclare("", false, false, true, false, nil, )
	if err != nil{
		fmt.Printf("创建队列错误：%v\n",err.Error())
	}
	//队列绑定交换机以及指定的key
	err = mq.channel.QueueBind(q.Name,mq.Key,mq.Exchange,false,nil)
	mq.failOrError(err,"队列绑定交换机失败")


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

	fmt.Printf("等待消息...\n")
	<- chn
}

//话题模式消费消息
func (mq *RabbitMq)ConsumeTopic()  {
	err := mq.channel.ExchangeDeclare(mq.Exchange,"topic",true,false,false,false,nil)
	mq.failOrError(err,"创建交换机失败")

	//队列名为空,排他性
	var q amqp.Queue
	q,err = mq.channel.QueueDeclare("", false, false, true, false, nil, )
	if err != nil{
		fmt.Printf("创建队列错误：%v\n",err.Error())
	}
	//队列绑定交换机以及指定的key
	err = mq.channel.QueueBind(q.Name,mq.Key,mq.Exchange,false,nil)
	mq.failOrError(err,"队列绑定交换机失败")


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

	fmt.Printf("等待消息...\n")
	<- chn
}

//消费简单模式
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