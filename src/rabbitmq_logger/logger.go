package rabbitmq_logger

import (
	"context"
	"fmt"
	"log"
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

func (r *rabbitMQLogger) publishMessage(level, body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("%s | %s | %s", formattedTime, level, body)

	err := r.ch.PublishWithContext(ctx,
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Printf("failed to publish a message. Error: %s\n", err)
	}
}

func (r *rabbitMQLogger) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}
