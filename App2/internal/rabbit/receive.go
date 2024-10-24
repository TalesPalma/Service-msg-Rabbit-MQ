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
		"App1Msg", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
			ResponseMsg(string(d.Body))
			// r.SendMessage(string("Testando canal:Se ta vendo isso é porque a app2 salvou no db"))
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
