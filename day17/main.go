package main

import (
	"fmt"
	"os"
)

func parse() string {
	file, _ := os.ReadFile("input.txt")

	return string(file)
}

type Point struct {
	x int
	y int
}

type Shape struct {
	points []Point // Relative to position
	height int
}

func isObstacle(value byte) bool {
	if value == '@' || value == '#' {
		return true
	}
	return false
}

func isOutOfBounds(x int, y int, grid [][7]byte) bool {
	if len(grid) <= y || y < 0 || x < 0 || len(grid[0]) <= x {
		return true
	}
	return false
}

func canMove(dir Point, position Point, rock Shape, grid [][7]byte) bool {
	// TODO return real value
	for _, point := range rock.points {
		nextPoint := Point{
			x: position.x + dir.x + point.x,
			y: position.y + dir.y + point.y,
		}
		if isOutOfBounds(nextPoint.x, nextPoint.y, grid) || isObstacle(grid[nextPoint.y][nextPoint.x]) {
			return false
		}
	}
	return true
}

func rockefeller(jets []byte, startingPosition Point, rock Shape, grid [][7]byte, jetIdx int) (int, int) {
	position := startingPosition
	idx := jetIdx

	// xDir := -1
	jet := jets[idx/2%len(jets)]
	// if jet == '>' {
	// 	xDir = 1
	// }
	var dir Point
	// dir := Point{
	// 	x: xDir,
	// 	y: 0,
	// }

	for {
		if idx%2 == 0 {
			jet = jets[(idx/2)%len(jets)]
			xDir := -1
			if jet == '>' {
				xDir = 1
			}
			// Move with jet
			dir = Point{
				x: xDir,
				y: 0,
			}
		} else {
			// Move with falling
			dir = Point{
				x: 0,
				y: -1,
			}
		}
		moved := canMove(dir, position, rock, grid)

		idx++
		if !moved {
			if dir.x == 0 {
				// Need to settle at this point
				break
			}
		} else {
			// Update positions
			position = Point{
				x: position.x + dir.x,
				y: position.y + dir.y,
			}
		}
	}

	for _, point := range rock.points {
		grid[position.y+point.y][position.x+point.x] = '#'
	}

	// Return the index of the last obstacle

	return position.y + rock.height, idx
}

func printGrid(grid [][7]byte, p Point) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	jets := parse()
	rounds := 2022

	shapes := [5]Shape{
		{
			points: []Point{ // -
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 2, y: 0},
				{x: 3, y: 0},
			},
			height: 1,
		},
		{
			points: []Point{ // +
				{x: 1, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
				{x: 2, y: 1},
				{x: 1, y: 2},
			},
			height: 3,
		},
		{
			points: []Point{ // L reverse
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 2, y: 0},
				{x: 2, y: 1},
				{x: 2, y: 2},
			},
			height: 3,
		},
		{
			points: []Point{ // |
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 0, y: 2},
				{x: 0, y: 3},
			},
			height: 4,
		},
		{
			points: []Point{ // square
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
			},
			height: 2,
		},
	}

	grid := make([][7]byte, 5)
	// grid[0] = make([]byte, 7)
	lastObstacle := 0
	jetIdx := 0
	lastHeight := 0

	for i := 0; i < rounds; i++ {
		currRock := shapes[i%len(shapes)]
		startingPosition := Point{
			x: 2,
			y: lastObstacle + 3,
		}
		lastHeight, jetIdx = rockefeller([]byte(jets), startingPosition, currRock, grid, jetIdx)
		if lastHeight > lastObstacle {
			lastObstacle = lastHeight
		}
		for j := len(grid); j < lastObstacle+3+shapes[(i+1)%len(shapes)].height; j++ {
			grid = append(grid, [7]byte{})
		}

		printGrid(grid, startingPosition)
		fmt.Println(lastObstacle)
	}

	fmt.Println(lastObstacle)
}
