package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type memorymap map[int]int64

func main() {
	memory := make(memorymap)
	memV2 := make(memorymap)
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mask := map[int]bool{}
	maskV2 := map[int]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mask") {
			mask = getMask(line[7:])
			maskV2 = getMaskV2(line[7:])
		} else if strings.Contains(line, "mem") {
			memory.updateMemory(line, mask)
			memV2.updateMemoryV2(line, maskV2)
		}
	}
	sum := memory.sum()
	sum2 := memV2.sum()
	fmt.Printf("sum: %v sumV2: %v\n", sum, sum2)
}

func getMaskV2(line string) (mask map[int]rune) {
	mask = map[int]rune{}
	for i, v := range line {
		mask[i] = v
	}
	return
}

func getMask(line string) (mask map[int]bool) {
	mask = map[int]bool{}
	for i, v := range line {
		switch v {
		case '1':
			mask[i] = true
			break
		case '0':
			mask[i] = false
			break
		}
	}
	return
}

func (m *memorymap) updateMemoryV2(line string, mask map[int]rune) {
	mreg := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	input := mreg.FindStringSubmatch(line)
	address, _ := strconv.Atoi(input[1])
	value, _ := strconv.ParseInt(input[2], 10, 64)
	xs := map[int]bool{}
	for i, v := range mask {
		switch v {
		case '1':
			address |= 1 << (35 - i)
			break
		case 'X':
			xs[i] = true
			break
		}
	}
	newAddresses := applyXs(xs, address, value)
	for i, v := range newAddresses {
		(*m)[i] = v
	}
}

func applyXs(xs map[int]bool, address int, value int64) (naddresses memorymap) {
	naddresses = make(memorymap)
	add1, add2 := address, address
	for i := range xs {
		add1 |= 1 << (35 - i)
		add2 &= ^(1 << (35 - i))
		delete(xs, i)
		xs2 := make(map[int]bool)
		for j, v := range xs {
			xs2[j] = v
		}
		nadd1 := applyXs(xs, add1, value)
		nadd2 := applyXs(xs2, add2, value)
		for j, v := range nadd1 {
			naddresses[j] = v
		}
		for j, v := range nadd2 {
			naddresses[j] = v
		}
		break
	}
	naddresses[add1] = value
	naddresses[add2] = value
	return
}

func (m *memorymap) updateMemory(line string, mask map[int]bool) {
	mreg := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	input := mreg.FindStringSubmatch(line)
	address, _ := strconv.Atoi(input[1])
	value, _ := strconv.ParseInt(input[2], 10, 64)
	for i, v := range mask {
		if v {
			value |= 1 << (35 - i)
		} else if !v {
			value &= ^(1 << (35 - i))
		}
	}
	(*m)[address] = value
}

func (m *memorymap) sum() (sum int64) {
	for _, v := range *m {
		sum += v
	}
	return
}
