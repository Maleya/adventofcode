package main

import (
	"adventofcode/pkg/file"
	"fmt"
)

func main() {
	// fmt.Println(abs)
	a := file.ReadInput("src/2023/day08/example.txt")
	for _, line := range a {
		fmt.Println(line)

	}
}
