package main

import (
	"flag"
	"fmt"
	"log"

	"protomq/rabbitmq"
)

var pFlag = flag.Bool("p", false, "RabbitMQ publisher start flag")
var cFlag = flag.Bool("c", false, "RabbitMQ consumner start flag")

func main() {
	flag.Parse()

	if *pFlag {
		defer func() {
			fmt.Println("message sent successfully")
		}()

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

		p.Publish(message)
	}

	if *cFlag {
		fmt.Println("start consumer service")
	}
}
