package main

import (
	"fmt"

	"github.com/Guilherme-Joviniano/go.expert/packging/1/math"
)

func main() {
	// fmt.Println("Hello, World!")
	m := math.Math{
		A: 2,
		B: 1,
	}
	fmt.Println(m.Add())
}
 