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

type node struct {
	name  string
	left  string
	right string
}

func (n *node) getNodeName(s string) string {
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

type graph struct {
	nodes        map[string]node
	instructions string
	steps        int
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

		}
	}
}

func (g *graph) runInstructions(startNodeName string) string {
	endNodeName := "ZZZ"
	var newNodename string

	fmt.Println(startNodeName, endNodeName)
	n := g.nodes[startNodeName]

	for _, instr := range g.instructions {
		instr := string(instr)
		newNodename = n.getNodeName(string(instr))
		fmt.Println("At", n.name, "going", instr, "to", newNodename)
		if newNodename == endNodeName {
			g.steps++
			fmt.Println("found end node, after steps:", g.steps)
			return newNodename
		}
		n = g.nodes[newNodename]
		g.steps++

	}
	return newNodename
}

func (g *graph) loopInstructions() {
	g.steps = 0
	latest_node := "AAA"
	for latest_node != "ZZZ" {
		latest_node = g.runInstructions(latest_node)
	}
}

func main() {
	// load_file := example_input
	load_file := input
	splitInput := strings.Split(strings.TrimSpace(string(load_file)), "\n")

	g := graph{
		nodes: make(map[string]node),
		steps: 0,
	}
	g.process_input(splitInput)
	// g.runInstructions()
	g.loopInstructions()
	// fmt.Println(g)
}
