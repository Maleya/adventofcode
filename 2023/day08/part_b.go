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
	end_dist     map[int]int
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

// func (g *graph) runInstructions_old(startNodeName string) string {
// 	endNodeName := "ZZZ"
// 	var newNodename string

// 	fmt.Println(startNodeName, endNodeName)
// 	n := g.nodes[startNodeName]

// 	for _, instr := range g.instructions {
// 		instr := string(instr)
// 		newNodename = n.getNodeName(string(instr))
// 		fmt.Println("At", n.name, "going", instr, "to", newNodename)
// 		if newNodename == endNodeName {
// 			g.steps++
// 			fmt.Println("found end node, after steps:", g.steps)
// 			return newNodename
// 		}
// 		n = g.nodes[newNodename]
// 		g.steps++

// 	}
// 	return newNodename
// }

func (g *graph) runInstructions(startNodeNames []string) ([]string, bool) {
	n_starts := len(g.starterNodes)
	fmt.Println(n_starts, "starts")
	fmt.Println(startNodeNames)
	current_nodes := startNodeNames

	for _, instr := range g.instructions {
		var newNodes []string
		instr := string(instr)
		counter := 0

		// apply instruction to current nodes.
		for i, node_name := range current_nodes {
			new_name := g.nodes[node_name].getNodeName(instr)
			// fmt.Println("step:", g.steps, "At", node_name, "going", instr, "to", new_name, "counter:", counter)
			if strings.HasSuffix(new_name, "Z") {
				g.end_dist[i] = g.steps
				counter++
			}
			newNodes = append(newNodes, new_name)
			current_nodes = newNodes

		}
		g.steps++
		// fmt.Println(g.steps)
		if counter == n_starts {
			fmt.Println("found end node", g.steps)
			return current_nodes, true
			// break
		}
	}
	return current_nodes, false

}

func (g *graph) loopInstructions() {
	g.steps = 0
	latest_node := g.starterNodes
	ok := false
	for ok == false {
		// fmt.Println("looped instructions")
		latest_node, ok = g.runInstructions(latest_node)
		fmt.Println(g.end_dist)
	}
}

func main() {
	// fileName := "example.txt"
	// fileName := "example1.txt"
	// fileName := "example2.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")

	g := graph{
		nodes:    make(map[string]node),
		steps:    0,
		end_dist: make(map[int]int),
	}
	g.process_input(splitInput)
	// fmt.Println(g)

	// g.runInstructions(g.starterNodes)
	g.loopInstructions()
}

// lcm of the distance to the first z on each path
// assumed separate loops and no that take up full
