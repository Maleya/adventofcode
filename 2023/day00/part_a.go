package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// fileName := "example.txt"
	fileName := "input.txt"

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
}
