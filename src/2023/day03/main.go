package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type symbol_location struct {
	//
	symbol  string
	loc     location
	part_no []int
}
type location struct {
	//
	row_no    int
	row_start int
	row_end   int
}

func find_symbols(input []string) map[string][]symbol_location {
	locations := make(map[string][]symbol_location)
	re := regexp.MustCompile(`[^A-Za-z0-9.]`)

	for line_nr, line := range input {

		// fmt.Println(line)
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			symbol := line[match[0]:match[1]]
			loc := location{row_no: line_nr, row_start: match[0], row_end: match[1]}
			symbol_data := symbol_location{symbol: symbol, loc: loc}

			locations[symbol] = append(locations[symbol], symbol_data)
		}
	}
	return locations
}
func findDigits(input []string) map[string][]symbol_location {
	locations := make(map[string][]symbol_location)
	re := regexp.MustCompile(`\d+`)

	for line_nr, line := range input {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			digits := line[match[0]:match[1]]
			loc := location{row_no: line_nr, row_start: match[0], row_end: match[1]}
			symbol_data := symbol_location{symbol: digits, loc: loc}

			locations[digits] = append(locations[digits], symbol_data)
		}
	}
	return locations
}

// updates symbols in place with part_no fields.
func match_numbers_to_symbols(input []string, symbols, numbers map[string][]symbol_location) {
	sum_part1 := 0

	for i, symbol_type := range symbols {
		for j, symbol_loc := range symbol_type {

			for _, number := range numbers {
				for _, number_loc := range number {
					number_row := number_loc.loc.row_no
					symbol_row := symbol_loc.loc.row_no
					if number_row <= symbol_row+1 && number_row >= symbol_row-1 {

						search_start := number_loc.loc.row_start - 1
						search_end := number_loc.loc.row_end

						if symbol_loc.loc.row_start >= search_start && symbol_loc.loc.row_start <= search_end {
							part_int, _ := strconv.Atoi(number_loc.symbol)
							// update the symbols struct:
							symbols[i][j].part_no = append(symbols[i][j].part_no, part_int)
							sum_part1 += part_int
						} else {
							continue
						}
					} else {
						continue
					}
				}
			}
		}
	}
	fmt.Println("part_a", sum_part1)
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	symbols_loc := find_symbols(splitInput)
	digits_loc := findDigits(splitInput)

	match_numbers_to_symbols(splitInput, symbols_loc, digits_loc)

	//part 2
	final_sum := 0
	for _, symb := range symbols_loc["*"] {
		if prod := 1; len(symb.part_no) > 1 {
			for _, part := range symb.part_no {
				prod *= part
			}
			final_sum += prod
		}
	}

	fmt.Println("part_b", final_sum)
}
