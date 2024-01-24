package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

type Grid struct {
	grid    [][]string
	counter map[Coords]bool
	visited map[Beam]bool
}

func (g Grid) inBounds(c Coords) bool {
	return c.x >= 0 && c.x < len(g.grid[0]) && c.y >= 0 && c.y < len(g.grid)
}

func (g Grid) print() {
	for _, elem := range g.grid {
		fmt.Println(elem)
	}
}

// enum Direction:
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

type Coords struct {
	y int
	x int
}

func (c Coords) traverseGrid(d Direction) Coords {

	switch d {
	case North:
		return Coords{c.y - 1, c.x}
	case East:
		return Coords{c.y, c.x + 1}
	case South:
		return Coords{c.y + 1, c.x}
	case West:
		return Coords{c.y, c.x - 1}
	default:
		panic("problem with direction")
	}

}

type Beam struct {
	coords    Coords
	direction Direction
}

func (g *Grid) evaluateSpace(b Beam) {
	if _, ok := g.visited[b]; ok {
		return
	} else {
		g.visited[b] = true
		g.counter[b.coords] = true
	}

	bounceDirections := g.beamBounceDirections(b)
	for _, dir := range bounceDirections {
		newspace := b.coords.traverseGrid(dir)
		if g.inBounds(newspace) {
			g.evaluateSpace(Beam{newspace, dir})
		}
	}
}

func (g Grid) beamBounceDirections(b Beam) []Direction {
	current := g.grid[b.coords.y][b.coords.x]

	switch current {
	case ".":
		// fmt.Println("empty space")
		return []Direction{b.direction}
	case "|":
		if b.direction == East || b.direction == West {
			return []Direction{North, South}

		} else if b.direction == North || b.direction == South {
			return []Direction{b.direction}
		}
	case "-":
		if b.direction == East || b.direction == West {
			return []Direction{b.direction}

		} else if b.direction == North || b.direction == South {
			return []Direction{East, West}
		}
	case "/":
		if b.direction == North {
			return []Direction{East}
		} else if b.direction == East {
			return []Direction{North}
		} else if b.direction == South {
			return []Direction{West}
		} else if b.direction == West {
			return []Direction{South}
		}
	case `\`:
		if b.direction == North {
			return []Direction{West}
		} else if b.direction == East {
			return []Direction{South}
		} else if b.direction == South {
			return []Direction{East}
		} else if b.direction == West {
			return []Direction{North}
		}
	}
	return []Direction{b.direction}

}

func initGrid(input []string) Grid {
	visited := make(map[Beam]bool)
	count := make(map[Coords]bool)
	grid := make([][]string, len(input))
	for i, elem := range input {
		line := strings.Split(elem, "")
		grid[i] = line
	}
	return Grid{grid, count, visited}
}
func countEnergized(input []string, startbeam Beam) int {
	g := initGrid(input)
	g.evaluateSpace(startbeam)
	return len(g.counter)

}

func isSplitter(char string) bool {
	return char == "|" || char == "-"
}

func isEmptySpace(char string) bool {
	return char == "."
}

func partA(input []string) {
	g := initGrid(input)
	startCoords := Coords{0, 0}
	startSpace := Beam{startCoords, East}
	g.evaluateSpace(startSpace)
	fmt.Println("part_a:", len(g.counter))

}
func partB(input []string) {
	g := initGrid(input)
	maxenergy := 0
	ncols := len(g.grid[0]) - 1
	nrows := len(g.grid) - 1

	for i := 0; i <= ncols; i++ {
		g = initGrid(input)
		beam := Beam{Coords{nrows, i}, North}
		n := countEnergized(input, beam)
		if n > maxenergy {
			maxenergy = n
		}

		g = initGrid(input)
		beam = Beam{Coords{0, i}, South}
		n = countEnergized(input, beam)

		if n > maxenergy {
			maxenergy = n
		}
	}
	for i := 0; i <= nrows; i++ {
		g = initGrid(input)
		beam := Beam{Coords{i, ncols}, West}
		n := countEnergized(input, beam)
		if n > maxenergy {
			maxenergy = n
		}

		g = initGrid(input)
		beam = Beam{Coords{i, 0}, East} //* here
		n = countEnergized(input, beam)

		if n > maxenergy {
			maxenergy = n
		}
	}

	fmt.Println("part_b:", maxenergy)
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	// fmt.Println(splitInput)
	partA(splitInput)
	partB(splitInput)

}
