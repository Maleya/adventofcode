package intercardinaldirection

type Direction int

const (
    NE Direction = iota
    SE
    SW
    NW
)

func (d Direction) String() string {
    return [...]string{"NE", "SE", "SW", "NW"}[d]
}

func (d Direction) Opposite() Direction {
    // Define opposite direction for each intercardinal direction
    opposite := [...]Direction{SW, NW, NE, SE}
    return opposite[d]
}

func (d Direction) Add(steps int) Direction {
    numDirections := 4 // Number of intercardinal directions
    newDirection := int(d) + steps
    if newDirection < 0 {
        newDirection += numDirections * ((-newDirection / numDirections) + 1)
    }
    return Direction(newDirection % numDirections)
}