package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

type report struct {
	levels []int
	errors int
}

func newReport() *report {
	return &report{
		levels: []int{},
		errors: 0,
	}
}

func (r *report) strictlyIncreasingOrDecreasing() bool {
	var increase_counter, decrease_counter int

	for i := 0; i < len(r.levels)-1; i++ {
		if r.levels[i] < r.levels[i+1] {
			increase_counter++
		}
		if r.levels[i] > r.levels[i+1] {
			decrease_counter++
		}
	}
	max_counter := max(increase_counter, decrease_counter)
	r.errors += len(r.levels) - 1 - max_counter
	return increase_counter == len(r.levels)-1 || decrease_counter == len(r.levels)-1
}

func (r *report) differby(least, most int) bool {
	var differ_counter int
	for i := 0; i < len(r.levels)-1; i++ {
		diff := math.Abs(float64(r.levels[i] - r.levels[i+1]))
		if diff >= float64(least) && diff <= float64(most) {
			differ_counter++
		}
	}
	errors := len(r.levels) - 1 - differ_counter
	r.errors += errors
	return differ_counter == len(r.levels)-1
}

// func (r *report) safe() bool {
// 	return r.strictlyIncreasingOrDecreasing() && r.differby(1, 3)
// }

func parseInput(input []string) []report {
	var reports []report
	for _, line := range input {
		r := newReport()
		levels := strings.Split(line, " ")
		for _, level := range levels {
			level = strings.TrimRight(level, "\r") // this is needed for windows?
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				fmt.Printf("Error converting %s to integer: %v\n", level, err)
				continue
			}
			r.levels = append(r.levels, levelInt)
		}
		reports = append(reports, *r)
	}
	return reports
}

func partA(input []string) {
	reports := parseInput(input)
	var safe_reports int
	for i := 0; i < len(reports); i++ {
		reports[i].differby(1, 3)
		reports[i].strictlyIncreasingOrDecreasing()

		if reports[i].errors == 0 {
			safe_reports++
		}

		// fmt.Println("report:", reports[i])
	}
	fmt.Println("part a:", safe_reports)

	// fmt.Println("Number of safe reports:", safe_reports)
}
func removeElement(slice []int, i int) []int {
    return append(slice[:i], slice[i+1:]...)
}

func partB(input []string) {
	reports := parseInput(input)
	var safe_reports int
	for i := 0; i < len(reports); i++ {
		for j := 0; j < len(reports[i].levels); j++ {
			
		if reports[i].errors <= 1 {
			safe_reports++
		}

		// fmt.Println("report:", reports[i])
	}
	fmt.Println("part b:", safe_reports)
}
func main() {
	load_file := example_input
	// load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

}
