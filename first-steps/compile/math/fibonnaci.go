package math


func Fibonacci(number int) int {
	if (number <= 1)  {
		return number
	}	

	var penultResult, lastResult, result int
	
	penultResult, lastResult = 0, 1

	for index := 2; index <= number; index++ {
		result = lastResult + penultResult
		penultResult = lastResult
		lastResult = result
	}

	return result
}