package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answers := make(map[string]int)
	total := 0
	group := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += checkAnswers(answers, group)
			answers = make(map[string]int)
			group = 0
		} else {
			answers = getAnswers(answers, line)
			group++
		}
	}
	fmt.Printf("total: %v\n", total)
}

func checkAnswers(a map[string]int, group int) int {
	count := 0
	for _, v := range a {
		if v == group {
			count++
		}
	}
	return count
}

func getAnswers(a map[string]int, in string) map[string]int {
	for _, v := range in {
		a[string(v)]++
	}
	return a
}
