package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitmqLogsReaderSettings struct {
	RabbitMQUrl  string
	QueueName    string
	ConsumerName string
	LogLevel     string
	ExchangeName string
}

func NewEnvSettings() rabbitmqLogsReaderSettings {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: can not load dot env: %v\n", err)
	}

	return rabbitmqLogsReaderSettings{
		RabbitMQUrl:  os.Getenv("localRabbitUrl"),
		QueueName:    os.Getenv("queueName"),
		ConsumerName: os.Getenv("consumerName"),
		LogLevel:     os.Getenv("logLevel"),
		ExchangeName: os.Getenv("exchangeName"),
	}
}

func main() {
	settings := NewEnvSettings()
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

	err = ch.QueueBind(
		settings.QueueName,
		settings.LogLevel,
		settings.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to queue bind. Error: %v", err)
	}

	messages, err := ch.Consume(
		settings.QueueName,
		settings.ConsumerName,
		true,
		false,
		false,
		false,
		nil,
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
