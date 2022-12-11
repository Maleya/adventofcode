package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rock_paper_scisors_game1(a, b int) int {
	// we get b+1 points for the shape on top of the round points
	var round_points int

	switch {
	case a == b: // draw
		// draw
		round_points = 3
	case (a+1)%3 == b:
		// win
		round_points = 6

	case (a+2)%3 == b:
		// lose
		round_points = 0

	}
	return (b + 1) + round_points

}

func rock_paper_scisors_game2(a, b int) int {
	// same point rules but decides the outcome
	var throw_shape int
	var round_points int

	switch {
	case b == 1:
		// draw needed
		throw_shape = a
		round_points = 3

	case b == 2:
		// win needed
		throw_shape = (a + 1) % 3
		round_points = 6

	case b == 0:
		// loss needed
		throw_shape = (a + 2) % 3
		round_points = 0

	}
	return round_points + throw_shape + 1

}

func main() {
	shape_to_int := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	var total int

	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pair := strings.Fields(fileScanner.Text()) //split by whitespace
		elf, you := shape_to_int[pair[0]], shape_to_int[pair[1]]
		total += rock_paper_scisors_game1(elf, you)
	}
	readFile.Close()
	fmt.Println("total:", total)

}
