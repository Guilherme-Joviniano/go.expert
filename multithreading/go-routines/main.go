package main

import (
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		println(name, ":", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)

	go task("A", &wg)
	go task("B", &wg)

	go func(name string, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			println(name, ":", i)
			time.Sleep(1 * time.Second)
		}
		time.Sleep(5 * time.Second)
		wg.Done()

	}("C", &wg)

	wg.Wait()
}
