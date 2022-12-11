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

func make_set(items string) map[rune]bool {
	set := make(map[rune]bool)

	for _, item := range items {
		set[item] = true
	}
	return set
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	total_priority := 0
	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()

	for fileScanner.Scan() {
		first_set := make_set(fileScanner.Text())
		fileScanner.Scan()
		second_set := make_set(fileScanner.Text())
		fileScanner.Scan()
		third_set := make_set(fileScanner.Text())

		// check for mutal pairs:

		for first_elf_item := range first_set {
			if second_set[first_elf_item] && third_set[first_elf_item] {
				total_priority += letter_priority(string(first_elf_item))
			}
		}
	}
	fmt.Println(total_priority)
}
