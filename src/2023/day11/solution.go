package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func containsOnly(search string, slice string) bool {
	l := len(slice)
	count := 0
	for _, element := range slice {
		if string(element) == search {
			count++
		}
	}
	return count == l
}
func countBelow(value int, slice []int) int {
	count := 0
	for _, element := range slice {
		if element < value {
			count++
		}
	}
	return count
}

type coords struct {
	y int
	x int
}

type galaxies struct {
	locations map[int]coords
	lastElem  coords
	emptyrows []int
	emptycols []int
}

func init_galaxy(input []string) *galaxies {
	loc := make(map[int]coords)
	var emptyrows []int
	var emptycols []int
	galaxyInCol := make(map[int]bool)
	counter := 0

	for y, line := range input {
		if containsOnly(".", line) {
			emptyrows = append(emptyrows, y)
		}
		for x, char := range line {
			if string(char) == "#" {
				loc[counter] = coords{y, x}
				counter++
				galaxyInCol[x] = true
			}
		}
	}
	// check for columns without galaxies:
	for y := 0; y < len(input[0]); y++ {
		if !galaxyInCol[y] {
			emptycols = append(emptycols, y)
		}
	}

	g := galaxies{
		locations: loc,
		lastElem:  coords{len(input), len(input[0])},
		emptyrows: emptyrows,
		emptycols: emptycols,
	}
	return &g
}

func (g *galaxies) expandSpace(factor int) {
	new_map := make(map[int]coords)

	for i, gal := range g.locations {
		dx := countBelow(gal.x, g.emptycols)
		dy := countBelow(gal.y, g.emptyrows)
		new_gal := coords{gal.y - dy + dy*factor, gal.x - dx + dx*factor}
		new_map[i] = new_gal

	}
	g.locations = new_map
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func distance(p1 coords, p2 coords) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
}

func (g *galaxies) uniquedistanceSum() int {
	sum := 0

	keys := make([]int, 0, len(g.locations))
	for key := range g.locations {
		keys = append(keys, key)
	}
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			key1 := keys[i]
			key2 := keys[j]

			gal1 := g.locations[key1]
			gal2 := g.locations[key2]

			sum += distance(gal1, gal2)
		}
	}
	return sum
}

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

func partA(input []string) {
	g := init_galaxy(input)
	g.expandSpace(2)
	sum := g.uniquedistanceSum()
	fmt.Println("distance sum part a:", sum)
}
func partB(input []string) {
	g := init_galaxy(input)
	g.expandSpace(1000000)
	sum := g.uniquedistanceSum()
	fmt.Println("distance sum part b", sum)
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	partA(splitInput)
	partB(splitInput)
}
