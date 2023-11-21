package main

func Soma(m map[string]int) int {
	var total int
	for _, v := range m {
		total += v
	}
	return total
}


func SomaFloat(m map[string]float64) float64 {
	var total float64
	for _, v := range m {
		total += v
	}
	return total
}

type Number interface {
	~int | float64 | float32
}

type MyNumber int

func SomaGeneric[T Number] (m map[string]T) T {
	var soma T 
	for _,v := range m {
		soma += v
	}
	return soma
}

func Compare[T comparable] (a, b T) bool {
	 if a == b {
			return true
		} 
	 return false 
}

func main() {
	m := map[string]int{"Wesley": 1000, "Guilherme": 10000}
	mFloat := map[string]float64{"Wesley": 1000.40, "Guilherme": 10000.20}

	mMyNumber := map[string]MyNumber{"Wesley": 1000, "Guilherme": 10000}

	println(Soma(m))
	println(SomaFloat(mFloat))

	println(SomaGeneric(m))
	println(SomaGeneric(mFloat))
	println(SomaGeneric(mMyNumber))

	println(Compare(10, 10))
}
