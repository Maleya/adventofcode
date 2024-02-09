package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// priority queue
// L2 distance to goal early stopping
// min full run distance early stopping.
// track states from grid
// from a spot, findlegaldirections() and then spawn children

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

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

type crucible struct {
	y, x           int
	heatloss       int
	lastthreemoves []Direction
	allmoves       []Direction
}

func (c crucible) String() string {
	return fmt.Sprintf("location (%v, %v) heatloss %v history %v", c.y, c.x, c.heatloss, c.lastthreemoves)
}

func (c *crucible) SpawnNewCrucibleInDirection(dir Direction, g Grid) crucible {
	var y, x int
	var newDirection, allmoveshist []Direction

	switch dir {
	case North:
		y = c.y - 1
		x = c.x
	case East:
		x = c.x + 1
		y = c.y
	case South:
		y = c.y + 1
		x = c.x
	case West:
		x = c.x - 1
		y = c.y
	}
	if !g.isonGrid(y, x) {
		panic("not on grid")
	}

	newDirection = append(newDirection, c.lastthreemoves...)
	newDirection = append(newDirection, dir)
	if len(newDirection) > 3 {
		newDirection = newDirection[1:]
	}

	// temp:
	// allmoveshist = append(allmoveshist, c.lastthreemoves...)
	allmoveshist = append(allmoveshist, dir)

	return crucible{
		y:              y,
		x:              x,
		heatloss:       c.heatloss + g.grid[y][x],
		lastthreemoves: newDirection,
		allmoves:       allmoveshist,
	}

}

func (c *crucible) FindLegalDirections(g Grid) []Direction {
	var legalDirections []Direction
	options := []Direction{North, East, South, West}

	for _, dir := range options {
		// dont move off grid
		if !g.DirectionOnGrid(c.y, c.x, dir) {
			continue
		}
		// no more than 3 in a row
		if len(c.lastthreemoves) == 3 {
			if c.lastthreemoves[0] == dir && c.lastthreemoves[1] == dir && c.lastthreemoves[2] == dir {
				continue
			}
		}

		// dont move backwards
		if len(c.lastthreemoves) > 0 {
			lastmove := c.lastthreemoves[len(c.lastthreemoves)-1]
			if lastmove == North && dir == South {
				continue
			}
			if lastmove == East && dir == West {
				continue
			}
			if lastmove == South && dir == North {
				continue
			}
			if lastmove == West && dir == East {
				continue
			}
		}

		legalDirections = append(legalDirections, dir)
	}
	return legalDirections

}

type Grid struct {
	grid          [][]int
	visitedStates map[string]bool
}

func (g Grid) isonGrid(y int, x int) bool {
	return y >= 0 && y < len(g.grid) && x >= 0 && x < len(g.grid[0])
}

func (g Grid) DirectionOnGrid(y int, x int, dir Direction) bool {
	var y_new, x_new int

	switch dir {
	case North:
		y_new = y - 1
		x_new = x
	case East:
		x_new = x + 1
		y_new = y
	case South:
		y_new = y + 1
		x_new = x
	case West:
		x_new = x - 1
		y_new = y
	}
	return g.isonGrid(y_new, x_new)

}

func (g Grid) String() string {
	var s string
	for _, row := range g.grid {
		s += fmt.Sprintf("%v\n", row)
	}
	return s
}
func (g Grid) inGoal(y, x int) bool {
	return y == len(g.grid)-1 && x == len(g.grid[0])-1
}

func parseInput(input []string) Grid {
	grid := make([][]int, len(input))
	visitedStates := make(map[string]bool)

	for i, row := range input {
		grid[i] = make([]int, len(row))
		for j, char := range row {
			grid[i][j], _ = strconv.Atoi(string(char))
		}
	}
	return Grid{
		grid:          grid,
		visitedStates: visitedStates,
	}
}

func partA(input []string) {
	g := parseInput(input)
	fmt.Println(g)
	leastHeatloss := math.MaxInt64
	cache := make(map[string]int)

	c := crucible{y: 0, x: 0, heatloss: 0, lastthreemoves: []Direction{}}
	pq := []crucible{c}

	for len(pq) > 0 {
		c := pq[0]
		// fmt.Println(c)
		pq = pq[1:]

		if c.heatloss > leastHeatloss {
			continue
		}
		if g.inGoal(c.y, c.x) {
			if c.heatloss < leastHeatloss {
				fmt.Println("GOAL", c)
				fmt.Println("allmoves", c.allmoves) // rewrite
				leastHeatloss = c.heatloss
			}
			continue
		}
		// early stopping if you are suboptimal

		enc := fmt.Sprintf("(%v, %v),%v", c.y, c.x, c.lastthreemoves)
		if _, ok := cache[enc]; ok {
			continue
		} else {
			cache[enc] = c.heatloss
		}

		// make sure we want to spawn more?
		newDirecitons := c.FindLegalDirections(g)
		if len(newDirecitons) > 0 {
			for _, dir := range newDirecitons {
				NewCrucible := c.SpawnNewCrucibleInDirection(dir, g)
				pq = append(pq, NewCrucible)
			}
		}
		sort.SliceStable(pq, func(i, j int) bool {
			return pq[i].heatloss < pq[j].heatloss
		})
	}


}
func partB(input []string) {
	fmt.Println("part_b:")
}

func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

	//todo:
	// - track all moves correctly
	// track last 10 moves
	// make ultracurible struct
	// remove "all moves"

}
