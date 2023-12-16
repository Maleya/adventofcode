package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruction struct {
	command string
	arg     int
}

func main() {
	// input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("example1.txt")
	input, _ := ioutil.ReadFile("input.txt")

	cycle := 1
	x := 1
	var sig_str, sum, duration int
	const dark string = "."
	const light string = "█"
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	for i := 0; i < len(splitInput); i++ {

		inst := instruction{}
		fmt.Sscanf(splitInput[i], "%s %d", &inst.command, &inst.arg)
		current_op := inst
		if current_op.command == "noop" {
			duration = 1
		} else if current_op.command == "addx" {
			duration = 2
		}

		for d := 0; d < duration; d++ {
			sig_str = x * cycle
			if (cycle-20)%40 == 0 {
				sum += sig_str
			}
			col := (cycle - 1) % 40
			if col <= x+1 && col >= x-1 {
				fmt.Print(light)

			} else {
				fmt.Print(dark)

			}

			cycle += 1
		}
		x += current_op.arg

	}
	fmt.Println("sum", sum)
	// make terminal 40 wide.

}
