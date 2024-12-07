package compass

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  string
	}{
		{North, "North"},
		{NE, "NE"},
		{East, "East"},
		{SE, "SE"},
		{South, "South"},
		{SW, "SW"},
		{West, "West"},
		{NW, "NW"},
	}

	for _, test := range tests {
		got := test.direction.String()
		if got != test.expected {
			t.Errorf("String() = %v, want %v", got, test.expected)
		}
	}
}

func TestOpposite(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  Direction
	}{
		{North, South},
		{NE, SW},
		{East, West},
		{SE, NW},
		{South, North},
		{SW, NE},
		{West, East},
		{NW, SE},
	}

	for _, test := range tests {
		got := test.direction.Opposite()
		if got != test.expected {
			t.Errorf("Opposite() = %v, want %v", got, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		direction Direction
		steps     int
		expected  Direction
	}{
		{North, 1, NE},
		{North, 2, East},
		{North, -1, NW},
		{West, 1, NW},
		{West, -1, SW},
		{NE, 1, East},
		{NE, -1, North},
		{NW, 1, North},
		{NW, -1, West},
		{South, 4, North}, // Full circle
		{NE, 8, NE},       // Full circle for intercardinal
	}

	for _, test := range tests {
		got := test.direction.Add(test.steps)
		if got != test.expected {
			t.Errorf("Add(%d) = %v, want %v", test.steps, got, test.expected)
		}
	}
}
