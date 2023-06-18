package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func visibility(grid [][]int, x int, y int) int {
	// go up
	here := grid[y][x]

	upvis := 0
	for i := y - 1; i >= 0; i-- {
		curr := grid[i][x]
		upvis += 1
		if curr >= here {
			break
		}
	}

	// go right
	rightvis := 0
	for i := x + 1; i < len(grid[0]); i++ {
		curr := grid[y][i]
		rightvis += 1
		if curr >= here {
			break
		}
	}
	// go down
	downvis := 0
	for i := y + 1; i < len(grid); i++ {
		curr := grid[i][x]
		downvis += 1
		if curr >= here {
			break
		}
	}
	// go left
	leftvis := 0
	for i := x - 1; i >= 0; i-- {
		curr := grid[y][i]
		leftvis += 1
		if curr >= here {
			break
		}
	}
	return upvis * rightvis * downvis * leftvis
}

func main() {
	// Look at the grid from all 4 directions and determine the visible trees
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, value := range line {
			intValue := int(value - '0')
			row = append(row, intValue)
		}
		grid = append(grid, row)
	}

	var min int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			// current := grid[x][y]
			vis := visibility(grid, x, y)
			if (x == 0 && y == 0) || vis > min {
				min = vis
			}
		}
	}
	fmt.Println(min)
}
