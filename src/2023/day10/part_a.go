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

// to be refactored for future use qq
type maze struct {
	mazemap  [][]string
	startLoc coords
	lastElem coords
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

// todo: rename!
// to be refactored for future use qq
func (m *maze) nextPipeLoc(loc coords) (output []coords) {
	if !m.inbounds(loc) {
		panic("location out of map bounds")
	}
	// directions:
	up_loc := loc.up()
	down_loc := loc.down()
	left_loc := loc.left()
	right_loc := loc.right()

	// requirements:
	req_up := []string{"F", "7", "|"}
	req_down := []string{"L", "J", "|"}
	req_left := []string{"F", "-", "L"}
	req_right := []string{"-", "7", "J"}

	up, ok := m.get(up_loc)
	if ok == nil && Isin(up, req_up) {
		fmt.Println("went up:", up, up_loc)
		output = append(output, up_loc)
	}

	down, ok := m.get(down_loc)
	if ok == nil && Isin(down, req_down) {
		fmt.Println("went down:", down, down_loc)
		output = append(output, down_loc)
	}
	left, ok := m.get(left_loc)
	if ok == nil && Isin(left, req_left) {
		fmt.Println("went left:", left, left_loc)
		output = append(output, left_loc)
	}
	right, ok := m.get(right_loc)
	if ok == nil && Isin(right, req_right) {
		fmt.Println("went right:", right, right_loc)
		output = append(output, right_loc)

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
		output = m.mazemap[loc.y][loc.x] // check that this works.
	}
	return output, err

}

// check if location is in bounds
// to be refactored for future use qq
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
	m := maze{
		mazemap:  mazemap,
		startLoc: startingLoc,
		lastElem: coords{y: dy - 1, x: dx - 1},
	}
	return &m

}
func part_a(splitInput []string) {
	m := init_maze(splitInput)
	start_loc := m.startLoc
	step_counter := 0

	var queue []coords
	visited := make(map[coords]bool)

	queue = append(queue, m.nextPipeLoc(start_loc)...)

	// todo: fix this janky ass queue.
	for len(queue) > 0 {
		new_locs := m.nextPipeLoc(queue[0])
		step_counter++
		for _, loc := range new_locs {
			if !visited[loc] {
				queue = append(queue, loc)
				visited[loc] = true
			}
		}
		// queue = append(queue, m.nextPipeLoc(queue[0])...) // First element
		queue = queue[1:] // Dequeue
	}
	fmt.Println("final steps", step_counter)

}

func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	part_a(splitInput)
	// m.get(up_loc)
	// m := init_maze(splitInput)
	// loc := coords{y: -1, x: 0}
	// fmt.Println(m.inbounds(loc))
	// fmt.Println(m.get(loc))
	// result :=
	// fmt.Println(m.nextPipeLoc(loc))
}
