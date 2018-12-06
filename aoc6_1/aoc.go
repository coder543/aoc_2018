package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func getCoords() []coord {
	input, _ := ioutil.ReadFile("input")
	values := strings.Split(string(input), "\n")

	coords := []coord{}
	for _, val := range values {
		valxy := strings.Split(val, ", ")
		x, _ := strconv.ParseInt(valxy[0], 10, 32)
		y, _ := strconv.ParseInt(valxy[1], 10, 32)
		coords = append(coords, coord{int(x), int(y)})
	}

	return coords
}

func calcBox(coords []coord) (left, top, right, bottom int) {
	left, bottom = coords[0].x, coords[0].y
	for _, coord := range coords {
		if coord.x < left {
			left = coord.x
		}
		if coord.x > right {
			right = coord.x
		}
		if coord.y < bottom {
			bottom = coord.y
		}
		if coord.y > top {
			top = coord.y
		}
	}

	return
}

func makeGrid(coords []coord) [][]int {
	left, top, right, bottom := calcBox(coords)
	width := right - left + 1
	height := top - bottom + 1

	// recenter the coordinates in the new grid
	for i := range coords {
		coords[i].x -= left
		coords[i].y -= bottom
	}

	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	return grid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestCoord(x, y int, coords []coord) int {
	nearest := 0
	distance := 100000
	inbetween := false
	for i, coord := range coords {
		distX := abs(coord.x - x)
		distY := abs(coord.y - y)
		dist := distX + distY
		if dist == distance {
			inbetween = true
		} else if dist < distance {
			nearest = i
			distance = dist
			inbetween = false
		}
	}

	if inbetween {
		return -1
	}

	return nearest
}

func main() {
	coords := getCoords()
	grid := makeGrid(coords)
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = nearestCoord(x, y, coords)
		}
		// fmt.Println(grid[y])
	}

	coordArea := map[int]int{}
	for y := range grid {
		for x := range grid[y] {
			coord := grid[y][x]
			if coord == -1 {
				continue
			}
			_, ok := coordArea[coord]
			if !ok {
				coordArea[coord] = 0
			}
			coordArea[coord] += 1
		}
	}

	biggest := 0
	for k, v := range coordArea {
		if v > biggest {
			fmt.Println(k, ",", v)
			biggest = v
		}
	}
}
