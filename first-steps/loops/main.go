package main

func main() {
	for i := 0; i < 10; i++  {
		println(i)
	}
	
	numbers := []string{"um", "dois", "tres"}
	for k,v := range numbers {
		println(k,v)
	}
	
	initial := 0
	for initial < 10 {
		println(initial)
		initial ++
	}
	
	// Infinite loop (consume queue and others)
	for {
		println("hello, world!")
	}
	
}