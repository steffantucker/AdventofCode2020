package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const PREAMBLE = 25

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
	missing := findMissing(nums)
	fmt.Printf("missing: %v\n", missing)
	fmt.Printf("weakness: %v\n", findWeakness(nums, missing))
}

func findWeakness(nums []int, missing int) (weak int) {
	var contiguous []int
	for size := 2; size < len(nums); size++ {
		for i := 0; i < len(nums)-size; i++ {
			contiguous = nums[i : i+size]
			s := sum(contiguous)
			if s == missing {
				small, big := minMax(contiguous)
				return small + big
			}
		}
	}
	return -1
}

func minMax(nums []int) (min, max int) {
	sort.Ints(nums)
	return nums[0], nums[len(nums)-1]
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func findMissing(nums []int) int {
	for i := PREAMBLE; i < len(nums); i++ {
		match := false
		for j := i - PREAMBLE; j <= i; j++ {
			for k := j + 1; k < i; k++ {
				if nums[i] == nums[j]+nums[k] {
					match = true
					break
				}
			}
			if match {
				break
			}
		}
		if !match {
			return nums[i]
		}
	}
	return -1
}
