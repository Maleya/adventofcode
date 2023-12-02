package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// fileName := "example_a.txt"
	// fileName := "example_b.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range splitInput {

		fmt.Println(line)

	}
}
