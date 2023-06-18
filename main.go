package main

import (
	"rabbitmq-demo/rbmq"
	"strconv"
	"time"
)

func main() {
	go TopicRecieve()
	TopicSend()
}

// func PublishRecieve() {
// 	rabbitmq := rbmq.NewRabbitMQPublish("publishrbmq")
// 	rabbitmq.RecieveSub()
// }

// func PublishSend() {
// 	rabbitmq := rbmq.NewRabbitMQPublish("publishrbmq")
// 	for i := 0; i < 100; i++ {
// 		rabbitmq.PublishPub("Hello World   " +
// 			strconv.Itoa(i))
// 		fmt.Println("Hello World   " +
// 			strconv.Itoa(i))
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func RoutingSend() {
// 	mq1 := rbmq.NewRabbitMQRouting("su15vte", "mq_1")
// 	mq2 := rbmq.NewRabbitMQRouting("su15vte", "mq_2")
// 	for i := 0; i <= 100; i++ {
// 		mq1.PublishRouting("Hello mq1!!!" + strconv.Itoa(i))
// 		mq2.PublishRouting("Hello mq2!!!" + strconv.Itoa(i))
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func RoutingRecieve() {
// 	mq1 := rbmq.NewRabbitMQRouting("su15vte", "mq_1")
// 	mq2 := rbmq.NewRabbitMQRouting("su15vte", "mq_2")
// 	go mq1.RecieveRouting()
// 	go mq2.RecieveRouting()
// }

func TopicSend() {
	mq1 := rbmq.NewRabbitMQTopic("topicmq", "su15vte.topic.one")
	mq2 := rbmq.NewRabbitMQTopic("topicmq", "su15vte.topic.two")
	for i := 0; i <= 100; i++ {
		mq1.PublishTopic("Hello mq1!!!" + strconv.Itoa(i))
		mq2.PublishTopic("Hello mq2!!!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

func TopicRecieve() {
	mq1 := rbmq.NewRabbitMQTopic("topicmq", "#")
	mq2 := rbmq.NewRabbitMQTopic("topicmq", "su.*")
	mq1.RecieveTopic()
	mq2.RecieveTopic()
}
