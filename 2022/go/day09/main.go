package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func make_set(items string) map[rune]bool {
	set := make(map[rune]bool)

	for _, item := range items {
		set[item] = true
	}
	return set
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	headX, headY, tailX, tailY := 0, 0, 0, 0
	tail_visited := make(map[[2]int]bool) // "x,y" as string

	for i := 0; i < len(splitInput); i++ {
		var direction string
		var steps int
		fmt.Sscanf(splitInput[i], "%s %d", &direction, &steps)
		for j := 0; j < steps; j++ {

			// fmt.Println(direction, steps)

			switch direction {
			case "R":
				fmt.Println("R")
				headX += 1
			case "L":
				fmt.Println("L")
				headX -= 1
			case "U":
				fmt.Println("U")
				headY += 1
			case "D":
				fmt.Println("D")
				headY -= 1
			}

			fmt.Println("head", headX, headY)
			// head and tail on top of eachother
			// if headX == tailX && headY == tailY {
			// 	// fmt.Println("head", headX, headY)
			// 	// fmt.Println("tail", tailX, tailY)
			// 	continue
			// }
			// head and tail touching
			if headX-tailX <= Abs(1) && headY-tailY <= Abs(1) {

				// fmt.Println("toucin yo")
				// fmt.Println("head", headX, headY)
				// fmt.Println("tail", tailX, tailY)
				continue
			}

			if headX == tailX || headY == tailY {
				// tail and head in line:
				if headX > tailX {
					tailX += 1
				} else if headX < tailX {
					tailX -= 1
				} else if headY > tailY {
					tailY += 1
				} else if headY < tailY {
					tailY -= 1
				}
			} else {
				// tail and head diagonal:
				if headX > tailX {
					tailX += 1
				} else if headX < tailX {
					tailX -= 1
				}
				if headY > tailY {
					tailY += 1
				} else if headY < tailY {
					tailY -= 1
				}
			}

			loc := [2]int{tailX, tailY}
			fmt.Println("tail", tailX, tailY)
			tail_visited[loc] = true

		}

	}

	fmt.Println("p1 ans:", len(tail_visited))

}

// func calc_tail_motion(Headdirection string, steps int, headX *int, headY *int, tailX *int, tailY *int) [2]int {

// 	var new_tail_coords [2]int
// 	switch Headdirection {
// 	case "R":
// 		headX += steps
// 	case "L":
// 		headX -= steps
// 	case "U":
// 		headY += steps
// 	case "D":
// 		headY -= steps
// 	}

// 	if headX == tailX && headY == tailY {

// 	}

// 	return new_tail_coords
// }
