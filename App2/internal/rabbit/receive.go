package rabbit

import (
	"log"

	"github.com/TalesPalma/App2/internal/db"
	"github.com/TalesPalma/App2/internal/models"
)

func (r Rabbit) ReceiveMessage() {

	ch, err := r.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
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
			ResponseMsg(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

func ResponseMsg(msg string) {
	if response := db.Db.Create(&models.Message{Text: msg}); response.Error != nil {
		log.Fatal(response.Error)
	}
	log.Println("New message received and save database")
}
