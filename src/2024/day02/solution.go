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
	if increase_counter == len(r.levels)-1 || decrease_counter == len(r.levels)-1 {
		return true
	}
	return false

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

func (r report) safe() bool {
	return r.strictlyIncreasingOrDecreasing() && r.differby(1, 3)
}
func (r report) popAll() []report {
	var reports []report
	for i := 0; i < len(r.levels); i++ {
		popedLevels := make([]int, len(r.levels)-1)
		copy(popedLevels, r.levels[:i])
		copy(popedLevels[i:], r.levels[i+1:])
		re := report{popedLevels, r.errors}
		reports = append(reports, re)
	}
	return reports
}

func parseInputPartA(input []string) []report {
	var reports []report
	for _, line := range input {
		r := report{} // needs a &
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
		reports = append(reports, r)
	}
	return reports
}

func partA(input []string) {

	reports := parseInputPartA(input)
	var safe_reports int
	for _, r := range reports {

		if r.safe() {
			safe_reports++
		}
	}
	fmt.Println("part_a:", safe_reports)

}
func partB(input []string) {
	reports := parseInputPartA(input)
	var safe_reports int
	for _, r := range reports {
		if r.safe() {
			safe_reports++
		} else {
			for _, r := range r.popAll() {
				if r.safe() {
					safe_reports++
					break
				}
			}
		}
	}
	fmt.Println("part_b:", safe_reports)

}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)

}
