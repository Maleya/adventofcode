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

//go:embed example1.txt
var example_input1 string

type Pattern struct {
	Pattern [][]string
}

type Collection struct {
	Patterns []Pattern
}

func (p Pattern) Print() {
	for _, line := range p.Pattern {
		fmt.Println(line)
	}
}

// refactor to seperate lib?
func (p Pattern) transpose() Pattern {
	xl := len(p.Pattern[0])
	yl := len(p.Pattern)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = p.Pattern[j][i]
		}
	}
	return Pattern{Pattern: result}
}

func (p Pattern) findvertical() int {
	// ax is the 0 indexed rightmost element of the mirror plane.
	var potentialAxes []int
	rows := len(p.Pattern)
	cols := len(p.Pattern[0])
	// fmt.Printf("rows: %d, cols: %d\n", rows, cols)
	// find compare the first two columns:
	for ax := 1; ax < len(p.Pattern[0]); ax++ {
		var counter int
		for _, line := range p.Pattern {
			if line[ax-1] != line[ax] {
				continue
			} else {
				counter++
			}
		}
		if counter == rows {
			potentialAxes = append(potentialAxes, ax)
		}
	}
	if len(potentialAxes) == 0 {
		return 0
	}
	// problems: The test case from reddit is giving 0,0
	// and you dont catch if you dont find anything initially.
	// its on the edge.. meaning the bottom part of the code cant be used and is not handled.
	// Expand the search outwards
	// fmt.Println("mirror axes to check:", potentialAxes)
	for _, ax := range potentialAxes {
		// fmt.Println("checking mirror axis", ax)

		maxdist := min(ax-1, cols-ax-1) //! broken for ax=1
		if maxdist == 0 {
			return ax
		}
		//* catch special case of no expansion needed.

		var counter int
		for i := 1; i <= maxdist; i++ {
			// fmt.Println("expanding search to", ax-i-1, ax+i)
			for j := 0; j < rows; j++ {
				if p.Pattern[j][ax-i-1] == p.Pattern[j][ax+i] {
					counter++
					// fmt.Println("match on row", j, p.Pattern[j][ax-i-1], p.Pattern[j][ax+i])
				}
			}
			// fmt.Println("counter", counter)
			if counter == rows*maxdist {
				// fmt.Println("found axis", ax)
				return ax
			} else {
				// fmt.Println("axis", ax, "test incomplete", counter, "/", rows*maxdist)
			}
		}
	}
	return 0
}
func (p Pattern) findhorizontal() int {
	p = p.transpose()
	return p.findvertical()
	// fmt.Println(ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func createCollection(input []string) Collection {
	var patterns []Pattern
	var currentPattern [][]string

	for _, line := range input {
		if line == "" {
			if len(currentPattern) > 0 {
				patterns = append(patterns, Pattern{Pattern: currentPattern})
				currentPattern = nil
			}
		} else {
			currentPattern = append(currentPattern, strings.Split(line, ""))
		}
	}

	// Handle the last pattern if not followed by an empty line
	if len(currentPattern) > 0 {
		patterns = append(patterns, Pattern{Pattern: currentPattern})
	}

	return Collection{Patterns: patterns}
}

func partA(input []string) {
	c := createCollection(input)
	var vertical, horizontal int
	for i, pattern := range c.Patterns {
		vert := pattern.findvertical()
		hort := pattern.findhorizontal()
		if vert == 0 && hort == 0 {
			fmt.Println("no axis found for pattern!!!!!", i)
			break
		}
		fmt.Printf("for pattern %d found %d horizontal, %d vertical \n", i, hort, vert)
		vertical += vert
		horizontal += hort * 100
	}
	fmt.Println("ans:", vertical+horizontal)

}
func partB(input []string) {
	fmt.Println("part_b:")
}

func main() {
	// load_file := example_input
	// load_file := example_input1
	load_file := input

	splitInput := strings.Split(strings.TrimSpace(load_file), "\n")

	partA(splitInput)
	// partB(splitInput)

}
