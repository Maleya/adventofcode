package main

import (
	"adventofcode/pkg/enums/cardinaldirection"
	_ "embed"
	"fmt"
)

func main() {

	// a := North
	a := cardinaldirection.North
	fmt.Println(a, a.Add(2), a.Opposite(), a.Add(60))

}
