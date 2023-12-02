package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func game_possible(results, max_colours map[string]int) bool {
	for colour, count := range results {
		if count > max_colours[colour] {
			return false
		}
	}
	return true
}

func main() {
	// fileName := "example.txt"
	// fileName := "example_b.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	var prod_sum int

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, fullgame := range splitInput {
		var game_id, product_game int
		cube_max := make(map[string]int)

		header_split := strings.Split(fullgame, ":")
		fmt.Sscanf(header_split[0], "Game %d", &game_id)

		for _, round := range strings.Split(header_split[1], ";") {
			var digit int
			var color string
			product_game = 1 // reset for each round

			for _, cube := range strings.Split(round, ",") {
				fmt.Sscanf(cube, "%d %s", &digit, &color)
				if digit > cube_max[color] {
					cube_max[color] = digit
				}
			}
			for key := range cube_max {
				product_game *= cube_max[key]
			}
		}
		prod_sum += product_game
	}

	fmt.Println("sum of products:", prod_sum)
}
