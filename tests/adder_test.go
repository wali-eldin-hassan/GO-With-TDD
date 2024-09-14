package tests

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
