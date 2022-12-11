package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// moves like a stack
func move_boxes(stack *[9][]string, move, from, to int) {

	for i := 0; i < move; i++ {

		n := len(stack[from]) - 1
		elem := stack[from][n]        // Top element on stack
		stack[from] = stack[from][:n] // pop the last

		stack[to] = append(stack[to], elem) // add to new stack.
	}

}

// moves `move` boxes at once.
func move_several_boxes(stack *[9][]string, move, from, to int) {

	n := len(stack[from])
	elem := stack[from][n-move : n]        // Top elements on stack
	stack[from] = stack[from][:n-move]     // pop the last set as one
	stack[to] = append(stack[to], elem...) // add to new stack.
}

func main() {

	// hardcode the stacks :(
	var arrayofstacks [9][]string
	arrayofstacks[0] = []string{"S", "M", "R", "N", "W", "J", "V", "T"}
	arrayofstacks[1] = []string{"B", "W", "D", "J", "Q", "P", "C", "V"}
	arrayofstacks[2] = []string{"B", "J", "F", "H", "D", "R", "P"}
	arrayofstacks[3] = []string{"F", "R", "P", "B", "M", "N", "D"}
	arrayofstacks[4] = []string{"H", "V", "R", "P", "T", "B"}
	arrayofstacks[5] = []string{"C", "B", "P", "T"}
	arrayofstacks[6] = []string{"B", "J", "R", "P", "L"}
	arrayofstacks[7] = []string{"N", "C", "S", "L", "T", "Z", "B", "W"}
	arrayofstacks[8] = []string{"L", "S", "G"}

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(readFile)
	defer readFile.Close()
	for sc.Scan() {
		line := sc.Text()

		re := regexp.MustCompile("[0-9]+")
		move_instructions := re.FindAllString(line, -1)
		move, _ := strconv.Atoi(move_instructions[0])
		from, _ := strconv.Atoi(move_instructions[1])
		to, _ := strconv.Atoi(move_instructions[2])

		move_several_boxes(&arrayofstacks, move, from-1, to-1)
		// move_boxes(&arrayofstacks, move, from-1, to-1)

	}
	var answer string
	for i, elem := range arrayofstacks {
		answer += elem[len(arrayofstacks[i])-1]
	}
	println("answer:", answer)
}
