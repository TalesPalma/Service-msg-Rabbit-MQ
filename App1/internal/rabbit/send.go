package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (r Rabbit) SendMessage(bodyMessage string) {

	ch, err := r.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_exchange", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"App1Msg", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // Fanout doesn't use routing keys
		"logs_exchange", // exchange name
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to bind a queue")

	err = ch.Publish(
		"logs_exchange", // name exchange
		"",              // fanout doesn't use routing keys
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyMessage),
		},
	)

	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", bodyMessage)

}
