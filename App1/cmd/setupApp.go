package main

import (
	"github.com/TalesPalma/GolangRabbitMQ/internal/web"
)

func initWebServer(infinite chan struct{}) {
	web.InitServer()
	close(infinite)
}
