package main

import "github.com/TalesPalma/serviceLog/internal/rabbit"

func main() {
	rabbit := rabbit.NewRabbit()
	rabbit.ReceiveMessage()
}
