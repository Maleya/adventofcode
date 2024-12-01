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

func partA(input []string) {
	for _, elem := range input {
		fmt.Println(elem)
	}
	fmt.Println("part_a:")

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
