package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

type Row struct {
	SpringString string
	Sizes        []int
	DoneInGroup  int
}

// make a string representation of Row for a map
func (r Row) ToString() string {
	strSlice := make([]string, len(r.Sizes))

	for i, num := range r.Sizes {
		strSlice[i] = strconv.Itoa(num)
	}

	sizesStr := strings.Join(strSlice, ",")
	return fmt.Sprintf("%s-%s-%d", r.SpringString, sizesStr, r.DoneInGroup)
}

func parseInputString(input string) Row {
	parts := strings.Split(input, " ")

	leftSide := parts[0]
	variablesStr := parts[1]

	// Split the variables string on comma and convert to integers
	var sizes []int
	for _, vStr := range strings.Split(variablesStr, ",") {
		variable, err := parseVariable(vStr)
		if err != nil {
			// Handle error if conversion fails
			fmt.Println("Error parsing variable:", err)
			return Row{}
		}
		sizes = append(sizes, variable)
	}

	// Create and return the struct
	return Row{
		SpringString: leftSide + ".", //append a "." to find EOL
		Sizes:        sizes,
		DoneInGroup:  0,
	}
}

func parseInputStringPartB(input string) Row {
	// ugly copy-paste job to parse for partB. Should be refactored. One day.
	parts := strings.Split(input, " ")

	leftSide := parts[0]
	variablesStr := parts[1]

	// Split the variables string on comma and convert to integers
	var sizes []int
	for _, vStr := range strings.Split(variablesStr, ",") {
		variable, err := parseVariable(vStr)
		if err != nil {
			// Handle error if conversion fails
			fmt.Println("Error parsing variable:", err)
			return Row{}
		}
		sizes = append(sizes, variable)
	}

	// repeat the left side 5 times
	repeatedSpring := strings.Join(strings.Split(strings.Repeat(leftSide+"?", 5), ""), "")
	repeatedSpring = repeatedSpring[:len(repeatedSpring)-1]

	// Repeat Sizes 5 times
	var repeatedSizes []int
	for i := 0; i < 5; i++ {
		repeatedSizes = append(repeatedSizes, sizes...)
	}

	return Row{
		SpringString: repeatedSpring + ".", //append a "." to find EOL,
		Sizes:        repeatedSizes,
		DoneInGroup:  0,
	}
}

func countSolutions(r Row, cache map[string]int) int {
	var solutions int
	var possible []string

	// check return condition:
	if len(r.SpringString) == 0 {
		if len(r.Sizes) == 0 && r.DoneInGroup == 0 {
			// we correctly handled and closed all groups:
			return 1
		}
		return 0
	}

	if char := string(r.SpringString[0]); char == "?" {
		possible = []string{"#", "."}
	} else {
		possible = []string{char}
	}

	for _, spring := range possible {
		if spring == "#" {
			// extend current group
			NextRow := Row{
				SpringString: r.SpringString[1:],
				Sizes:        r.Sizes,
				DoneInGroup:  r.DoneInGroup + 1,
			}

			// load from cache or update if not in cache:
			if cached_sol, ok := cache[NextRow.ToString()]; ok {
				solutions += cached_sol
			} else {
				sol := countSolutions(NextRow, cache)
				cache[NextRow.ToString()] = sol
				solutions += sol
			}
		} else {
			if r.DoneInGroup > 0 {
				// close a group if its saturated:
				if len(r.Sizes) > 0 && r.DoneInGroup == r.Sizes[0] {

					NextRow := Row{
						SpringString: r.SpringString[1:],
						Sizes:        r.Sizes[1:],
						DoneInGroup:  0,
					}
					// load from cache or update if not in cache:
					if cached_sol, ok := cache[NextRow.ToString()]; ok {
						solutions += cached_sol
					} else {
						sol := countSolutions(NextRow, cache)
						cache[NextRow.ToString()] = sol
						solutions += sol
					}
				}

			} else {
				// not in a group, move on to next symbol
				NextRow := Row{
					SpringString: r.SpringString[1:],
					Sizes:        r.Sizes,
					DoneInGroup:  0,
				}
				// load from cache or update if not in cache:
				if cached_sol, ok := cache[NextRow.ToString()]; ok {
					solutions += cached_sol
				} else {
					sol := countSolutions(NextRow, cache)
					cache[NextRow.ToString()] = sol
					solutions += sol
				}
			}
		}
	}
	return solutions
}

func parseVariable(vStr string) (int, error) {
	// Convert the variable string to an integer
	variable, err := strconv.Atoi(vStr)
	if err != nil {
		return 0, err
	}
	return variable, nil
}

func partA(input []string) {
	sum := 0
	cache := make(map[string]int)

	for _, line := range input {

		r := parseInputString(line)
		sol := countSolutions(r, cache)
		sum += sol
	}
	fmt.Println("sum part a:", sum)
}

func partB(input []string) {
	sum := 0
	cache := make(map[string]int)

	for _, line := range input {
		r := parseInputStringPartB(line)
		sol := countSolutions(r, cache)
		sum += sol
	}
	fmt.Println("sum part b:", sum)
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)
}
