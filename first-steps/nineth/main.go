package main

import "fmt"

// Funções Variádicas

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := func() int {
		return sum(2, 2, 2, 2, 2, 2, 2, 2, 2, 2)
	}()
	fmt.Println(total)
}
