package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id      uint64
	Message string
}

func main() {
	i := uint64(0)

	c1 := make(chan Message)
	c2 := make(chan Message)

	// RabbitMQ
	go func() {
		for {
			atomic.AddUint64(&i, 1)
			time.Sleep(time.Second * 4)
			msg := Message{i, "Hello from Rabbit"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddUint64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()
	// for i := 0; i < 3; i++ {
	for {
		select {
		case value := <-c2: // Kafka
			fmt.Printf("From: [Id] %d Message: %s\n", value.id, value.Message)
		case value := <-c1: // RabbitMQ
			fmt.Printf("From: [Id] %d Message: %s\n", value.id, value.Message)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
