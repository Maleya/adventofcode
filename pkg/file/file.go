package file

import (
	"io"
	"log"
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
