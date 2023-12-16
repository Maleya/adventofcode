package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// rewrite this to be useful in the future.
func findAllOccurrences(input, substr string) []int {
	re := regexp.MustCompile(substr)
	matches := re.FindAllStringIndex(input, -1)

	var occurrences []int
	for _, match := range matches {
		occurrences = append(occurrences, match[0])
	}

	return occurrences
}

// this really should be cleaned up and split into smaller functions
func extract_digits_extended(input string, wordmap map[string]string) string {
	var resultBuilder strings.Builder
	var pos_list []int                   // list of keys we will sort
	pos_to_value := make(map[int]string) // corresponding values

	for key, digit := range wordmap {

		occurances := findAllOccurrences(input, key)
		if len(occurances) >= 0 {
			for _, pos := range occurances {
				pos_to_value[pos] = digit
				pos_list = append(pos_list, pos)
			}
		}
	}
	// loop and check for literal string digits
	for pos, char := range input {
		if unicode.IsDigit(char) {
			pos_to_value[pos] = string(char)
			pos_list = append(pos_list, pos)
		}
	}

	// build the final string as only digits
	sort.Ints(pos_list)
	for _, key := range pos_list {
		resultBuilder.WriteString(pos_to_value[key])

	}
	return strings.TrimSpace(resultBuilder.String())
}

// return the first and last digits
func combine_first_last(s string) int {
	firstDigit := string(s[0])
	lastDigit := string(s[len(s)-1])
	combined_digits, _ := strconv.Atoi(firstDigit + lastDigit)
	return combined_digits
}

func main() {
	// fileName := "example_a.txt"
	// fileName := "example_b.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	// make map of strings
	replacements_map := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	var final_sum int
	for _, line := range splitInput {

		// fmt.Println(line)
		digits_str := extract_digits_extended(line, replacements_map)
		// fmt.Println(digits_str)
		combined := combine_first_last(digits_str)
		// fmt.Println(combined)
		final_sum += combined
	}
	fmt.Println("final answer", final_sum)
}
