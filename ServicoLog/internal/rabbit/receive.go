package rabbit

import "log"

func (r Rabbit) ReceiveMessage() {

	ch, err := r.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Mensagens",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
