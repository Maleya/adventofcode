package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	Name string
	Size int
}

type Directory struct {
	Name       string
	Size       int
	Parent     *Directory
	ChildDirs  map[string]*Directory
	ChildFiles []file
}

func makeDirectory(name string, parentDir *Directory) *Directory {
	return &Directory{Name: name, Parent: parentDir, ChildDirs: map[string]*Directory{}}
}

func (d *Directory) add_file(file_string string) {
	var size int
	var name string
	string_parts := strings.Split(file_string, " ")
	size, _ = strconv.Atoi(string_parts[0])
	name = string_parts[1]
	f := file{Name: name, Size: size}
	d.ChildFiles = append(d.ChildFiles, f)
}

func (d *Directory) FindDirectorySize() {
	for _, file := range d.ChildFiles {
		d.Size += file.Size
	}
	for _, Child_Dir := range d.ChildDirs {
		Child_Dir.FindDirectorySize()
		d.Size += Child_Dir.Size

	}
}

func findDirectoriesOfAtMost(dir *Directory, n int) []*Directory {
	var directories []*Directory
	if dir.Size <= n {
		directories = append(directories, dir)
	}
	for _, subdir := range dir.ChildDirs {
		directories = append(directories, findDirectoriesOfAtMost(subdir, n)...)
	}
	return directories
}

func list_sizes(dir *Directory) []int {
	var sizes []int
	sizes = append(sizes, dir.Size)
	for _, subdir := range dir.ChildDirs {
		// fmt.Println(dir.Size)
		sizes = append(sizes, list_sizes(subdir)...)

	}
	return sizes
}

func main() {

	// input, _ := ioutil.ReadFile("example.txt")
	input, _ := ioutil.ReadFile("input.txt")
	splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	root := makeDirectory("root", nil)
	var currentDir *Directory

	for i := 0; i < len(splitInput); i++ {
		// catch commands
		var command, arg, p1, p2 string

		fmt.Sscanf(splitInput[i], "$ %s %s", &command, &arg)

		if command == "cd" {
			if arg == "/" {
				currentDir = root

			} else if arg == ".." {
				currentDir = currentDir.Parent

			} else {

				currentDir = currentDir.ChildDirs[arg]
			}

		} else if command == "ls" {
			continue
		} else {
			// its not a command:
			fmt.Sscanf(splitInput[i], "%s %s", &p1, &p2)
			if p1 == "dir" {
				// its a dir
				dir := makeDirectory(p2, currentDir)
				currentDir.ChildDirs[dir.Name] = dir
			} else {
				currentDir.add_file(p1 + " " + p2)

			}
		}
	}
	root.FindDirectorySize()
	total_size := 0
	small_dirs := findDirectoriesOfAtMost(root, 100000)

	for _, dir := range small_dirs {
		total_size += dir.Size
	}
	fmt.Println("ans", total_size)

	// part 2:
	needed_space := 30000000 - (70000000 - root.Size)

	ans2 := list_sizes(root)
	sort.Ints(ans2)
	for _, size := range ans2 {
		if size >= needed_space {
			fmt.Println("smallest dir:", size)
			break
		}
	}
	// fmt.Println(ans2)

}
