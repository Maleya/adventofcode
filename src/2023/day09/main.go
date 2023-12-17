package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type all_history struct {
	histories []history
}

func (ah *all_history) process_input(input []string) {

	for _, line := range input {
		digits := strings.Fields(line)
		var single_history []int
		for _, digit := range digits {
			digit_n, _ := strconv.Atoi(digit)
			single_history = append(single_history, digit_n)
		}
		h := newHistory(single_history)
		ah.histories = append(ah.histories, *h)

	}
}

type history struct {
	initalHistory []int
	diffmap       map[int][]int
}

func (h *history) CalcAlldiff() {
	i := 0
	diff, done := calcDiff(h.initalHistory)
	h.diffmap[i] = diff

	for !done {
		diff, done = calcDiff(h.diffmap[i])
		i++
		h.diffmap[i] = diff
		if len(diff) <= 1 {
			break
		}
	}
}

func (h *history) projectHistory() int {
	// reminder to run calcalldiff first.
	n_rows := len(h.diffmap)
	initial_length := len(h.initalHistory) - 1
	sum := 0

	for i := 0; i < n_rows-1; i++ {
		row := h.diffmap[i]
		last := row[len(row)-1]
		sum += last

	}
	output := h.initalHistory[initial_length] + sum
	return output

}

func (h *history) projectHistory_backward() int {
	// reminder to run calcalldiff first.
	n_rows := len(h.diffmap)
	total := 0
	for i := n_rows - 1; i >= 0; i-- {
		row := h.diffmap[i]
		first := row[0]
		total = first - total
	}
	output := h.initalHistory[0] - total
	return output

}

func calcDiff(seq []int) (diff_seq []int, allzero bool) {
	nseq := len(seq)
	zeroCounter := 0
	for i := 0; i < nseq-1; i++ {
		diff := seq[i+1] - seq[i]
		if diff == 0 {
			zeroCounter++
		}
		// fmt.Println(seq[i+1], "-", seq[i], "=", diff)
		diff_seq = append(diff_seq, diff)
	}
	if zeroCounter == len(diff_seq) {
		allzero = true
	} else {
		allzero = false
	}
	return diff_seq, allzero
}

func newHistory(input []int) *history {

	h := history{
		initalHistory: input,
		diffmap:       make(map[int][]int)}
	return &h
}

func part_a(input []string) int {
	ah := all_history{}
	ah.process_input(input)

	sum := 0
	for _, h := range ah.histories {
		h.CalcAlldiff()
		sum += h.projectHistory()
	}

	fmt.Println("part_a:", sum)
	return sum
}

func part_b(input []string) int {
	ah := all_history{}
	ah.process_input(input)

	sum := 0
	for _, h := range ah.histories {
		h.CalcAlldiff()
		sum += h.projectHistory_backward()
	}

	fmt.Println("part_b:", sum)
	return sum
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")
	part_a(splitInput)
	part_b(splitInput)

}
