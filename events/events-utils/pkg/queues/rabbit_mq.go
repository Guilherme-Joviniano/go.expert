package queues

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct{}

type ConsumerConfigs struct {
	Queue        string
	ConsumerName string
	NoLocal      bool
	NoAck        bool
	Exclusive    bool
	NoWait       bool
	Arguments    amqp.Table
}

func (r *RabbitMQ) OpenChannel(connectionString string) (*amqp.Channel, error) {
	connection, err := amqp.Dial(connectionString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	channel, err := connection.Channel()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return channel, nil
}

func (r *RabbitMQ) Consume(channel *amqp.Channel, configs *ConsumerConfigs) <-chan amqp.Delivery {
	messages, err := channel.Consume(
		configs.Queue,
		configs.ConsumerName,
		configs.NoAck,
		configs.Exclusive,
		configs.NoLocal,
		configs.NoWait,
		configs.Arguments,
	)

	if err != nil {
		log.Fatalf("Consumer [%s] fails to start, error: %s", configs.ConsumerName, err.Error())
		panic(err)
	}

	return messages
}
