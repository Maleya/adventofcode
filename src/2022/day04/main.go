package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fully_contains_range(r1_start, r1_end, r2_start, r2_end int) bool {

	// check if range 1 inside 2:
	if r1_start <= r2_end && r1_start >= r2_start && r1_end <= r2_end && r1_end >= r2_start {
		return true
		//  check if range 2 inside 1
	} else if r2_start <= r1_end && r2_start >= r1_start && r2_end <= r1_end && r2_end >= r1_start {
		return true
	} else {
		return false
	}
}

func contains_overlap(r1_start, r1_end, r2_start, r2_end int) bool {

	// r1_start or r1 end contained in range 2
	if (r1_start <= r2_end && r1_start >= r2_start) || (r1_end <= r2_end && r1_end >= r2_start) {
		return true
		//  check if range 2 inside 1
	} else if (r2_start <= r1_end && r2_start >= r1_start) || (r2_end <= r1_end && r2_end >= r1_start) {
		return true
	} else {
		return false
	}
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(readFile)
	defer readFile.Close()
	counter := 0
	for sc.Scan() {
		line := sc.Text()
		elf_ranges := strings.Split(line, ",")

		first_elf_range := strings.Split(elf_ranges[0], "-")
		second_elf_range := strings.Split(elf_ranges[1], "-")

		r1_start, _ := strconv.Atoi(first_elf_range[0])
		r1_end, _ := strconv.Atoi(first_elf_range[1])
		r2_start, _ := strconv.Atoi(second_elf_range[0])
		r2_end, _ := strconv.Atoi(second_elf_range[1])

		if contains_overlap(r1_start, r1_end, r2_start, r2_end) {
			counter += 1
		}
	}
	fmt.Println("total:", counter)
}
