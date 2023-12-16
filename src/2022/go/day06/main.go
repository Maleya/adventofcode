package main

import (
	"bufio"
	"fmt"
	"os"
)

func char_repeated(input string, nr_repeats int) bool {
	set := make(map[rune]bool)
	for _, char := range input {
		set[char] = true
	}
	size := len(set)
	if size == nr_repeats {
		return false
	}
	return true
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(readFile)
	defer readFile.Close()
	sc.Scan()
	line := sc.Text()

	reapeat_length := 4

	for i := reapeat_length; i <= len(line); i++ {
		four_block := line[i-reapeat_length : i]
		if char_repeated(four_block, reapeat_length) == false {
			fmt.Println("ans:", i)
			break

		}

	}

}
