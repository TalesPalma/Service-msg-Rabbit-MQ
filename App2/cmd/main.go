package main

import (
	"github.com/TalesPalma/App2/internal/db"
	"github.com/TalesPalma/App2/internal/rabbit"
	"github.com/TalesPalma/App2/internal/web"
)

func main() {
	infinitChannel := make(chan struct{})
	db.InitDatabase()

	go web.InitWebServer()
	go messageServiceRabbit()

	<-infinitChannel
}

func messageServiceRabbit() {
	rabbit.NewRabbit().ReceiveMessage()
}
