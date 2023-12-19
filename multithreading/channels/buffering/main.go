package main

// Buffering Channels
// Sending N messages to channel support
// Important details, for that works, the volumes of messages attaching into channel and the number of thread makes difference

func main() {
	// make(chan string) --> going to throws a deadLock
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
