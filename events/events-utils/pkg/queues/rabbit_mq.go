package queues

import (
	"context"
	"errors"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrTimeoutPublishMessage = errors.New("timeout reached to the publish message")
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

func (r *RabbitMQ) PublishWithContextTimeout(
	ctx context.Context,
	channel *amqp.Channel,
	timeout uint32,
	exchange, key string,
	mandatory bool,
	imediate bool,
	message *amqp.Publishing,
) error {
	context, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(timeout))

	errChan := make(chan error)

	go func() {
		err := channel.PublishWithContext(context, exchange, key, mandatory, imediate, *message)
		errChan <- err
	}()

	defer cancel()

	select {
	case <-context.Done():
		return ErrTimeoutPublishMessage
	case err := <-errChan:
		return err
	}
}
