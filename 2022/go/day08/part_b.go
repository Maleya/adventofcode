package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func look_right(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {

			current_tree := trees[i][j]
			for jj := j + 1; jj < dx; jj++ {
				if trees[i][jj] >= current_tree {
					vis[i][j] = jj - j
					break
				}
				if jj == dx-1 {
					vis[i][j] = jj - j
				}

			}
		}
	}
	return vis
}

// same as look right with transposed tree grid and vis
func look_down(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	for i := 0; i < dy; i++ {
		// tree_row := trees[i][:]

		// fmt.Println(tree_row)
		for j := 0; j < dx; j++ {
			current_tree := trees[j][i]
			for jj := j + 1; jj < dx; jj++ {
				// fmt.Println(trees[j][i], tree_row[jj])
				if trees[jj][i] >= current_tree {
					vis[j][i] = jj - j
					// fmt.Println("visible:", jj-j)
					break
				}
				if jj == dx-1 {
					vis[j][i] = jj - j
				}

			}
		}
	}
	return vis
}

func look_left(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	for i := 0; i < dy; i++ {

		for j := dx - 1; j > 0; j-- {
			current_tree := trees[i][j]

			for jj := j - 1; jj >= 0; jj-- {
				if trees[i][jj] >= current_tree {
					vis[i][j] = j - jj
					// fmt.Println("visible:", jj-j)
					break
				}
				if jj == 0 {
					vis[i][j] = j - jj
				}

			}
		}
	}
	return vis
}

// look up is look left with a transposed tree grid and vis
func look_up(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	for i := 0; i < dy; i++ {

		for j := dx - 1; j > 0; j-- {
			current_tree := trees[j][i]

			for jj := j - 1; jj >= 0; jj-- {
				if trees[jj][i] >= current_tree { // swapped
					vis[j][i] = j - jj
					// fmt.Println("visible:", jj-j)
					break
				}
				if jj == 0 {
					vis[j][i] = j - jj
				}

			}
		}
	}
	return vis
}

func main() {

	// input, _ := ioutil.ReadFile("example.txt")
	input, _ := ioutil.ReadFile("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// use (y,x) convention. Y is down. and first in tuple
	var dx, dy int
	dx = len(strings.TrimSpace(splitInput[0]))
	dy = len(splitInput)

	tree_grid := make([][]int, dy)
	vis_l := make([][]int, dy)
	vis_r := make([][]int, dy)
	vis_u := make([][]int, dy)
	vis_d := make([][]int, dy)

	for i := 0; i < dy; i++ {
		tree_grid[i] = make([]int, dx)
		vis_l[i] = make([]int, dx)
		vis_r[i] = make([]int, dx)
		vis_u[i] = make([]int, dx)
		vis_d[i] = make([]int, dx)

		row := splitInput[i]
		for j := 0; j < dx; j++ {
			height, _ := strconv.Atoi(string(row[j]))
			tree_grid[i][j] = height

		}
	}
	right := look_right(tree_grid, dy, dx, vis_r)
	left := look_left(tree_grid, dy, dx, vis_l)
	up := look_up(tree_grid, dy, dx, vis_u)
	down := look_down(tree_grid, dy, dx, vis_d)

	max := 0
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			product := right[i][j] * left[i][j] * up[i][j] * down[i][j]
			if product > max {
				max = product
			}
		}
	}
	fmt.Println("ans:", max)
}
