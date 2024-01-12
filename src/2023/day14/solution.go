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
func (r rocks) countRocks() {
	// number of round stones per group
	ncols := len(r.Positions[0])
	nrows := len(r.Positions)

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

func (r rocks) calculateWeight() int {
	// each square rock is at: len(r.Positions)-squarePos
	// subsequent round rocks are incremented till another square is reached.
	sum := 0
	nrows := len(r.Positions)
	// for each element in round, if not 0, get a stone to shift it by index.

	for col := 0; col < len(r.roundRocksCount); col++ {
		// fmt.Println("COLUMN:", col)
		// fmt.Println("has round:", r.roundRocksCount[col])
		// fmt.Println("has square:", r.squareRocksLoc[col])
		incrementer := 0

		for nround, round := range r.roundRocksCount[col] {
			// fmt.Println("round =", round)
			if round == 0 {
				// no round stones in group, move on.
				if nround != 0 {
					incrementer++
				}
				continue
			}
			// fmt.Println(r.squareRocksLoc[col], "at index", incrementer)

			if nround == 0 {
				// the first group
				startidx := nrows
				for s := startidx; s > nrows-round; s-- {
					sum += s
					// fmt.Println(s)
				}
				continue
			}

			// if square stones exist in col
			if len(r.squareRocksLoc[col]) > 0 {
				// fmt.Println("incrementer", incrementer)
				squareLoc := r.squareRocksLoc[col][incrementer]
				// fmt.Println("sq loc", squareLoc)
				startidx := nrows - squareLoc - 1
				// incrementer++
				// fmt.Println("loop goes from", startidx, "to", nrows-round)
				// first group starts from the top:
				for s := startidx; s > startidx-round; s-- {
					// fmt.Println(s)
					sum += s

				}
				incrementer++
			}

			// fmt.Println("we in here", startidx, nrows-round)

		}
	}
	return sum
	// fmt.Println("sum:", sum)
}

func (r rocks) rotatemap() rocks {
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

	return rocks{Positions: rotated}
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
	rocks.print()
	rocks.countRocks()
	ans := rocks.calculateWeight()
	fmt.Println("part_a:", ans)

}

func partB(input []string) {
	rocks := parseInput(input)
	rocks.print()
	rocks.countRocks()
	// ans := rocks.calculateWeight()

	// fmt.Println("part_b:")
}

func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	// for _, elem := range splitInput {
	// 	fmt.Println(elem)
	// }
	partA(splitInput)
	partB(splitInput)

}
