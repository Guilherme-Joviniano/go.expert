package main

type ID int
type JSON []byte

var (
	f ID = 1
	j JSON 
)

func main() {
	println(f)
	println(j)
}
