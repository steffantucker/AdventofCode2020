package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	number int
	pos    int
}

func main() {
	buses := []bus{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliest, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	line := scanner.Text()
	for i, v := range strings.Split(line, ",") {
		n, err := strconv.Atoi(v)
		if err == nil {
			buses = append(buses, bus{number: n, pos: i + 1})
		}
	}
	fmt.Printf("earliest: %v\n", findEarliest(buses, earliest))
	sum := findConsecutive(buses)
	fmt.Printf("sum: %v\n", sum)
}

// could work, but too slow
func findConsecutive(buses []bus) uint64 {
	fmt.Printf("buses: %#v\n", buses)
	var t uint64
	t = 200000000000000
	found := 0
	for ; found != len(buses); t++ {
		found = 0
		if t%1000000000 == 0 {
			fmt.Println(t)
		}
		for _, b := range buses {
			if (t+uint64(b.pos))%uint64(b.number) != 0 {
				break
			} else {
				found++
			}
		}
	}

	return t
}

func findEarliest(buses []bus, earliest int) int {
	shortest := earliest
	bus := 0
	for _, b := range buses {
		next := earliest / b.number
		next++
		next = b.number*next - earliest
		fmt.Printf("next: %v\n", next)
		if next < shortest {
			shortest = next
			bus = b.number
		}
	}
	return bus * shortest
}
