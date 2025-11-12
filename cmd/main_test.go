package main

import "testing"

func TestMain(t *testing.T) {}

// TestSum_Basic tests the basic functionality of the Sum function
func TestSum_Basic(t *testing.T) {
	// The function we are testing
	result := Sum(5, 3)

	// The expected outcome
	var expected byte = 8

	// Check if the result matches the expectation
	if result != expected {
		t.Errorf("Sum(5, 3) failed. Expected %d, Got %d", expected, result)
	}
}

// TestSum_TableDriven tests the Sum function using a variety of inputs
func TestSum_TableDriven(t *testing.T) {
	// 1. Define the test cases in a slice of structs
	var tests = []struct {
		name     string // A descriptive name for the test case
		a        byte   // first input value
		b        byte   // second input value
		expected byte   // expected result
	}{
		{"Positive Numbers", 5, 3, 8},
		// {"Negative Numbers", -5, -3, -8},
		{"Zero Input", 10, 0, 10},
		// {"Negative and Positive", -10, 5, -5},
		// {"Large Numbers", 1000, 2000, 3000},
	}

	// 2. Iterate over the test cases
	for _, tt := range tests {
		// 3. Use t.Run for isolated subtests (allows you to see which case failed)
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.a, tt.b)

			// 4. Assertion
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) failed. Expected %d, Got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}
