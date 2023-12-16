package main

import (
	"fmt"
	"io"
	"os"
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

func main() {
	fileName := "example_a.txt"
	// fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		os.Exit(1)
	}

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	var final_sum int
	for _, line := range splitInput {

		digits_str := extract_digits(line)
		final_sum += combine_first_last(digits_str)
	}
	fmt.Println("final answer", final_sum)
}
