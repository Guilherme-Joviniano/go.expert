package main

import (
	"fmt"
	"compile-project/math"
)

func main() {
	fibonacci := math.Fibonacci(1000)
	println(fibonacci)
}

func init() {
	fmt.Println("Iniciando o sistema...")
}