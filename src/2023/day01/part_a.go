package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func extract_digits(s string) string {
	var sb strings.Builder

	for _, char := range s {
		if unicode.IsDigit(char) {
			sb.WriteString(string(char))
		}
	}
	return sb.String()
}

// return the first and last digits
func combine_first_last(s string) int {
	firstDigit := string(s[0])
	lastDigit := string(s[len(s)-1])
	combined_digits, _ := strconv.Atoi(firstDigit + lastDigit)
	return combined_digits
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	var final_sum int
	for _, line := range splitInput {

		digits_str := extract_digits(line)
		final_sum += combine_first_last(digits_str)
	}
	fmt.Println("final answer", final_sum)
}
