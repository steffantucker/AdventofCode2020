package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type dirs struct {
	dir byte
	amp float64
}
type planedesc struct {
	angle float64
	north float64
	east  float64
}

func main() {
	nums := []dirs{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n := dirs{dir: line[0]}
		a, err := strconv.ParseFloat(string(line[1:]), 64)
		if err != nil {
			panic(err)
		}
		n.amp = a
		nums = append(nums, n)
	}
	plane := planedesc{angle: 0, north: 0, east: 0}
	end := followDirs(nums, plane)
	manhattan := getManhattanDist(end)
	fmt.Printf("plane: %#v\nman dist: %v\n", end, manhattan)
	waypoint := planedesc{east: 10, north: 1, angle: 0}
	wend := followWaypoint(nums, plane, waypoint)
	manhattan = getManhattanDist(wend)
	fmt.Printf("plane: %#v\nman: %v\n", wend, manhattan)
}

func followWaypoint(dirs []dirs, plane, waypoint planedesc) planedesc {
	for _, dir := range dirs {
		switch dir.dir {
		case 'N':
			waypoint.north += dir.amp
			break
		case 'S':
			waypoint.north -= dir.amp
			break
		case 'E':
			waypoint.east += dir.amp
			break
		case 'W':
			waypoint.east -= dir.amp
			break
		case 'L':
			angle := dir.amp * (math.Pi / 180)
			y, x := waypoint.north, waypoint.east
			dx := x*math.Cos(angle) - y*math.Sin(angle)
			dy := y*math.Cos(angle) + x*math.Sin(angle)
			sx := fmt.Sprintf("%.0f", dx)
			sy := fmt.Sprintf("%.0f", dy)
			ix, _ := strconv.Atoi(sx)
			iy, _ := strconv.Atoi(sy)
			waypoint.east = float64(ix)
			waypoint.north = float64(iy)
			fmt.Printf("angle: %v east: %v, north: %v\n", angle, waypoint.east, waypoint.north)
			break
		case 'R':
			angle := -dir.amp * (math.Pi / 180)
			y, x := waypoint.north, waypoint.east
			dx := x*math.Cos(angle) - y*math.Sin(angle)
			dy := y*math.Cos(angle) + x*math.Sin(angle)
			sx := fmt.Sprintf("%.0f", dx)
			sy := fmt.Sprintf("%.0f", dy)
			ix, _ := strconv.Atoi(sx)
			iy, _ := strconv.Atoi(sy)
			waypoint.east = float64(ix)
			waypoint.north = float64(iy)
			fmt.Printf("angle: %v east: %v, north: %v\n", angle, waypoint.east, waypoint.north)
			break
		case 'F':
			plane.east += dir.amp * waypoint.east
			plane.north += dir.amp * waypoint.north
			break
		}
		//fmt.Printf("plane: %#v\nway: %#v\ndirs: %#v\n\n", plane, waypoint, dir)
	}
	return plane
}

func followDirs(dirs []dirs, plane planedesc) planedesc {
	for _, dir := range dirs {
		switch dir.dir {
		case 'N':
			plane.north += dir.amp
			break
		case 'S':
			plane.north -= dir.amp
			break
		case 'E':
			plane.east += dir.amp
			break
		case 'W':
			plane.east -= dir.amp
			break
		case 'L':
			plane.angle += dir.amp
			break
		case 'R':
			plane.angle -= dir.amp
			break
		case 'F':
			dx := dir.amp * math.Cos(plane.angle*(math.Pi/180))
			dy := dir.amp * math.Sin(plane.angle*(math.Pi/180))
			plane.east += dx
			plane.north += dy
			break
		}
	}
	return plane
}

func getManhattanDist(plane planedesc) int {
	return int(math.Abs(float64(plane.north)) + math.Abs(float64(plane.east)))
}
