package helper

import (
	"bufio"
	"os"
)

// LoadInputLines loads the input assuming each line is a string
func LoadInputLines(f string) (lines []string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
