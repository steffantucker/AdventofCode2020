package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	world := []string{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			panic(err)
		}
		world = append(world, line)
	}
	oldWorld := make([]string, len(world))
	copy(oldWorld, world)
	world = runSim2(oldWorld)
	simCount := 0
	for !compareWorlds(oldWorld, world) {
		fmt.Printf("sim: %v\n", simCount)
		copy(oldWorld, world)
		world = runSim2(oldWorld)
		simCount++
	}
	filledSeats := countSeats(world)
	fmt.Printf("seats: %#v\ncount: %v\n", world, filledSeats)
	/*once := runSim2(world)
	twice := runSim2(once)
	fmt.Printf("once:%#v\ntwice:%#v\n", once, twice)*/
}

func compareWorlds(x, y []string) bool {
	if len(x) != len(y) {
		panic("different world sizes")
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func runSim2(w []string) (world []string) {
	for i, y := range w {
		newRow := []byte{}
		for j, x := range y {
			if i == 0 && j == 3 {
				fmt.Println("here")
			}
			if x == '.' {
				newRow = append(newRow, '.')
				continue
			}
			cstr := []byte{}
			lr := findSeats([]byte(y), j)
			cstr = append(cstr, lr...)
			d := []byte{}
			for u := 0; u < len(w); u++ {
				d = append(d, w[u][j])
			}
			ud := findSeats(d, i)
			cstr = append(cstr, ud...)

			diag := []byte{}
			dloc := -1
			for dx, dy, dcount := j-i, i-j, 0; dx < len(y) && dy < len(w); dy, dx, dcount = dy+1, dx+1, dcount+1 {
				if dx < 0 {
					dx = 0
				}
				if dy < 0 {
					dy = 0
				}
				if dloc == -1 && dx == j && dy == i {
					dloc = dcount
				}
				diag = append(diag, w[dy][dx])
			}
			di := findSeats(diag, dloc)
			cstr = append(cstr, di...)

			diag2 := []byte{}
			dloc = -1
			dx := i + j
			dy := dx - (len(y) - 1)
			if dx >= len(y) {
				dx = len(y) - 1
			}
			if dy < 0 {
				dy = 0
			}
			for dcount := 0; dx >= 0 && dy < len(w); dy, dx, dcount = dy+1, dx-1, dcount+1 {
				if dloc == -1 && dx == j && dy == i {
					dloc = dcount
				}
				diag2 = append(diag2, w[dy][dx])
			}
			di2 := findSeats(diag2, dloc)
			cstr = append(cstr, di2...)
			newRow = append(newRow, checkArea(string(cstr), x))
		}
		world = append(world, string(newRow))
	}
	return
}

func runSim(w []string) (world []string) {
	for i, y := range w {
		newRow := []byte{}
		for j, x := range y {
			if x == '.' {
				newRow = append(newRow, '.')
				continue
			}
			cstr := []byte{}
			// top-bottom internal
			if i-1 >= 0 {
				cstr = append(cstr, w[i-1][j])
			}
			if i+1 < len(w) {
				cstr = append(cstr, w[i+1][j])
			}
			if j-1 >= 0 {
				cstr = append(cstr, w[i][j-1])
			}
			if j+1 < len(y) {
				cstr = append(cstr, w[i][j+1])
			}
			if i-1 >= 0 && j-1 >= 0 {
				cstr = append(cstr, w[i-1][j-1])
			}
			if i-1 >= 0 && j+1 < len(y) {
				cstr = append(cstr, w[i-1][j+1])
			}
			if i+1 < len(w) && j-1 >= 0 {
				cstr = append(cstr, w[i+1][j-1])
			}
			if i+1 < len(w) && j+1 < len(y) {
				cstr = append(cstr, w[i+1][j+1])
			}
			newRow = append(newRow, checkArea(string(cstr), x))
		}
		world = append(world, string(newRow))
	}
	return
}

func findSeats(seats []byte, l int) (s []byte) {
	for i := l - 1; i >= 0; i-- {
		if seats[i] == '#' || seats[i] == 'L' {
			s = append(s, seats[i])
			break
		}
	}
	for i := l + 1; i < len(seats); i++ {
		if seats[i] == '#' || seats[i] == 'L' {
			s = append(s, seats[i])
			break
		}
	}
	return
}

func checkArea(cstr string, x rune) byte {
	count := strings.Count(cstr, "#")
	if x == 'L' && count == 0 {
		return '#'
	} else if x == 'L' {
		return 'L'
	}
	if x == '#' && count < 5 {
		return '#'
	}
	return 'L'

}

func countSeats(w []string) (seats int) {
	for _, i := range w {
		seats += strings.Count(i, "#")
	}
	return
}
