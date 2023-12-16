package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	input, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var letter string
	var word string
	for i := 0; scanner.Scan(); i++ {
		word = string(scanner.Text())
		letter = string(word[0])
		fmt.Println(i, letter, word)
	}
}
