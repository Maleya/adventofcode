package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example_input string

//go:embed example2.txt
var example_input2 string

type node struct {
	name  string
	left  string
	right string
}

func (n node) getNodeName(s string) string {
	switch strings.ToLower(s) {
	case "l":
		return n.left
	case "r":
		return n.right
	default:
		fmt.Println("invalid string", s)
		panic(1)
	}
}
func (n *node) nameEndson(s string) bool {
	return strings.HasSuffix(n.name, s)
}

type graph struct {
	nodes        map[string]node
	instructions string
	steps        int
	starterNodes []string
	dist_to_Z    map[int]int
}

func (g *graph) addNode(n node) {
	g.nodes[n.name] = n
}

func (g *graph) process_input(input []string) {
	for i, line := range input {
		if i == 0 {
			g.instructions = strings.TrimSpace(line)
		}

		parts := strings.Split(strings.TrimSpace(line), "=")

		if len(parts) == 2 {
			nodeName := strings.TrimSpace(parts[0])

			leftright := strings.TrimSpace(parts[1])

			trimmed := string(strings.Trim(leftright, "()"))
			split := strings.Split(strings.TrimSpace(trimmed), ",")

			left := strings.TrimSpace(split[0])
			right := strings.TrimSpace(split[1])

			n := node{nodeName, left, right}
			g.addNode(n)
			if n.nameEndson("A") {
				g.starterNodes = append(g.starterNodes, nodeName)

			}
		}
	}
}

func (g *graph) runtillZ() {

	fmt.Println(len(g.starterNodes))
	for _, nodename := range g.starterNodes {
		counter := 0
		counter_2 := 0
		i := 0
		fmt.Println(nodename)
		for counter_2 < len(g.starterNodes) {
			idx := i % len(g.instructions)
			instruction := string(g.instructions[idx])
			nodename = g.nodes[nodename].getNodeName(instruction)
			i++
			counter++

			if strings.HasSuffix(nodename, "Z") {
				fmt.Println("reached", nodename, "in:", counter)
				counter_2++
				break
			}
		}
	}
}

func main() {
	// load_file := example_input
	// load_file := example_input2
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	g := graph{
		nodes:     make(map[string]node),
		steps:     0,
		dist_to_Z: make(map[int]int),
	}
	g.process_input(splitInput)
	// fmt.Println(g)

	// g.runInstructions(g.starterNodes)
	// g.loopInstructions()
	g.runtillZ()
}

// lcm of the distance to the first z on each path
// assumed separate loops and no that take up full
