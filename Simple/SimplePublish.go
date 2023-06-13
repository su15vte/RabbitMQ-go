package main

import (
	"fmt"
	"rabbitmq-demo/rbmq"
)

func main() {
	rabbitmq := rbmq.NewRabbitMQSimple("su15vte")
	rabbitmq.PulishSimple("Hello su15vte!")
	fmt.Println("send!")
}
