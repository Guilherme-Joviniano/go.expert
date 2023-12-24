package main

import (
	"fmt"

	"github.com/gjovs/events-utils/pkg/queues"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMQ := queues.RabbitMQ{}
	channel, err := rabbitMQ.OpenChannel("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	messages := make(chan amqp.Delivery)

	go func() {
		val := rabbitMQ.Consume(channel, &queues.ConsumerConfigs{
			Queue:        "default_queue",
			ConsumerName: "default_consumer_name",
			NoLocal:      false,
			NoAck:        false,
			Exclusive:    false,
			NoWait:       false,
			Arguments:    nil,
		})

		messages <- <-val
	}()

	defer channel.Close()

	for message := range messages {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
