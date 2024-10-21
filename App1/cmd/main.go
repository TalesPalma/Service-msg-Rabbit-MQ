package main

import (
	"github.com/TalesPalma/GolangRabbitMQ/internal/rabbit"
)

func main() {
	infinite := make(chan struct{})
	go initWebServer(infinite)
	go rabbit.NewRabbit().ReceiveMessage()
	<-infinite
}
