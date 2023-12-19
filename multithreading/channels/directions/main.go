package main

// only attach new data to the channel
// Receive-Only
func receive(name string, hello chan<- string) {
	hello <- name
}

// only read the channel data
// Send-Only
func read(data <-chan string) {
	println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}
