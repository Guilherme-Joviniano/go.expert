package math

type Number interface {
	int | float32 | float64
}

func Plus[T Number](a, b T) T {
	return a + b
}

func Less[T Number] (a,b T) T {
	return a - b
}
