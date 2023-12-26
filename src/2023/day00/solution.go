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
	fmt.Println("part_a:")

}
func partB(input []string) {
	fmt.Println("part_b:")
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	for _, elem := range splitInput {
		fmt.Println(elem)
	}
	partA(splitInput)
	partB(splitInput)

}
