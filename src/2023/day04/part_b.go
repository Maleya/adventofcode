package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) (int, []int, []int, error) {
	var card_no int

	parts := strings.Split(input, "|")
	left_split := strings.Split(parts[0], ":")

	fmt.Sscanf(input, "Card %d", &card_no)

	num_left := strings.TrimSpace(left_split[1])
	num_right := strings.TrimSpace(parts[1])
	numbers1, _ := parseNumbers(strings.TrimSpace(num_left))
	numbers2, _ := parseNumbers(strings.TrimSpace(num_right))
	return card_no, numbers1, numbers2, nil
}

func parseNumbers(numbersStr string) ([]int, error) {
	numberStrings := strings.Fields(numbersStr)
	numbers := make([]int, len(numberStrings))

	for i, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %v", err)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func countoverlap(number1, numbers2 []int) int {
	//count overlap. How many of numbers2 are in numbers1
	counter := 0
	for _, num2 := range numbers2 {
		for _, num1 := range number1 {
			if num1 == num2 {
				counter++
			}

		}
	}
	return counter
}

func update_counter(card_counter map[int]int, card, overlap, max_update_thresh int) map[int]int {
	// fmt.Println()
	upper_lim := card + overlap + 1
	// fmt.Println("upper_lim", upper_lim)
	if upper_lim > max_update_thresh+1 { //testing here remove +1 {
		upper_lim = max_update_thresh
		// fmt.Println("adjusted_lim", card, card+overlap+1, upper_lim)
	}

	for i := card + 1; i < upper_lim; i++ {
		// fmt.Println("i", i)
		// fmt.Println("card", i, "incremented via", card)
		card_counter[i]++
	}
	return card_counter
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	part_a := 0
	part_b := 0
	card_copies := make(map[int]int)

	table_len := len(splitInput) //! change back
	for _, line := range splitInput {

		card_no, numbers1, numbers2, _ := parseInput(line)
		overlap := countoverlap(numbers1, numbers2)
		fmt.Println("Card:", card_no, "overlap", overlap)

		for i := 0; i < 1+card_copies[card_no]; i++ {
			update_counter(card_copies, card_no, overlap, table_len)
		}
		// update_counter(card_copies, card_no, adjusted_overlap, table_len)

		// fmt.Println("Numbers Set 1:", numbers1)
		// fmt.Println("Numbers Set 2:", numbers2)

	}
	fmt.Println("part_a:", part_a)
	fmt.Println(card_copies)
	for _, elem := range card_copies {
		part_b += elem
	}
	part_b += table_len
	fmt.Println("part_b:", part_b)
}
