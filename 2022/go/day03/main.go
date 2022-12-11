package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// array sum function
func addArray(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	// fileScanner.Split(bufio.ScanLines)

	sum := 0
	var max_slice []int

	for fileScanner.Scan() {
		text_line := fileScanner.Text()

		if text_line == "" {
			// new elf:
			max_slice = append(max_slice, sum)
			sum = 0 // reset counter

		} else {
			// its a int line
			snack, err := strconv.Atoi(text_line)
			if err != nil {
				panic(err)
			}
			sum += snack
		}
	}
	readFile.Close()

	// fmt.Println("max", max_slice)
	sort.Ints(max_slice)
	top_3 := max_slice[len(max_slice)-3:]
	fmt.Println("top 3:", top_3)

	ans := addArray(max_slice[len(max_slice)-3:]...)
	fmt.Println("sum:", ans)

}
