package main

func main() {
	// empty
	forever := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		// full
		// Make die without deadlock
		forever <- struct{}{}
	}()

	// always is going to be empty
	// Can deadlock if  all  goroutines are asleep
	<-forever
}
