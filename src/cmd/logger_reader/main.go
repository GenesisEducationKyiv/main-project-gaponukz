package main

import (
	"btcapp/src/settings"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	settings := settings.NewDotEnvSettings().Load()
	conn, err := amqp.Dial(settings.RabbitMQUrl)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %v", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close()
	}()

	q, err := ch.QueueDeclare(
		"logging", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %v", err)
	}

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %v", err)
	}

	var forever chan struct{}

	go func() {
		for message := range messages {
			fmt.Println(string(message.Body))
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
