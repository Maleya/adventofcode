package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

//go:embed example1.txt
var example_input1 string

//go:embed example2.txt
var example_input2 string

// to be refactored for future use qq
type maze struct {
	mazemap         [][]string
	startLoc        coords
	lastElem        coords
	visited         map[coords]bool
	pipeToDirection map[string][]string
	DirectionToPipe map[string][]string
}

// bool for checking if S should be counted in raytracing
func (m *maze) countSinRay() bool {
	req := coords{y: 1, x: 0}
	// we as long as one of the arms of is south we need to count it.

	output := m.nextPipeLoc(m.startLoc)
	for _, elem := range output {
		diff := coords{elem.y - m.startLoc.y, elem.x - m.startLoc.x}
		if diff == req {
			return true
		}
	}
	return false
}

// returns false if on the boundary
// uses visited to check the boundary.
func (m *maze) isInside(loc coords) bool {
	counter := 0
	countPipeParts := []string{"|", "7", "F"}
	if !m.inbounds(loc) {
		panic("location out of map bounds")
	}
	if m.visited[loc] {
		return false
	}
	for dx := loc.x; dx < m.lastElem.x+1; dx++ {
		ray_loc := coords{loc.y, dx}
		if m.visited[ray_loc] {
			pipe, _ := m.get(ray_loc)
			// identify pipe under S:
			if pipe == "S" {
				if m.countSinRay() {
					counter++ // continue
				}
			}
			if Isin(pipe, countPipeParts) {
				counter++
				// fmt.Println("encountered:", pipe, "at loc:", ray_loc, "counter", counter)
			}

		}

	}
	return counter%2 != 0
}

func (m *maze) checkallinside() int {
	counter := 0
	// implement raytracing
	for y := 0; y <= m.lastElem.y; y++ {
		for x := 0; x <= m.lastElem.x; x++ {
			newLoc := coords{y, x}
			if m.isInside(newLoc) {
				counter++

			}
		}

	}
	return counter
}

// to be refactored for future use qq
type coords struct {
	y int
	x int
}

func (c coords) up() coords {
	new := c
	new.y--
	return new
}
func (c coords) down() coords {
	new := c
	new.y++
	return new
}
func (c coords) left() coords {
	new := c
	new.x--
	return new
}
func (c coords) right() coords {
	new := c
	new.x++
	return new
}
func Isin(s string, inlist []string) bool {
	for _, a := range inlist {
		if a == s {
			return true
		}
	}
	return false
}

func (m *maze) nextPipeLoc(loc coords) (output []coords) {

	curr_pipe, ok := m.get(loc)
	if ok != nil {
		panic("location out of map bounds")
	}
	// fmt.Println("current pipe:", curr_pipe)

	viable_directions := m.pipeToDirection[curr_pipe]

	for _, dir := range viable_directions {
		required_pipes := m.DirectionToPipe[dir]
		switch dir {
		case "N":
			next_loc := loc.up()
			next_pipe, ok := m.get(next_loc)
			if ok == nil && Isin(next_pipe, required_pipes) {
				// fmt.Println("went north to:", next_pipe, "at loc:", next_loc)
				output = append(output, next_loc)
			}
		case "E":
			next_loc := loc.right()
			next_pipe, ok := m.get(next_loc)
			if ok == nil && Isin(next_pipe, required_pipes) {
				// fmt.Println("went east to:", next_pipe, "at loc:", next_loc)
				output = append(output, next_loc)
			}
		case "S":
			next_loc := loc.down()
			next_pipe, ok := m.get(next_loc)
			if ok == nil && Isin(next_pipe, required_pipes) {
				// fmt.Println("went south to:", next_pipe, "at loc:", next_loc)
				output = append(output, next_loc)
			}
		case "W":
			next_loc := loc.left()
			next_pipe, ok := m.get(next_loc)
			if ok == nil && Isin(next_pipe, required_pipes) {
				// fmt.Println("went west to:", next_pipe, "at loc:", next_loc)
				output = append(output, next_loc)
			}
		}
	}

	return output
}

// get a mazeitem if it is in bounds
// to be refactored for future use qq
func (m *maze) get(loc coords) (output string, err error) {
	if !m.inbounds(loc) {
		err = errors.New("loc coordinate out of map bounds")
		return "", err
	}
	if m.inbounds(loc) {
		output = m.mazemap[loc.y][loc.x]
	}
	return output, err

}

// check if location is in bounds
func (m *maze) inbounds(loc coords) bool {
	if (loc.x <= m.lastElem.x && loc.x >= 0) && (loc.y <= m.lastElem.y && loc.y >= 0) {
		return true
	} else {
		return false
	}
}

func init_maze(input []string) *maze {
	dy := len(input)
	dx := len(input[0])
	var startingLoc coords

	mazemap := make([][]string, dy)
	for i := range mazemap {
		mazemap[i] = make([]string, dx)
	}
	for i, line := range input {
		for j, char := range line {
			mazemap[i][j] = string(char)
			if string(char) == "S" {
				startingLoc.y = i
				startingLoc.x = j
			}
		}
	}
	pipeToDirection := map[string][]string{
		"|": {"N", "S"},
		"-": {"E", "W"},
		"L": {"N", "E"},
		"J": {"N", "W"},
		"7": {"S", "W"},
		"F": {"S", "E"},
		"S": {"N", "E", "S", "W"},
	}
	DirectionToPipe := map[string][]string{
		"N": {"F", "7", "|"},
		"E": {"-", "7", "J"},
		"S": {"L", "J", "|"},
		"W": {"F", "-", "L"},
	}

	m := maze{
		mazemap:         mazemap,
		startLoc:        startingLoc,
		lastElem:        coords{y: dy - 1, x: dx - 1},
		visited:         make(map[coords]bool),
		pipeToDirection: pipeToDirection,
		DirectionToPipe: DirectionToPipe,
	}
	return &m

}
func part_a_b(splitInput []string) {
	m := init_maze(splitInput)
	start_loc := m.startLoc
	step_counter := 0

	var queue []coords
	m.visited[start_loc] = true

	// go only one direction from S by taking first element.
	queue = append(queue, m.nextPipeLoc(start_loc)[0])

	for len(queue) > 0 {
		current_loc := queue[0] // pop
		queue = queue[1:]       // deque
		current, _ := m.get(current_loc)
		if current == "S" {
			break
		}
		step_counter++
		m.visited[current_loc] = true
		new_locs := m.nextPipeLoc(current_loc)
		for _, loc := range new_locs {
			if !m.visited[loc] {
				queue = append(queue, loc)
				m.visited[loc] = true
			}
		}
	}
	fmt.Println("final steps", step_counter)
	fmt.Println("part a answer:", (step_counter+1)/2)
	fmt.Println("points inside (part b):", m.checkallinside())
}

func main() {
	// load_file := example_input
	load_file := input
	// load_file := example_input2
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	part_a_b(splitInput)
	// part_b(splitInput)
}
