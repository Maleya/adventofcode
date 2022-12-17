package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// use (y,x) convention. Y is down. and first in tuple
	var x_len, y_len int
	x_len = len(strings.TrimSpace(splitInput[0]))
	y_len = len(splitInput)

	tree_grid := make([][]int, y_len)
	vis := make([][]int, y_len)

	fmt.Println("gridsize:", y_len, x_len)

	for i := 0; i < y_len; i++ {
		tree_grid[i] = make([]int, x_len)
		vis[i] = make([]int, x_len)

		row := splitInput[i]
		for j := 0; j < x_len; j++ {
			height, _ := strconv.Atoi(string(row[j]))
			tree_grid[i][j] = height

			if i == 0 || i == y_len-1 {
				vis[i][j] = 1

			}
			if j == 0 || j == x_len-1 {
				vis[i][j] = 1

			}
		}
	}

	fmt.Println(tree_grid)
	fmt.Println(vis)
}
