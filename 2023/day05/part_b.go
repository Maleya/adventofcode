package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type MapData struct {
	Title string
	Rows  [][]int
}

type Funcmap struct {
	Title string
	funcs []func(input, dest, source, length int) int
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

func computePairs(input []int) []int {
	var result []int

	for i := 0; i < len(input); i += 2 {
		start := input[i]
		rng := input[i+1]

		// Populate the result slice with computed values
		for j := 0; j < rng; j++ {
			result = append(result, start+j)
		}
	}

	return result
}

func (m *MapData) calculate(input int) int {
	// fmt.Println(m.Title)
	for i := 0; i < len(m.Rows); i++ {
		dest := m.Rows[i][0]
		source := m.Rows[i][1]
		length := m.Rows[i][2]

		if input >= source && input < source+length {
			// fmt.Println(input, input+(dest-source))
			return input + (dest - source)
		}
	}
	// fmt.Println(input)
	return input
}

func compute_loc(input int, mapData_list []MapData) int {
	var output int
	for _, m := range mapData_list {
		output = m.calculate(input)
		input = output
	}
	return output
}

func main() {
	// fileName := "example.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	seedData, mapData_list := parseInput(splitInput)
	start := time.Now()

	part_b := int(^uint(0) >> 1)
	seedData = computePairs(seedData)
	// seedData = []int{82}
	var output int

	for _, seed := range seedData {
		// fmt.Println("computed:", compute_loc(seed, mapData_list))
		output = compute_loc(seed, mapData_list)
		if output < part_b {
			part_b = output
		}
	}

	fmt.Println("part_b:", part_b)
	duration := time.Since(start)
	fmt.Println(duration)
}
