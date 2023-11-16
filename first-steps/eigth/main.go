package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sumWithErros(25, 26)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, bool) {
	value := a + b
	if value >= 50 {
		return value, true
	}
	return value, false
}

func sumWithErros(a, b int) (int, error) {
	value := a + b
	if value >= 50 {
		return value, errors.New("[MATH ERROR]: A soma é maior que 50")
	}
	return value, nil
}
