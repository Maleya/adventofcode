package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func stringToASCII(input string) []int {
	asciiCodes := make([]int, len(input))

	for i, char := range input {
		asciiCodes[i] = int(char)
	}

	return asciiCodes
}

func hash(s string) int {
	currentVal := 0
	ascii := stringToASCII(s)

	for _, elem := range ascii {
		currentVal += elem
		currentVal *= 17
		currentVal %= 256
	}

	return currentVal
}

//go:embed example.txt
var example_input string

func partA(input []string) {
	sum := 0
	for _, elem := range input {
		sum += hash(elem)
	}
	fmt.Println("part_a:", sum)

}

type lens struct {
	// the elements inside the linked list
	label       string
	focalLength int
}

type box struct {
	boxNo  int
	lenses *list.List
}

func initBoxes(n int) []box {
	boxes := make([]box, 256)
	for i := 0; i < n; i++ {
		boxes[i] = box{boxNo: i, lenses: list.New()}
	}
	return boxes
}

func focusingPower(boxes []box) int {
	sum := 0
	for boxNo, box := range boxes {
		if box.lenses.Len() != 0 {
			iter := 1
			for e := box.lenses.Front(); e != nil; e = e.Next() {
				sum += (1 + boxNo) * iter * e.Value.(lens).focalLength
				iter++
			}
		}
	}
	return sum
}

func seedBoxes(boxes []box, input []string) []box {
	for _, elem := range input {
		if label, _, ok := strings.Cut(elem, "-"); ok {
			// we remove lenses if they exist
			boxNo := hash(label)

			if lensObj, lensinBox := listcontains(boxes[boxNo].lenses, label); lensinBox {
				boxes[boxNo].lenses.Remove(lensObj)
			}

		} else if label, focalStr, ok := strings.Cut(elem, "="); ok {
			// we place lenses into new boxes
			focal, _ := strconv.Atoi(focalStr)
			boxNo := hash(label)

			lensObj, lensinBox := listcontains(boxes[boxNo].lenses, label)
			newLens := lens{label: label, focalLength: focal}

			if !lensinBox {
				// add to box at the back
				boxes[boxNo].lenses.PushBack(newLens)

			} else {
				// replace existing label
				boxes[boxNo].lenses.InsertBefore(newLens, lensObj)
				boxes[boxNo].lenses.Remove(lensObj)
			}
		}
	}
	return boxes
}

func partB(input []string) {
	boxes := initBoxes(256)
	boxes = seedBoxes(boxes, input)
	ans := focusingPower(boxes)
	fmt.Println("part_b:", ans)

}

func listcontains(l *list.List, label string) (*list.Element, bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(lens).label == label {
			return e, true
		}
	}
	return nil, false
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), ",")

	partA(splitInput)
	partB(splitInput)

}
