package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (r Rabbit) SendMessage(bodyMessage string) {

	ch, err := r.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeDeclare(ch, "logs_exchange")
	queueDeclares(ch, "App1Msg", "Logs")

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

func queueDeclares(ch *amqp.Channel, queueNames ...string) {

	args := map[string]interface{}{
		"x-dead-letter-exchange": "logs_exchange",
	}

	for _, nameQueue := range queueNames {
		q, err := ch.QueueDeclare(
			nameQueue, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			args,      // arguments
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

	}
}

func exchangeDeclare(ch *amqp.Channel, nameExchange ...string) {
	for _, name := range nameExchange {
		err := ch.ExchangeDeclare(
			name,     // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		failOnError(err, "Failed to declare an exchange")
	}
}
