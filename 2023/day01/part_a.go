package main

// do this with channels?
import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// return the digits in the string in order
func strtoint(s string) int {
	fmt.Println(s)
	for _, char := range s {
		// println("char", string(char))
		if unicode.IsDigit(char) {
			if i, _ := strconv.Atoi(string(char)); i != 0 { //no need for if here
				println("test", i)
				fmt.Println("it was a digit:", string(char))
			}

		}

	}
	return 0
}

func main() {
	fileName := "example.txt"
	// fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, line := range splitInput {
		strtoint(line)
	}
}
