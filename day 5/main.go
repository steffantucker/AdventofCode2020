package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	splitreg := regexp.MustCompile(`((?:F|B){7})((?:L|R){3})`)
	biggest := -1
	var sids []int

	for scanner.Scan() {
		line := scanner.Text()
		directions := splitreg.FindStringSubmatch(line)
		drow := directions[1]
		dseat := directions[2]

		rownum := findRow(drow)
		seatnum := findSeat(dseat)
		//fmt.Printf("rownum: %#v seat: %v\n", rownum, seatnum)

		sid := rownum*8 + seatnum
		//fmt.Printf("sid: %v\n", sid)
		sids = append(sids, sid)
		if sid > biggest {
			biggest = sid
		}

	}
	sort.Ints(sids)
	fmt.Printf("list: %#v\n", sids)
	//fmt.Printf("biggest: %v\n", biggest)
}

func findRow(dirs string) int {
	top := 128
	bottom := 1
	for _, dir := range dirs {
		if dir == 'F' {
			top -= (top - bottom + 1) / 2
		}
		if dir == 'B' {
			bottom += (top - bottom + 1) / 2
		}
		//fmt.Printf("in: %#v top: %v bottom: %v\n", dir, top, bottom)
	}
	return top - 1
}

func findSeat(dirs string) int {
	top := 8
	bottom := 1
	for _, dir := range dirs {
		if dir == 'L' {
			top -= (top - bottom + 1) / 2
		}
		if dir == 'R' {
			bottom += (top - bottom + 1) / 2
		}
		//fmt.Printf("in: %#v top: %v bottom: %v\n", dir, top, bottom)
	}
	return top - 1
}
