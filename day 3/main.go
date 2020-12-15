package main

import (
	"bufio"
	"fmt"
	"os"
)

type movement struct {
	right int
	down  int
}

func main() {
	var in []string
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		in = append(in, line)
	}

	move := []movement{
		{
			right: 1,
			down:  1,
		},
		{
			right: 3,
			down:  1,
		},
		{
			right: 5,
			down:  1,
		},
		{
			right: 7,
			down:  1,
		},
		{
			right: 1,
			down:  2,
		},
	}

	total := 1
	for _, m := range move {
		total *= slope(m, in)
	}
	fmt.Println(total)
}

func slope(m movement, in []string) int {
	var x, y int
	width := len(in[0])
	length := len(in)
	treecount := 0
	for y = m.down; y < length; y += m.down {
		x += m.right
		if rune(in[y][x%width]) == '#' {
			treecount++
		}
	}
	fmt.Printf("right: %v down: %v total: %v\n", m.right, m.down, treecount)
	return treecount
}
