package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type MapData struct {
	Title      string
	Rows       [][]int
	range_maps []func(input int) (int, bool)
}

type Funcmap struct {
	Title string
	funcs []func(input, dest, source, length int) int
}

func (m *MapData) calculate(input int) int {
	var in_range bool
	// fmt.Println("function input:", input)
	for i := 0; i < len(m.Rows); i++ {
		// fmt.Println("loop start:", input)
		input, in_range = range_map(input, m.Rows[i][0], m.Rows[i][1], m.Rows[i][2])
		// fmt.Println("loop end:", input)
		if in_range {
			// remember: numbers pass only one map.
			break
		}
	}
	return input
}

func parseInput(lines []string) ([]int, []MapData) {
	var seedData []int
	var mapData []MapData

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "seeds:") {
			seedData = parseSeeds(line)
		} else if strings.Contains(line, "map:") {
			mapTitle := line
			var rows [][]int
			i++ // Move to the next line
			for i < len(lines) && lines[i] != "" {
				row := parseRow(lines[i])
				rows = append(rows, row)
				i++
			}
			mapData = append(mapData, MapData{Title: mapTitle, Rows: rows})
		}
	}

	return seedData, mapData
}

func parseSeeds(line string) []int {
	seedStr := strings.TrimPrefix(line, "seeds:")
	seedValues := parseIntSlice(seedStr)
	return seedValues
}

func parseRow(line string) []int {
	return parseIntSlice(line)
}

func parseIntSlice(s string) []int {
	fields := strings.Fields(s)
	var nums []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return nil
		}
		nums = append(nums, num)
	}
	return nums
}

func range_map(input, dest, source, length int) (int, bool) {
	if input < source || input > source+length {
		return input, false
	} else {
		return input + (dest - source), true
	}
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	seedData, mapData_list := parseInput(splitInput)
	part_a := 99999999999

	for _, seed := range seedData {
		for i, m := range mapData_list {
			output := m.calculate(seed)
			if i == len(mapData_list)-1 {
				// fmt.Println(m.Title)
				// fmt.Printf("Input: %d, Output: %d\n", seed, output)
				if output < part_a {
					part_a = output
				}

			}
			seed = output
		}
	}
	fmt.Println("part_a:", part_a)
}
