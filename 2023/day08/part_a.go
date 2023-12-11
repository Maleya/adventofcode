package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type node struct {
	name  string
	left  string
	right string
}

func (n *node) run(s string) string {
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
	nodes []node
}

func (g *graph) addNode(n node) {
	g.nodes = append(g.nodes, n)
}

func (g *graph) process_input(input []string) {
	for _, line := range input {

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

func (g *graph) traverse(instructions string) {
}

func main() {
	fileName := "example.txt"
	// fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")

	var g graph
	g.process_input(splitInput)
	fmt.Println(g)
}
