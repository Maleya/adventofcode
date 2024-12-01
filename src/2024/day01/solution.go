package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func parseInputPartA(input []string) ([]int, []int) {
	var leftList, rightList []int

	for _, elem := range input {
		nums := strings.Split(elem, "   ")
		if len(nums) != 2 {
			continue
		}

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			// Handle the error appropriately
			continue
		}

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			// Handle the error appropriately
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	return leftList, rightList
}

func distance(x int, y int) int {
	return int(math.Abs(float64(x - y)))
}

func partA(input []string) {
	var sum int
	leftList, rightList := parseInputPartA(input)

	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Println(leftList, rightList)

	for i := 0; i < len(leftList); i++ {
		sum += distance(leftList[i], rightList[i])
		fmt.Println(distance(leftList[i], rightList[i]))
	}

	fmt.Println("part_a:", sum)

}
func countOccurences(input int, array []int) int {
	occurences := 0

	for _, elem := range array {
		if elem == input {
			occurences++
		}
	}

	return occurences
}

func partB(input []string) {
	cache := make(map[int]int)
	leftList, rightList := parseInputPartA(input)

	for _, elem := range leftList {
		// check if elem in cache

		

		fmt.Println("part_b:")
	}
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

}
