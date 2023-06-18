package rbmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://su15vte:su15vte@localhost:5672/su15vte"

type RabbitMQ struct {
	conn      *amqp.Connection //连接
	channel   *amqp.Channel    //通道
	QueueName string           //队列名
	Exchange  string           //交换机名称
	Key       string           //bind key
	Mqurl     string           //连接信息
}

//创建一个新的RabbitMQ实例
func NewRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queuename, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//销毁RabbitMQ实例
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
	//关闭连接和通道
}

//错误处理
func (r *RabbitMQ) ErrorHandling(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.ErrorHandling(err, "failed to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PulishSimple(message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //持久化
		false, //自动删除
		false, //排他性
		false, //阻塞处理
		nil,   //额外的处理
	)
	r.ErrorHandling(err, "failed to consumesimple")
	msg, err := r.channel.Consume(
		q.Name,
		"",
		true,  //是否自动应答
		false, //是否独有
		false, //no-local 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-wait
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msg {
			log.Printf("收到一条信息消息：%s:%s", d.AppId, d.Body)
		}
	}()
	log.Printf("Waiting for messages...")
	<-forever
}
