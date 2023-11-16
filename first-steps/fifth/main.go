package main

import "fmt"

func main() {
	var minesArray [3]int
	minesArray[0] = 1
	minesArray[1] = 2
	minesArray[2] = 3

	fmt.Println(len(minesArray) - 1)
	fmt.Println(minesArray[len(minesArray)-1])
	fmt.Println(minesArray[0])

	for i, v := range minesArray {
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}
}
