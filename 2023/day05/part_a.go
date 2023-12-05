package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SeedData struct {
	Seeds []int
}

type MapData struct {
	Title string
	Rows  [][]int
}

func parseInput(lines []string) (SeedData, []MapData) {
	var seedData SeedData
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

func parseSeeds(line string) SeedData {
	seedStr := strings.TrimPrefix(line, "seeds:")
	seedValues := parseIntSlice(seedStr)
	return SeedData{Seeds: seedValues}
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

func main() {
	fileName := "example.txt"
	// fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")
	seedData, mapData := parseInput(splitInput)

	fmt.Println("Seeds:", seedData.Seeds)
	fmt.Println("Maps:")
	for _, m := range mapData {
		fmt.Println("Title:", m.Title)
		fmt.Println("Rows:", m.Rows)
		fmt.Println()
	}
}
