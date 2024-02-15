package main

import (
	direction "adventofcode/pkg/enums/cardinaldirection"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// L2 distance to goal early stopping
// min full run distance early stopping.

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

type crucible struct {
	y, x           int
	heatloss       int
	lastthreemoves []direction.Direction
}

func (c crucible) String() string {
	return fmt.Sprintf("location (%v, %v) heatloss %v, history %v", c.y, c.x, c.heatloss, c.lastthreemoves)
}

func (c *crucible) SpawnNewCrucibleInDirection(dir direction.Direction, g Grid) crucible {
	var newDirection []direction.Direction

	y, x := g.displaceByDirection(c.y, c.x, dir)
	if !g.isonGrid(y, x) {
		panic("not on grid")
	}

	newDirection = append(newDirection, c.lastthreemoves...)
	newDirection = append(newDirection, dir)
	if len(newDirection) > 3 {
		newDirection = newDirection[1:]
	}

	return crucible{
		y:              y,
		x:              x,
		heatloss:       c.heatloss + g.grid[y][x],
		lastthreemoves: newDirection,
	}

}

func (c *crucible) FindLegalDirections(g Grid) []direction.Direction {
	var legalDirections []direction.Direction
	options := []direction.Direction{direction.North, direction.East, direction.South, direction.West}

	for _, dir := range options {
		// dont move off grid
		if !g.directionOnGrid(c.y, c.x, dir) {
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
			if lastmove.Opposite() == dir {
				continue
			}
		}

		legalDirections = append(legalDirections, dir)
	}
	return legalDirections

}

type ultracrucible struct {
	y, x         int
	heatloss     int
	lastTenMoves []direction.Direction
	allmoves     []direction.Direction
}

func (c ultracrucible) String() string {
	return fmt.Sprintf("location (%v, %v) heatloss %v, history %v", c.y, c.x, c.heatloss, c.lastTenMoves)
}

func (c ultracrucible) lastNmovesEqualto(n int, dir direction.Direction) bool {

	noElements := len(c.lastTenMoves)
	// lowerbound := noElements - n
	// if noElements-n < 0 {
	// 	lowerbound = 0
	// }

	for i := noElements - 1; i >= noElements-n; i-- {
		if c.lastTenMoves[i] != dir {
			return false
		}
	}
	return true
}

func (c ultracrucible) lastNmovesEqual(n int) bool {

	noElements := len(c.lastTenMoves)
	dir := c.lastTenMoves[noElements-1]

	for i := noElements - 1; i >= noElements-n; i-- {
		if c.lastTenMoves[i] != dir {
			return false
		}
	}
	return true
}

func (c *ultracrucible) SpawnNewCrucibleInDirection(dir direction.Direction, g Grid) ultracrucible {
	var newDirection []direction.Direction
	var allmoves []direction.Direction

	y, x := g.displaceByDirection(c.y, c.x, dir)
	if !g.isonGrid(y, x) {
		panic("not on grid")
	}

	newDirection = append(newDirection, c.lastTenMoves...)
	newDirection = append(newDirection, dir)

	allmoves = append(allmoves, c.allmoves...)
	allmoves = append(allmoves, dir)
	if len(newDirection) > 10 {
		newDirection = newDirection[1:]
	}

	return ultracrucible{
		y:            y,
		x:            x,
		heatloss:     c.heatloss + g.grid[y][x],
		lastTenMoves: newDirection,
		allmoves:     allmoves,
	}
}

func (cru *ultracrucible) FindLegalDirections(g Grid) []direction.Direction {
	var legalDirections []direction.Direction
	options := []direction.Direction{direction.North, direction.East, direction.South, direction.West}

	for _, dir := range options {
		// dont move off grid
		if !g.directionOnGrid(cru.y, cru.x, dir) {
			continue
		}

		// maximum 10 blocks in one dir
		if len(cru.lastTenMoves) == 10 {
			if cru.lastNmovesEqualto(10, dir) {
				continue
			}
		}
		// minimum 4 blocks forward in same direction
		moveshistory := len(cru.lastTenMoves)
		if moveshistory > 4 {
			moveshistory = 4
		}
		if len(cru.lastTenMoves) > 0 {
			if !cru.lastNmovesEqual(moveshistory) && dir != cru.lastTenMoves[len(cru.lastTenMoves)-1] {
				continue
			}
		}

		// dont move backwards
		if len(cru.lastTenMoves) > 0 {
			lastmove := cru.lastTenMoves[len(cru.lastTenMoves)-1]
			if lastmove.Opposite() == dir {
				continue
			}
		}

		legalDirections = append(legalDirections, dir)
	}
	// fmt.Println(legalDirections)
	return legalDirections

}

type Grid struct {
	grid          [][]int
	visitedStates map[string]bool
}

func (g Grid) isonGrid(y int, x int) bool {
	return y >= 0 && y < len(g.grid) && x >= 0 && x < len(g.grid[0])
}

func (g Grid) displaceByDirection(y int, x int, dir direction.Direction) (int, int) {
	var y_new, x_new int

	switch dir {
	case direction.North:
		y_new = y - 1
		x_new = x
	case direction.East:
		x_new = x + 1
		y_new = y
	case direction.South:
		y_new = y + 1
		x_new = x
	case direction.West:
		x_new = x - 1
		y_new = y
	}
	return y_new, x_new
}

func (g Grid) directionOnGrid(y int, x int, dir direction.Direction) bool {
	y_new, x_new := g.displaceByDirection(y, x, dir)
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
	grid := parseInput(input)
	fmt.Println(grid)
	leastHeatloss := math.MaxInt64
	cache := make(map[string]int)

	cru := crucible{y: 0, x: 0, heatloss: 0, lastthreemoves: []direction.Direction{}}
	pq := []crucible{cru}

	for len(pq) > 0 {
		cru := pq[0]
		pq = pq[1:]

		if cru.heatloss > leastHeatloss {
			continue
		}

		if grid.inGoal(cru.y, cru.x) {
			if cru.heatloss < leastHeatloss {
				fmt.Println("GOAL", cru)
				leastHeatloss = cru.heatloss
			}
			continue
		}

		// early stopping if you are suboptimal
		enc := fmt.Sprintf("(%v, %v),%v", cru.y, cru.x, cru.lastthreemoves)
		if _, ok := cache[enc]; ok {
			continue
		} else {
			cache[enc] = cru.heatloss
		}

		// make sure we want to spawn more?
		newDirecitons := cru.FindLegalDirections(grid)
		if len(newDirecitons) > 0 {
			for _, dir := range newDirecitons {
				NewCrucible := cru.SpawnNewCrucibleInDirection(dir, grid)
				pq = append(pq, NewCrucible)
			}
		}
		sort.SliceStable(pq, func(i, j int) bool {
			return pq[i].heatloss < pq[j].heatloss
		})
	}

}
func partB(input []string) {
	grid := parseInput(input)
	fmt.Println(grid)
	leastHeatloss := math.MaxInt64
	cache := make(map[string]int)

	cru := ultracrucible{y: 0, x: 0, heatloss: 0, lastTenMoves: []direction.Direction{}}
	pq := []ultracrucible{cru}

	for len(pq) > 0 {
		cru := pq[0]
		pq = pq[1:]
		// fmt.Println(cru)

		if cru.heatloss > leastHeatloss {
			continue
		}
		// must end on 4x of the same direction
		if grid.inGoal(cru.y, cru.x) && cru.lastNmovesEqual(4) {
			if cru.heatloss < leastHeatloss {
				fmt.Println("GOAL", cru)
				fmt.Println(cru.allmoves)

				leastHeatloss = cru.heatloss
			}
			continue
		}

		// enc := fmt.Sprintf("(%v, %v),%v,%v,%v", cru.y, cru.x, latestMove, lastFourbool, lastTenBool)
		enc := fmt.Sprintf("(%v, %v),%v,", cru.y, cru.x, cru.lastTenMoves)
		// fmt.Println("enc", enc)
		if _, ok := cache[enc]; ok {
			continue
		} else {
			cache[enc] = cru.heatloss
		}

		// make sure we want to spawn more?
		newDirecitons := cru.FindLegalDirections(grid)
		if len(newDirecitons) > 0 {
			for _, dir := range newDirecitons {
				NewCrucible := cru.SpawnNewCrucibleInDirection(dir, grid)
				pq = append(pq, NewCrucible)
			}
		}
		sort.SliceStable(pq, func(i, j int) bool {
			return pq[i].heatloss < pq[j].heatloss
		})
	}

}

func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	// partA(splitInput)
	partB(splitInput)
}
