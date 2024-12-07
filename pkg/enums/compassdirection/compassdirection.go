package compass

type Direction int

const (
	North Direction = iota
	NE
	East
	SE
	South
	SW
	West
	NW
)

func (d Direction) String() string {
	return [...]string{"North", "NE", "East", "SE", "South", "SW", "West", "NW"}[d]
}

func (d Direction) Opposite() Direction {
	opposite := [...]Direction{South, SW, West, NW, North, NE, East, SE}
	return opposite[d]
}

func (d Direction) Add(steps int) Direction {
	numDirections := 8 // Total number of directions
	newDirection := int(d) + steps
	if newDirection < 0 {
		newDirection += numDirections * ((-newDirection / numDirections) + 1)
	}
	return Direction(newDirection % numDirections)
}
