package rabbitmq_logger

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQLogger struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitMQLogger(url string) (*rabbitMQLogger, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"logs_exchange",    // exchange name
		amqp.ExchangeTopic, // exchange type
		false,              // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"logging", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	return &rabbitMQLogger{
		conn: conn,
		ch:   ch,
		q:    q,
	}, nil
}

func (r *rabbitMQLogger) Info(message string) {
	r.publishMessage("INFO", message)
}

func (r *rabbitMQLogger) Warn(message string) {
	r.publishMessage("WARN", message)
}

func (r *rabbitMQLogger) Error(message string) {
	r.publishMessage("ERROR", message)
}

func (r rabbitMQLogger) formatMessage(level, message string) string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s | %s | %s", formattedTime, level, message)
}

func (r *rabbitMQLogger) publishMessage(level, body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := r.formatMessage(level, body)

	err := r.ch.PublishWithContext(ctx,
		"logs_exchange",        // exchange
		strings.ToLower(level), // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Printf("failed to publish a message. Error: %s\n", err)
	} else {
		fmt.Println(strings.ToLower(level), message)
	}
}

func (r *rabbitMQLogger) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}
