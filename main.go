package main

import (
	"fmt"
	"rabbitmq-demo/rbmq"
)

func main() {
	rabbitmq := rbmq.NewRabbitMQSimple("su15vte")
	for i := 0; i < 100; {
		sendMsg("Hello world")
	}
	rabbitmq.PulishSimple("Hello su15vte!")
	fmt.Println("send!")
	rabbitmq.ConsumeSimple()
}

func sendMsg(text string) {
	rabbitmq := rbmq.NewRabbitMQSimple("su15vte")
	rabbitmq.PulishSimple(text)
}
