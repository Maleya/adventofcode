package cardinaldirection

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func (d Direction) Opposite() Direction {
	// Define opposite direction for each direction
	opposite := [...]Direction{South, West, North, East}
	return opposite[d]
}
func (d Direction) Add(steps int) Direction {
	numDirections := 4 // Number of directions
	newDirection := int(d) + steps
	if newDirection < 0 {
		newDirection += numDirections * ((-newDirection / numDirections) + 1)
	}
	return Direction(newDirection % numDirections)
}
