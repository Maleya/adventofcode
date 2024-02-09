package cardinaldirection

import (
	"strconv"
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

func TestAdd(t *testing.T) {
	// Test cases for Opposite method
	tests := []struct {
		input    Direction
		addend   int
		expected Direction
	}{
		{North, 1, East},
		{North, 2, South},
		{North, 3, West},
		{North, 4, North},
		{West, 2, East},
		{East, 4, East},
		{East, 40, East},
		{East, 4000, East},
		{East, -4, East},
		{East, -4000, East},
	}
	for _, test := range tests {
		result := test.input.Add(test.addend)
		if result != test.expected {
			t.Errorf("%s + %s = %s; want %s", test.input, strconv.Itoa(test.addend), result, test.expected)
		}
	}
}
