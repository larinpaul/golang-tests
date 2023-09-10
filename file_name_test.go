package main

import (
	"reflect"
	"testing"
)

func MyFunction(input []int) []int {
	output := make([]int, len(input))
	for i, v := range input {
		output[i] = v + 2
	}
	return output
}

func TestMyFunction(t *testing.T) {
	// Set up any necessary test data or mocks here
	input := []int{1, 2, 3}
	expectedOutput := []int{3, 4, 5}

	// Call the function you want to test
	actualOutput := MyFunction(input)

	// Assert that the output matches the expected output
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, actualOutput)
	}

}
