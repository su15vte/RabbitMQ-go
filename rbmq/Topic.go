package rbmq

import (
	"log"

	"github.com/streadway/amqp"
)

func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.ErrorHandling(err, "failed to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.ErrorHandling(err, "failed to open a channel")
	return rabbitmq
}

//Topic模式发送信息
func (r *RabbitMQ) PublishTopic(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.ErrorHandling(err, "declare ERROR")
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

func (r *RabbitMQ) RecieveTopic() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.ErrorHandling(err, "declare ERROR")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.ErrorHandling(err, "declare a queue ERROR")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msg, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msg {
			log.Printf("接收到一条消息：%s", d.Body)
		}
	}()
	<-forever
}
