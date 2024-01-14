package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func stringToASCII(input string) []int {
	asciiCodes := make([]int, len(input))

	for i, char := range input {
		asciiCodes[i] = int(char)
	}

	return asciiCodes
}

func hash(s string) int {
	currentVal := 0
	ascii := stringToASCII(s)

	for _, elem := range ascii {
		currentVal += elem
		currentVal *= 17
		currentVal %= 256
	}

	return currentVal
}

//go:embed example.txt
var example_input string

func partA(input []string) {
	sum := 0
	for _, elem := range input {
		sum += hash(elem)
	}
	fmt.Println("part_a:", sum)

}
func partB(input []string) {
	fmt.Println("part_b:")
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), ",")

	partA(splitInput)
	partB(splitInput)

}
