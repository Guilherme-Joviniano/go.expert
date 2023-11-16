package main

import "fmt"

type ID int
type JSON []byte

var (
	f ID = 1
	j JSON
	k int32
)

func main() {
	fmt.Printf("O tipo de F é %T", f)
	fmt.Printf("O tipo de J é %T", j)
	fmt.Printf("O tipo de K é %T", k)
}