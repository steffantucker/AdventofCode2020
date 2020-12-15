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

func main() {
	nums := make([]int, CYCLES)
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	for i, v := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(v)
		nums[i] = n
	}
	ans := runGame(nums, CYCLES)
	fmt.Printf("ans: %v\n", ans)
}

func runGame(nums []int, cycles int) int {
	for i := 5; i < cycles-1; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		//if i < 0 {
		ch := make(chan int)
		go find(nums, i, 0, i-1, ch)
		a := <-ch
		if a == -1 {
			nums[i+1] = 0
		} else {
			nums[i+1] = a
		}
		/*} else {
			ch := make(chan int)
			go find(nums, i, i/2+1, i-1, ch)
			go find(nums, i, 0, i/2, ch)
			a, b := <-ch, <-ch
			if a == -1 && b == -1 {
				nums[i+1] = 0
				continue
			}
			if a < b {
				nums[i+1] = b
			} else {
				nums[i+1] = a
			}
		}*/
	}
	return nums[len(nums)-1]
}

func find(nums []int, target, start, end int, ch chan int) {
	for i := end; i >= start; i-- {
		if nums[i] == nums[target] {
			ch <- target - i
			return
		}
	}
	ch <- -1
}
