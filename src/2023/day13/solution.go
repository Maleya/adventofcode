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

// find the axis of reflection and returns [0] if none found (legacy)
func (p Pattern) findvertical() []int {
	// ax is the 0 indexed rightmost element of the mirror plane.
	// therefore no natural reflections occur at 0
	var potentialAxes, finalAxes []int
	rows := len(p.Pattern)
	cols := len(p.Pattern[0])

	// compare columns (pairwise)
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
		return []int{0}
	}
	for _, ax := range potentialAxes {
		// expand the search outward
		maxdist := min(ax-1, cols-ax-1)
		if maxdist == 0 {
			// catch case where no expansion needed:
			finalAxes = append(finalAxes, ax)
			continue
		}

		var counter int
		for i := 1; i <= maxdist; i++ {
			for j := 0; j < rows; j++ {
				if p.Pattern[j][ax-i-1] == p.Pattern[j][ax+i] {
					counter++
					// fmt.Println("match on row", j, p.Pattern[j][ax-i-1], p.Pattern[j][ax+i])
				}
			}
			if counter == rows*maxdist {
				finalAxes = append(finalAxes, ax)
				continue
			}
		}
	}
	if len(finalAxes) == 0 {
		return []int{0}
	} else {
		return finalAxes
	}
}
func (p Pattern) findhorizontal() []int {
	var axes []int
	p = p.transpose()
	axes = append(axes, p.findvertical()...)
	return axes
}

// iterate all combinations of symbol swaps and record results
func (p Pattern) iterateandtest() int {
	pat := p.Pattern
	var originalHort, originalVert int

	originalHort = p.findhorizontal()[0]
	originalVert = p.findvertical()[0]

	for i := 0; i < len(pat); i++ {
		for j := 0; j < len(pat[i]); j++ {

			// make a copy of pattern slice
			new_pat := make([][]string, len(pat))
			for k := range pat {
				new_pat[k] = make([]string, len(pat[k]))
				copy(new_pat[k], pat[k])
			}

			new_pat[i][j] = flipsymbol(p.Pattern[i][j])
			p := Pattern{Pattern: new_pat}
			vert := p.findvertical()

			for _, ax := range vert {
				if ax != 0 && ax != originalVert {
					// fmt.Printf("vertical match at %d, smudge in (%d,%d) \n", ax, i, j)
					// p.Print()
					return ax
				}

			}
			hort := p.findhorizontal()
			for _, ax := range hort {
				if ax != 0 && ax != originalHort {
					// fmt.Printf("horizontal match at %d, smudge in (%d,%d) \n", ax, i, j)
					// p.Print()
					return ax * 100
				}
			}
		}
	}
	return 0
}

func flipsymbol(r string) string {

	switch r {
	case "#":
		return "."
	case ".":
		return "#"
	default:
		panic("wrong string detected")
	}
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
		if vert[0] == 0 && hort[0] == 0 {
			fmt.Println("no axis found for pattern", i)
			break
		}
		// fmt.Printf("for pattern %d found %d horizontal, %d vertical \n", i, hort, vert)
		vertical += vert[0]
		horizontal += hort[0] * 100
	}
	fmt.Println("ans part a:", vertical+horizontal)

}

func partB(input []string) {
	sum := 0
	c := createCollection(input)
	for i, pattern := range c.Patterns {
		// fmt.Printf("pattern %d: \n", i)
		found := pattern.iterateandtest()
		if found == 0 {
			fmt.Println("no axis found for pattern", i)
		}
		sum += found
	}
	fmt.Println("ans part b", sum)
}

func main() {
	// load_file := example_input
	// load_file := example_input1
	load_file := input

	splitInput := strings.Split(strings.TrimSpace(load_file), "\n")

	partA(splitInput)
	partB(splitInput)

}
