package main

import (
	"flag"
	"fmt"
)

var pFlag = flag.Bool("p", false, "RabbitMQ publisher start flag")
var cFlag = flag.Bool("c", false, "RabbitMQ consumner start flag")

func main() {
	flag.Parse()

	if *pFlag {
		fmt.Println("start pulisher service")
	}

	if *cFlag {
		fmt.Println("start consumer service")
	}
}
