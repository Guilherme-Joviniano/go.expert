package main

import (
	"fmt"
	math "github.com/Guilherme-Joviniano/go.expert/first-steps/package/math"
)

func main () {
	plus := math.Plus(1, 0.2)
	less := math.Less(1, 1)
	fmt.Printf("O Resultado é %v \n", plus)	
	fmt.Printf("O Resultado é %v \n", less)	
}