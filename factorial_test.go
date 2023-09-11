package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	// Define some test cases with input and expected output values
	testCases := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{10, 3628800},
	}

	for _, tc := range testCases {
		// Call the Factorial function with the input value
		result := Factorial(tc.input)
		// Check if the result matches the expected output value
		if result != tc.output {
			// If not, report an error with the test case details
			t.Errorf("Factorial(%d) = %d; want %d", tc.input, result, tc.output)
		}
	}

}

//func TestPrimeFactorial(t *testing.T) {
//	// Define some test cases with input and expected output values
//	testCases := []struct {
//		input  int
//		output int
//	}{
//		{0, 1},
//		{1, 1},
//		{2, 2},
//		{3, 6},
//		{4, 24},
//		{5, 120},
//		{10, 3628800},
//	}
//
//	for _, tc := range testCases {
//		// Call the Factorial function with the input value
//		result := PrimeFactorial(tc.input)
//		// Check if the result matches the expected output value
//		if result != tc.output {
//			// If not, report an error with the test case details
//			t.Errorf("PrimeFactorial(%d) = %d; want %d", tc.input, result, tc.output)
//		}
//	}
//
//}
