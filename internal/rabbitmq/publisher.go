package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

const (
	QueueName     = "protomq"
	ConnectionURL = "amqp://guest:guest@localhost:5672/"
)

type Publisher interface {
	Publish(data interface{}) error
}

type RabbitMQPublisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbitMQPublisher(url, queueName string) (*RabbitMQPublisher, error) {
	mq := &RabbitMQPublisher{}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	mq.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	mq.channel = ch

	q, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, err
	}
	mq.queue = q

	return mq, nil
}

func (p *RabbitMQPublisher) Publish(data interface{}) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",           // Exchange
		p.queue.Name, // Routing key (queue name)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         jsonBody,
		},
	)
	return err
}
