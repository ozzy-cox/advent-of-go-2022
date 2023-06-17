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

	// | | | | Looking at the grid this way
	// v v v v
	//

	visibleIndices := make(map[string]bool)

	for x := 0; x < len(grid[0]); x++ {
		columnMax := grid[0][x]

		visibleIndices[fmt.Sprintf("%d-%d", x, 0)] = true
		fmt.Printf("%d-%d\n", x, 0)
		for y := 0; y < len(grid); y++ {
			currentPoint := grid[y][x]
			if currentPoint > columnMax {
				columnMax = currentPoint
				visibleIndices[fmt.Sprintf("%d-%d", x, y)] = true
				fmt.Printf("%d-%d\n", x, y)
			}
		}
	}
	fmt.Println(visibleIndices)

	// -> Looking at the grid this way
	// ->
	// ->
	// ->
	for y := 0; y < len(grid); y++ {
		rowMax := grid[y][0]

		visibleIndices[fmt.Sprintf("%d-%d", 0, y)] = true
		fmt.Printf("%d-%d\n", 0, y)
		for x := 0; x < len(grid[0]); x++ {
			currentPoint := grid[y][x]
			if currentPoint > rowMax {
				rowMax = currentPoint
				visibleIndices[fmt.Sprintf("%d-%d", x, y)] = true
				fmt.Printf("%d-%d\n", x, y)
			}
		}
	}
	fmt.Println(visibleIndices)

	// ^ ^ ^ ^
	// | | | | Looking at the grid this way
	//

	for x := 0; x < len(grid[0]); x++ {
		columnMax := grid[len(grid)-1][x]

		visibleIndices[fmt.Sprintf("%d-%d", x, len(grid)-1)] = true
		fmt.Printf("%d-%d\n", x, len(grid)-1)
		for y := len(grid) - 1; y >= 0; y-- {
			currentPoint := grid[y][x]
			if currentPoint > columnMax {
				columnMax = currentPoint
				visibleIndices[fmt.Sprintf("%d-%d", x, y)] = true
				fmt.Printf("%d-%d\n", x, y)
			}
		}
	}

	// <- Looking at the grid this way
	// <-
	// <-
	// <-
	for y := 0; y < len(grid); y++ {
		rowMax := grid[y][len(grid[0])-1]

		visibleIndices[fmt.Sprintf("%d-%d", len(grid[0])-1, y)] = true
		fmt.Printf("%d-%d\n", len(grid[0])-1, y)
		for x := len(grid[0]) - 1; x >= 0; x-- {
			currentPoint := grid[y][x]
			if currentPoint > rowMax {
				rowMax = currentPoint
				fmt.Printf("%d-%d\n", x, y)
				visibleIndices[fmt.Sprintf("%d-%d", x, y)] = true
			}
		}
	}
	fmt.Println(len(visibleIndices))
}
