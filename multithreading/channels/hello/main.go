package main

func main() {
	// empty
	stringChannel := make(chan string)

	go func() {
		// full
		stringChannel <- "Hello World!"
	}()

	// empty
	msg := <- stringChannel

	println(msg)
}
