package main

import (
	"fmt"
	"strings"
	// "io/ioutil"
	"log"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")


	input, err := os.Open("example.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%s", input)

	// input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("example1.txt")
	// input, _ := ioutil.ReadFile("input.txt")

	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	for i := 0; i < len(splitInput); i++ {

		fmt.Println(splitInput[i])
	}
}
