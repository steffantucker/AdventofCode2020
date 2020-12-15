package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type op struct {
	ins string
	pos bool
	off int
}

func main() {
	instructions := []op{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	insreg := regexp.MustCompile(`(.{3}) (\+|-)(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		ins := insreg.FindStringSubmatch(line)
		offset, err := strconv.Atoi(ins[3])
		s := op{
			ins: ins[1],
			off: offset,
		}
		if err != nil {
			panic(err)
		}
		if ins[2] == "+" {
			s.pos = true
		}
		instructions = append(instructions, s)
	}
	value, finished := runInfiniteProgram(instructions, true)
	fmt.Printf("last acc: %v finished: %v\n", value, finished)
}

func runInfiniteProgram(instructions []op, firstrun bool) (acc int, finished bool) {
	runIns := map[int]bool{}
	for opnum := 0; opnum < len(instructions); {
		fmt.Printf("opnum: %v\n", runIns[opnum])
		ins := instructions[opnum]
		if runIns[opnum] {
			fmt.Println("infinite after change")
			return acc, false
		}
		runIns[opnum] = true
		switch ins.ins {
		case "acc":
			if ins.pos {
				acc += ins.off
			} else {
				acc -= ins.off
			}
			opnum++
			break
		case "jmp":
			if firstrun {
				fmt.Printf("attempt jmp->nop opnum: %v\n", opnum)
				tinstructions := make([]op, len(instructions))
				copy(tinstructions, instructions)
				tinstructions[opnum].ins = "nop"
				tacc, tfinish := runInfiniteProgram(tinstructions, false)
				if tfinish {
					return tacc, true
				}
			}
			if ins.pos {
				opnum += ins.off
			} else {
				opnum -= ins.off
			}
			break
		case "nop":
			if firstrun && ins.off > 0 {
				fmt.Printf("attempt nop->jmp opnum: %v\n", opnum)
				tinstructions := make([]op, len(instructions))
				copy(tinstructions, instructions)
				tinstructions[opnum].ins = "jmp"
				tacc, tfinish := runInfiniteProgram(tinstructions, false)
				if tfinish {
					return tacc, true
				}
			}
			opnum++
			break
		default:
			break
		}
	}
	return acc, true
}
