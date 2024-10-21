package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Rabbit struct {
	Conn *amqp.Connection
}

func NewRabbit() Rabbit {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return Rabbit{
		Conn: conn,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
