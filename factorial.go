package main

func Factorial(n int) int {

	if n <= 0 {
		return 1
	}

	result := n

	for i := n - 1; i > 0; i-- {
		result *= i
	}

	return result
}
