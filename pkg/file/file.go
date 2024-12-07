package file

import (
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
)

func ReadInput(fileName string) (content []string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		// return content, err
	}
	defer file.Close()
	lineContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		// return content, err
	}
	content = strings.Split(strings.TrimSpace(string(lineContent)), "\n")
	return content
}

func NewInputReader() *os.File {
	inputReader, err := os.Open("input.txt")

	if err != nil {
		slog.Error("got an error while trying to open input file", "Error", err)
	}

	return inputReader
}
