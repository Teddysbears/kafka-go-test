package main

import (
	"github.com/Teddysbears/kafka-go-test/consumer"
	"github.com/Teddysbears/kafka-go-test/producer"

	"os"
)

func main() {
	if os.Args[1] == "consumer" {
		consumer.StartConsumer()
	} else if os.Args[1] == "producer" {
		producer.StartProducer()
	}
}
