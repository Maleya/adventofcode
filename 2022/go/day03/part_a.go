package main

import (
	"bufio"
	"fmt"
	"os"
)

func letter_priority(letter string) int {
	alphabet := make(map[string]int)

	var ch byte
	priority_counter := 1
	for ch = 'a'; ch <= 'z'; ch++ {
		alphabet[string(ch)] = priority_counter
		priority_counter += 1

	}
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet[string(ch)] = priority_counter
		priority_counter += 1

	}
	return alphabet[letter]
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()

	total_priority := 0
	for fileScanner.Scan() {
		text_line := fileScanner.Text()
		mid := len(text_line) / 2
		set_left := make(map[string]bool)
		set_right := make(map[string]bool)

		// make left and right into go routines because you can?
		for _, letter := range text_line[:mid] {

			set_left[string(letter)] = true
		}
		for _, letter := range text_line[mid:] {
			set_right[string(letter)] = true
		}

		for letter := range set_left {
			if set_right[letter] {
				total_priority += letter_priority(string(letter))
			}
		}
	}
	fmt.Println(total_priority)
	// fmt.Println(string(alphabet_small))
}
