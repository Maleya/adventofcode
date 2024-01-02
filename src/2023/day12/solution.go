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

func (r Row) ToString() string {
	sizesStr := fmt.Sprint(r.SpringString)
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

func countSolutions(r Row) int {
	// todo: use a cache of type map[string]bool
	var solutions int
	var possible []string

	// fmt.Println(r.SpringString, r.Sizes, r.DoneInGroup)

	// check return condition:
	if len(r.SpringString) == 0 {
		if len(r.Sizes) == 0 && r.DoneInGroup == 0 {
			// we correctly handled and closed all groups:
			return 1
		}
		return 0

	}
	char := string(r.SpringString[0])

	if char == "?" {
		possible = []string{"#", "."}
	} else {
		possible = []string{char}
	}

	for _, spring := range possible {
		if spring == "#" {
			// extend current group
			// fmt.Println("-->", r.SpringString)
			solutions += countSolutions(Row{
				SpringString: r.SpringString[1:],
				Sizes:        r.Sizes,
				DoneInGroup:  r.DoneInGroup + 1,
			})
		} else {
			if r.DoneInGroup > 0 {
				// close a group if its saturated:
				if len(r.Sizes) > 0 && r.DoneInGroup == r.Sizes[0] {
					// fmt.Println("-->", r.SpringString)

					solutions += countSolutions(Row{
						SpringString: r.SpringString[1:],
						Sizes:        r.Sizes[1:],
						DoneInGroup:  0,
					})
				}

			} else {
				// not in a group, move on to next symbol
				// fmt.Println("-->", r.SpringString)
				solutions += countSolutions(Row{
					SpringString: r.SpringString[1:],
					Sizes:        r.Sizes,
					DoneInGroup:  0,
				})

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
	cache := make(map[string]bool)

	for _, line := range input {

		r := parseInputString(line)
		fmt.Println("initial input", r)
		fmt.Println("adapted input", r.ToString())
		sol := countSolutions(r)
		fmt.Println("no of arrangements:", sol)
		sum += sol
	}
	fmt.Println("sum:", sum)

}
func partB(input []string) {
	fmt.Println("part_b:")
}

func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	// partB(splitInput)

}

// plan:
