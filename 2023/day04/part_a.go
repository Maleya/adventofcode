package main

import (
	"fmt"
	"io"
	"math"
	"os"
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
func pointscounter(no_matches int) int {

	if no_matches == 0 {
		return 0
	}
	ans := math.Pow(2.0, float64(no_matches-1))
	return int(ans)

}

func main() {
	// fileName := "example.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)
	part_a := 0

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range splitInput {
		// fmt.Println(line)

		_, numbers1, numbers2, _ := parseInput(line)
		// fmt.Println("Card Name:", cardName)
		overlap := countoverlap(numbers1, numbers2)
		// fmt.Println("overlap", overlap)
		final := pointscounter(overlap)
		part_a += final

		// fmt.Println("Numbers Set 1:", numbers1)
		// fmt.Println("Numbers Set 2:", numbers2)

	}
	fmt.Println("part_a:", part_a)
}
