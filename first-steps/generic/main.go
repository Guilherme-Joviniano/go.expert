package main

func Soma(m map[string]int) int {
	var total int
	for _, v := range m {
		total += v
	}
	return total
}

func main() {}
