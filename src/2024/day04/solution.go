package main

import (
	coord "adventofcode/pkg/coord"
	"adventofcode/pkg/enums/compass"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

// Grid struct to hold the grid data and dimensions
type Grid struct {
	data   [][]rune
	width  int
	height int
}

// NewGrid initializes a new Grid from a slice of strings
func NewGrid(input []string) *Grid {
	height := len(input)
	width := len(input[0])
	data := make([][]rune, height)
	for i := range input {
		data[i] = []rune(input[i])
	}
	return &Grid{
		data:   data,
		width:  width,
		height: height,
	}
}

// GetChar returns the character at the given coordinates, or a zero rune if out of bounds
func (g *Grid) GetChar(c coord.Coordinates) rune {
	if g.InBounds(c) {
		return g.data[c.Y][c.X]
	}
	panic(fmt.Sprintf("coordinates out of bounds: %v", c))
}

// InBounds checks if the given coordinates are within the grid boundaries
func (g *Grid) InBounds(c coord.Coordinates) bool {
	return c.X >= 0 && c.Y < g.width && c.Y >= 0 && c.Y < g.height
}

func partA(input []string) {
	magicWord := []rune{'X', 'M', 'A', 'S'}
	// visited := make(map[coord.Coordinates]bool)
	fmt.Println("XMAS in runes", magicWord)

	// look for X, then iterate through the compass directions to find the next rune.

	g := NewGrid(input)
	c := coord.Coordinates{X: 2, Y: 2}
	char := g.GetChar(c)

	// this is how you iterate through the compass directions
	for dir := compass.North; dir <= compass.NW; dir++ {
		fmt.Println("dir:", dir)
		new := c.Move(dir)
		fmt.Println("new:", new)
	}
	fmt.Printf("Character at %v: %c\n", c, char)
	fmt.Println("part_a:", magicWord)
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

}
