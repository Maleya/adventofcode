package coords

import (
	"adventofcode/pkg/enums/compass"
)

// Direction vectors for each compass direction
var directionVectors = map[compass.Direction][2]int{
	compass.North: {-1, 0},
	compass.NE:    {-1, 1},
	compass.East:  {0, 1},
	compass.SE:    {1, 1},
	compass.South: {1, 0},
	compass.SW:    {1, -1},
	compass.West:  {0, -1},
	compass.NW:    {-1, -1},
}

// todo: add a cardinalMove and a IntercardinalMove function

type Coordinates struct {
	Y int
	X int
}

// Move moves the coordinate in the given compass.Direction
func (c Coordinates) Move(direction compass.Direction) Coordinates {
	vector := directionVectors[direction]
	return Coordinates{Y: c.Y + vector[0], X: c.X + vector[1]}
}
