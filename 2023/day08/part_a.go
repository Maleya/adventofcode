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
	nodes        map[string]node
	instructions string
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

func (g *graph) traverse(instructions string) {
}

func main() {
	fileName := "example.txt"
	// fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")

	g := graph{
		nodes: make(map[string]node),
	}
	g.process_input(splitInput)
	fmt.Println(g)
}
