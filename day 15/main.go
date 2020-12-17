package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CYCLES is num cycles to run
const CYCLES = 30000000
const START = 3

func main() {
	nums := make(map[int]int)
	file, err := os.Open("inputtest")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	start := 0
	for i, v := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(v)
		nums[n] = i
		start = n
	}
	delete(nums, start)
	ans := runGame(nums, CYCLES, start)
	fmt.Printf("ans: %v\n", ans)
}

func runGame(nums map[int]int, cycles, start int) int {
	previous, current := start, 0
	for i := START; i < cycles; i++ {
		if j, ok := nums[current]; ok {
			nums[previous] = i - 1
			nums[current] = i
			previous = current
			current = i - j
		} else {
			nums[previous] = i - 1
			nums[current] = i
			previous = current
			current = 0
		}
	}
	return previous
}
