package main

import (
	"flag"
	"log"

	"protomq/rabbitmq"
)

var pFlag = flag.Bool("p", false, "RabbitMQ publisher start flag")
var cFlag = flag.Bool("c", false, "RabbitMQ consumner start flag")

func main() {
	flag.Parse()

	if *pFlag {
		p, err := rabbitmq.NewRabbitMQPublisher(
			rabbitmq.ConnectionURL,
			rabbitmq.QueueName,
		)
		if err != nil {
			log.Println("error create rabbitmq publisher ", err)
		}

		message := map[string]interface{}{
			"message": "test message",
		}

		go p.Publish(message)
	}

	if *cFlag {
		c, _ := rabbitmq.NewRabbitMQConsumer(
			rabbitmq.ConnectionURL,
			rabbitmq.QueueName,
		)

		go c.Consume()
	}

	select {}
}
