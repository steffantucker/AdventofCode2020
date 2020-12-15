package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums := []int{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)
	difs := []int{}
	prev := 0
	for _, v := range nums {
		n := v - prev
		difs = append(difs, n)
		prev = v
	}
	ones := count(difs, 1)
	threes := count(difs, 3)
	fmt.Printf("ans: %v\n", ones*threes)
	poss := countPoss(nums)
	fmt.Printf("diffs: %#v\npossibilities: %v\n", difs, poss)
}

func count(n []int, t int) (count int) {
	for _, v := range n {
		if v == t {
			count++
		}
	}
	return
}

func countPoss(nums []int) int {
	prev := 0
	ocount := 0
	perm := []int{1, 1, 1, 2, 4, 7}
	total := 1
	for _, v := range nums {
		ocount++
		n := v - prev
		prev = v
		if n == 3 {
			total *= perm[ocount]
			ocount = 0
		}
	}
	return total
}
