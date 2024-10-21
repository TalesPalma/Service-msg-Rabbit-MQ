package main

import "github.com/TalesPalma/GolangRabbitMQ/internal/db"

func main() {
	db.LoadDatabase()
	infinite := make(chan struct{})
	go initWebServer(infinite)
	<-infinite
}
