package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func find_vis_left_to_right(trees [][]int, dy int, dx int, vis [][]int) [][]int {

	for i := 0; i < dy; i++ {
		row_max := 0
		for j := 0; j < dx; j++ {
			if trees[i][j] > row_max {
				vis[i][j] = 1
				row_max = trees[i][j]
			}
		}
	}
	return vis
}
func find_vis_up_to_down(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	// this is really just left to right with transposed grids
	for i := 0; i < dy; i++ {
		row_max := 0
		// fmt.Println("row reset")
		for j := 0; j < dx; j++ {
			// fmt.Println(trees[j][i])
			// trees and vis transposed
			if trees[j][i] > row_max {
				vis[j][i] = 1

				// fmt.Println("vis=1", "at", j, i)
				row_max = trees[j][i]
				// fmt.Println("new max:", row_max)
			}
		}
	}
	return vis
}

func find_vis_down_to_up(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	// this is really just right to left with transposed grids

	for i := 0; i < dy; i++ {
		row_max := 0
		for j := dx - 1; j >= 0; j-- {
			if trees[j][i] > row_max {
				vis[j][i] = 1
				row_max = trees[j][i]
			}
		}
	}
	return vis
}
func find_vis_right_to_left(trees [][]int, dy int, dx int, vis [][]int) [][]int {
	for i := 0; i < dy; i++ {
		row_max := 0
		for j := dx - 1; j >= 0; j-- {
			if trees[i][j] > row_max {
				vis[i][j] = 1
				row_max = trees[i][j]
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
	vis := make([][]int, dy)
	final_vis := make([][]int, dy)

	fmt.Println("gridsize:", dy, dx)

	for i := 0; i < dy; i++ {
		tree_grid[i] = make([]int, dx)
		vis[i] = make([]int, dx)
		final_vis[i] = make([]int, dx)

		row := splitInput[i]
		for j := 0; j < dx; j++ {
			height, _ := strconv.Atoi(string(row[j]))
			tree_grid[i][j] = height

			if i == 0 || i == dy-1 {
				vis[i][j] = 1

			}
			if j == 0 || j == dx-1 {
				vis[i][j] = 1

			}
		}
	}

	vis1 := find_vis_left_to_right(tree_grid, dy, dx, vis)
	vis2 := find_vis_right_to_left(tree_grid, dy, dx, vis)
	vis3 := find_vis_up_to_down(tree_grid, dy, dx, vis)
	vis4 := find_vis_down_to_up(tree_grid, dy, dx, vis)

	counter := 0
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			if vis1[i][j] == 1 || vis2[i][j] == 1 || vis3[i][j] == 1 || vis4[i][j] == 1 {
				counter += 1
			}
		}
	}

	fmt.Println("ans:", counter)
}
