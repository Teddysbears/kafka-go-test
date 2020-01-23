package main

import (
	"github.com/"
	"os"
)

func main() {
	if os.Args[1] == "consumer" {
		consumer.StartConsumer()
	} else if os.Args[1] == "producer" {
		producer.StartProducer()
	}
}
