package main

import (
	"fmt"
	"time"
)

func Worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received value %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorkers := 1000000

	for i := 0; i < qtdWorkers; i++ {
		go Worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
