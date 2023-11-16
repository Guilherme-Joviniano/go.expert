package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
	// Capacidade inicial se mantêm!
	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v \n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])
	// Como aumentar a Capacidade de um slice!
	// Não necessariamente você aumenta a capacidade do slice, ele aumenta a capacidade da array
	// Ele dobra o valor do slice, para aguentar o aumento de capacidade redimensionando x2 o tamanho da "slice"
	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}

