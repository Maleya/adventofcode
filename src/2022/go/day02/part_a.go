package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rock_paper_scisors_game(a, b int) int {

	var round_points int
	switch {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	case a == b:
		// draw
		// fmt.Println("draw")
		round_points = 3
	case (a+1)%3 == b:
		// win
		// fmt.Println("you win")
		round_points = 6

	case (a+2)%3 == b:
		// loss
		// fmt.Println("you lose")
		round_points = 0

	}
	return (b + 1) + round_points

}

func main() {

	to_int := make(map[string]int)
	to_int["A"] = 0
	to_int["B"] = 1
	to_int["C"] = 2
	to_int["X"] = 0
	to_int["Y"] = 1
	to_int["Z"] = 2

	readFile, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var total int

	for fileScanner.Scan() {
		pair := strings.Fields(fileScanner.Text())
		fmt.Println((pair))

		elf, you := to_int[pair[0]], to_int[pair[1]]
		points := rock_paper_scisors_game(elf, you)

		total += points

	}
	readFile.Close()
	fmt.Println("total:", total)

}
