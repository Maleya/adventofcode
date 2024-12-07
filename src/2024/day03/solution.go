package main

import (
	"adventofcode/pkg/stringslice"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func extractMulPairs(inputStr string) [][2]int {
	// Define the regular expression pattern to match valid mul(X,Y) operations
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(inputStr, -1)

	// Slice to store the pairs of integers
	var pairs [][2]int

	// Iterate over the matches and convert them to integer pairs
	for _, match := range matches {
		if len(match) == 3 {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				pairs = append(pairs, [2]int{x, y})
			}
		}
	}

	return pairs
}

func partA(input []string) {
	var sum int
	for _, elem := range input {
		mulpairs := extractMulPairs(elem)
		for _, pair := range mulpairs {
			sum += pair[0] * pair[1]
		}

	}
	fmt.Println("part_a:", sum)

}
func partB(input []string) {
	const findOps = `((?:don't|do|mul)\(\d*,?\d*\))`
	const findMuli = `mul\((\d+),(\d+)\)`

	reOps := regexp.MustCompile(findOps)
	reMul := regexp.MustCompile(findMuli)

	var sum int
	do := true
	for _, elem := range input {
		matches := reOps.FindAllString(elem, -1)
		for _, match := range matches {
			switch match {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if do {
					multiplicand := reMul.FindStringSubmatch(match)
					multiplicandInt := stringslice.AtoiSlice(multiplicand[1:])
					sum += multiplicandInt[0] * multiplicandInt[1]
				}
			}
		}
	}
	fmt.Println("part_b:", sum)
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

}
