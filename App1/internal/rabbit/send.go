package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (r Rabbit) SendMessage(bodyMessage string) {

	ch, err := r.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"App1Msg", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyMessage),
		},
	)

	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", bodyMessage)

}
