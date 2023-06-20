package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point [2]int

func (p Point) x() int {
	return p[0]
}

func (p Point) y() int {
	return p[1]
}

var obstacles = map[byte]bool{
	'#': true,
	'o': true,
}

func getNextPos(p Point, grid [][]byte, abyss int) (Point, error) {
	y := p.y() + 1
	if y == abyss+2 {
		return Point{}, errors.New("no paths, should settle")
	}
	if _, ok := obstacles[grid[y][p.x()]]; !ok {
		next, err := getNextPos(Point{p.x(), y}, grid, abyss)
		if err != nil {
			return Point{p.x(), y}, nil
		} else {
			return next, nil
		}
	} else if _, ok := obstacles[grid[y][p.x()-1]]; !ok {
		next, err := getNextPos(Point{p.x() - 1, y}, grid, abyss)
		if err != nil {
			return Point{p.x() - 1, y}, nil
		} else {
			return next, nil
		}
	} else if _, ok := obstacles[grid[y][p.x()+1]]; !ok {
		next, err := getNextPos(Point{p.x() + 1, y}, grid, abyss)
		if err != nil {
			return Point{p.x() + 1, y}, nil
		} else {
			return next, nil
		}
	} else {
		return Point{}, errors.New("no paths, should settle")
	}
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	grid := make([][]byte, 500)

	for i := range grid {
		grid[i] = make([]byte, 1000)
	}

	startingPoint := Point{500, 0}
	abyss := 0
	for scanner.Scan() {
		line := scanner.Text()

		points := strings.Split(line, " -> ")

		for i := 0; i < len(points)-1; i++ {
			point1 := points[i]
			point2 := points[i+1]

			pointParts1 := strings.Split(point1, ",")
			pointParts2 := strings.Split(point2, ",")

			x1 := pointParts1[0]
			y1 := pointParts1[1]

			x2 := pointParts2[0]
			y2 := pointParts2[1]

			if x1 == x2 {
				x, _ := strconv.Atoi(x1)
				y1, err1 := strconv.Atoi(y1)
				y2, err2 := strconv.Atoi(y2)
				if err1 != nil || err2 != nil {
					break
				}
				incr := 0
				if y2 > y1 {
					incr = +1
					if abyss < y2 {
						abyss = y2
					}
				} else {
					if abyss < y1 {
						abyss = y1
					}
					incr = -1
				}
				for i := y1; ; i += incr {
					grid[i][x] = '#'
					if i == y2 {
						break
					}
				}
			} else if y1 == y2 {
				y, _ := strconv.Atoi(y1)
				x1, err1 := strconv.Atoi(x1)
				x2, err2 := strconv.Atoi(x2)
				if err1 != nil || err2 != nil {
					break
				}
				incr := 0
				if x2 > x1 {
					incr = +1
				} else {
					incr = -1
				}
				for i := x1; ; i += incr {
					grid[y][i] = '#'
					if i == x2 {
						break
					}
				}
				if abyss < y {
					abyss = y
				}
			}
		}
	}
	// Add floor
	grid = append(grid, make([]byte, 1000))
	grid = append(grid, make([]byte, 1000))
	for i := range grid[len(grid)-1] {
		grid[len(grid)-1][i] = '#'
	}
	settled := 0
	var currPos Point
	currPos = startingPoint

	for {

		nextPos, err := getNextPos(currPos, grid, abyss)
		grid[currPos.y()][currPos.x()] = byte('o')
		if err != nil {
			if err.Error() == "its over" {
				break
			}
			grid[currPos.y()][currPos.x()] = byte('o')
			settled++
			if currPos == startingPoint {
				break
			}
			currPos = startingPoint
			continue
		}

		grid[currPos.y()][currPos.x()] = 0
		currPos = nextPos
	}
	fmt.Println(settled)
}
