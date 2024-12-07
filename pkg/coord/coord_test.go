package coords

import (
	"adventofcode/pkg/enums/compassdirection"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		start     Coordinates
		direction compassdirection.Direction
		expected  Coordinates
	}{
		{Coordinates{0, 0}, compassdirection.North, Coordinates{-1, 0}},
		{Coordinates{0, 0}, compassdirection.NE, Coordinates{-1, 1}},
		{Coordinates{0, 0}, compassdirection.East, Coordinates{0, 1}},
		{Coordinates{0, 0}, compassdirection.SE, Coordinates{1, 1}},
		{Coordinates{0, 0}, compassdirection.South, Coordinates{1, 0}},
		{Coordinates{0, 0}, compassdirection.SW, Coordinates{1, -1}},
		{Coordinates{0, 0}, compassdirection.West, Coordinates{0, -1}},
		{Coordinates{0, 0}, compassdirection.NW, Coordinates{-1, -1}},
	}

	for _, test := range tests {
		result := test.start.Move(test.direction)
		if result != test.expected {
			t.Errorf("Move(%v, %v) = %v; want %v", test.start, test.direction, result, test.expected)
		}
	}
}
