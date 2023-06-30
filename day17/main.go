package main

import (
	"fmt"
	"os"
)

func parse() string {
	file, _ := os.ReadFile("smallinput.txt")

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

type Rock struct {
	position Point // Top left position of the Rock
	shape    Shape
}

func canMove(dir []Point, rock Rock, grid [][]byte) bool {
	for i := 0; i < len(rock.shape.points); i++ {
		point := rock.shape.points[i]
		if grid[point.x][point.y] != 0 {
			return false
		}
	}

	return true
}

func fellRock(jets []byte, rock Rock, grid [][]byte) {
}

func main() {
	jets := parse()
	fmt.Println(jets)
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
	lastObstacle := 0

	grid := make([][]byte, 1)

	for i := 0; i < rounds; i++ {
		currRock := Rock{
			shape: shapes[0],
			position: Point{
				x: 2,
				y: lastObstacle + 3,
			},
		}
		fellRock([]byte(jets), currRock, grid)
	}

	// fmt.Println(lastObstacle)
}
