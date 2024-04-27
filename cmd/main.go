package main

import (
	"flag"
	"log"

	"protomq/internal/rabbitmq"
	"protomq/internal/util"
	pb "protomq/proto"
)

var pFlag = flag.Bool("p", false, "RabbitMQ publisher start flag")
var cFlag = flag.Bool("c", false, "RabbitMQ consumner start flag")

func main() {
	flag.Parse()

	if *pFlag {
		defer func() {
			log.Println("message sent successfully")
		}()

		p, err := rabbitmq.NewRabbitMQPublisher(
			rabbitmq.ConnectionURL,
			rabbitmq.QueueName,
		)
		if err != nil {
			log.Println(util.ErrCreatePublisher, err)
		}

		message := &pb.MyMessage{
			Id:      "test Id",
			Content: "test Content",
		}

		p.Publish(message)
	}

	if *cFlag {
		c, err := rabbitmq.NewRabbitMQConsumer(
			rabbitmq.ConnectionURL,
			rabbitmq.QueueName,
		)
		if err != nil {
			log.Println(util.ErrCreateConsumer, err)
		}

		c.Consume()
		select {}
	}
}
