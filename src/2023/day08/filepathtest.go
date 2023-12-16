package main

import (
	"fmt"
	"os"
	"path/filepath"
	// "path/filepath"
)

func main() {
	// fileName := "example.txt"
	// fileName := "example_1.txt"
	// fileName := "src/2023/day07/input.txt"
	// fileName_abs, _ := filepath.Abs(fileName)

	// fmt.Println(fileName)
	// fmt.Println(fileName_abs)
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	ex, _ := os.Executable()
	fmt.Println(ex)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	filePath := filepath.Join(exPath, "example.txt")

	fmt.Println(exPath)

}
