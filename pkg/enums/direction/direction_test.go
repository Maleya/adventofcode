package direction

import (
	"testing"
)

func TestOpposite(t *testing.T) {
	// Test cases for Opposite method
	tests := []struct {
		input    Direction
		expected Direction
	}{
		{North, South},
		{East, West},
		{South, North},
		{West, East},
	}
	for _, test := range tests {
		result := test.input.Opposite()
		if result != test.expected {
			t.Errorf("Opposite(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
