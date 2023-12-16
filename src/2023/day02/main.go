package main

import (
	_ "embed"
	"fmt"
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

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	max_colours := map[string]int{"red": 12, "green": 13, "blue": 14}
	var game_id_sum int

	for _, fullgame := range splitInput {
		var game_id, round_counter int
		header_split := strings.Split(fullgame, ":")
		fmt.Sscanf(header_split[0], "Game %d", &game_id)

		all_rounds := strings.Split(header_split[1], ";")
		number_of_rounds := len(all_rounds)

		for _, round := range all_rounds {
			var digit int
			var color string
			cubes_drawn := make(map[string]int)

			for _, cube := range strings.Split(round, ",") {
				fmt.Sscanf(cube, "%d %s", &digit, &color)
				cubes_drawn[color] = digit
			}

			if game_possible(cubes_drawn, max_colours) {
				round_counter++
			}

		}
		if round_counter == number_of_rounds {
			game_id_sum += game_id
		}

	}
	fmt.Println("game_id_sum", game_id_sum)
}
