package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coords struct {
	x int
	y int
}

type motion struct {
	direction string
	distance  int
}

// includes overlap/super position
func is_adjacent(pos1 coords, pos2 coords) bool {
	if Abs(pos1.x-pos2.x) <= 1 && Abs(pos1.y-pos2.y) <= 1 {
		return true
	} else {
		return false
	}
}

func get_tail_position(head_pos coords, tail_pos coords) coords {
	if is_adjacent(head_pos, tail_pos) == true {
	} else if head_pos.y == tail_pos.y {
		if head_pos.x > tail_pos.x {
			tail_pos.x += 1
		} else if head_pos.x < tail_pos.x {
			tail_pos.x -= 1
		}
	} else if head_pos.x == tail_pos.x {

		if head_pos.y > tail_pos.y {
			tail_pos.y += 1
		} else if head_pos.y < tail_pos.y {
			tail_pos.y -= 1
		}

	} else {
		if head_pos.y > tail_pos.y {
			tail_pos.y += 1
		} else if head_pos.y < tail_pos.y {
			tail_pos.y -= 1
		}
		if head_pos.x > tail_pos.x {
			tail_pos.x += 1
		} else if head_pos.x < tail_pos.x {
			tail_pos.x -= 1
		}
	}
	return tail_pos
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	// input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("example1.txt")
	input, _ := ioutil.ReadFile("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	positions := make([]coords, 10)
	motion := motion{}
	tail_visited_part2 := make(map[coords]bool)
	tail_visited_part1 := make(map[coords]bool)

	for i := 0; i < len(splitInput); i++ {
		fmt.Sscanf(splitInput[i], "%s %d", &motion.direction, &motion.distance)
		// fmt.Println(motion)
		for j := 0; j < motion.distance; j++ {

			switch motion.direction {
			case "R":
				positions[0].x += 1
			case "L":
				positions[0].x -= 1
			case "U":
				positions[0].y += 1
			case "D":
				positions[0].y -= 1
			}
			for kn := 1; kn < len(positions); kn++ {
				positions[kn] = get_tail_position(positions[kn-1], positions[kn])

			}
			// fmt.Println("len", len(positions))
			current_tail := positions[len(positions)-1]
			current_tail_part1 := positions[1] //
			tail_visited_part2[current_tail] = true
			tail_visited_part1[current_tail_part1] = true
		}

	}

	fmt.Println("ans rope len 2:", len(tail_visited_part1))
	fmt.Println("ans: rop len 10:", len(tail_visited_part2))

}
