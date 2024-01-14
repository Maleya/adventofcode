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

type rocks struct {
	Positions       [][]string
	roundRocksCount map[int][]int // key: column, val: occurance count per group
	squareRocksLoc  map[int][]int // key: column, val: row locations
}

func (r rocks) print() {
	for _, row := range r.Positions {
		fmt.Println(row)
	}
}
func (r *rocks) countRocks() {
	ncols := len(r.Positions[0])
	nrows := len(r.Positions)

	// reset the existing:
	r.roundRocksCount = make(map[int][]int)
	r.squareRocksLoc = make(map[int][]int)

	for j := 0; j < ncols; j++ {
		count := 0
		for i := 0; i < nrows; i++ {

			if r.Positions[i][j] == "O" {
				count++
			}
			if r.Positions[i][j] == "#" {
				r.squareRocksLoc[j] = append(r.squareRocksLoc[j], i)
				r.roundRocksCount[j] = append(r.roundRocksCount[j], count)
				count = 0
			}
			// check end reached
			if i == nrows-1 {
				r.roundRocksCount[j] = append(r.roundRocksCount[j], count)

			}
		}
	}
}
func (r *rocks) cycle(n int) {
	for i := 0; i < 4*n; i++ {
		r.tiltNorth()
		r.rotatemap()
	}
}

func (r rocks) calculateWeight() int {
	points := 0
	nrows := len(r.Positions)
	for i, row := range r.Positions {
		for j := range row {
			if r.Positions[i][j] == "O" {
				points += nrows - i
			}
		}
	}
	return points
}

func (r *rocks) tiltNorth() {
	newPos := make([][]string, len(r.Positions))

	// make a new map and copy over the squares rocks
	for i, row := range r.Positions {
		newPos[i] = make([]string, len(row))
		for j, char := range row {
			if string(char) == "#" {
				newPos[i][j] = string(char)
			} else {
				newPos[i][j] = "."
			}
		}
	}
	r.countRocks()
	for col := 0; col < len(r.roundRocksCount); col++ {
		incrementer := 0

		for nround, round := range r.roundRocksCount[col] {
			if round == 0 {
				if nround != 0 {
					incrementer++
				}
				continue
			}

			if nround == 0 {
				for s := 0; s < round; s++ {
					newPos[s][col] = "O"
				}
				continue
			}
			// if square stones exist in col
			if len(r.squareRocksLoc[col]) > 0 {
				squareLoc := r.squareRocksLoc[col][incrementer]
				startidx := squareLoc + 1
				// fmt.Println("next loop:", startidx, startidx+round)
				for s := startidx; s < startidx+round; s++ {
					newPos[s][col] = "O"
				}
				incrementer++
			}
		}
	}
	r.Positions = newPos
}

func (r *rocks) rotatemap() {
	rows := len(r.Positions)
	cols := len(r.Positions[0])

	// Create a new rotated r.Positions
	rotated := make([][]string, cols)
	for i := range rotated {
		rotated[i] = make([]string, rows)
	}
	// Populate the rotated r.Positions
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-i-1] = r.Positions[i][j]
		}
	}

	r.Positions = rotated
}

func parseInput(input []string) rocks {
	Pos := make([][]string, len(input))

	for i, row := range input {
		Pos[i] = make([]string, len(row))
		for j, char := range row {
			Pos[i][j] = string(char)
		}
	}
	return rocks{Positions: Pos,
		roundRocksCount: make(map[int][]int),
		squareRocksLoc:  make(map[int][]int),
	}
}

func partA(input []string) {
	rocks := parseInput(input)
	rocks.tiltNorth()
	ans := rocks.calculateWeight()
	fmt.Println("part_a:", ans)

}

func partB(input []string) {
	rocks := parseInput(input)
	rocks.cycle(1000)
	ans := rocks.calculateWeight()
	fmt.Println("part_b:", ans)
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

}
