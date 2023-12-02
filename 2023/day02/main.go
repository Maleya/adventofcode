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

	max_colours := map[string]int{"red": 12, "green": 13, "blue": 14}
	var game_id_sum int

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
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
