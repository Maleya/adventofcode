package file

import (
	"io"
	"os"
	"strings"
)

func ReadInput(fileName string) (content []string, err error) {

	file, err := os.Open(fileName)
	if err != nil {
		return content, err
	}
	defer file.Close()
	lineContent, err := io.ReadAll(file)
	if err != nil {
		return content, err
	}
	content = strings.Split(strings.TrimSpace(string(lineContent)), "\n")
	return content, err

}
