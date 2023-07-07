package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

func parse() *map[Point]bool {
	points := make(map[Point]bool)
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		indices := [3]int{}
		for i, token := range tokens {
			indices[i], _ = strconv.Atoi(token)
		}
		points[Point{
			x: indices[0],
			y: indices[1],
			z: indices[2],
		}] = true
	}

	return &points
}

func main() {
	boulder := parse()
	fmt.Println(boulder)
	directions := [6]Point{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}

	totalSides := 0
	for point := range *boulder {
		for _, direction := range directions {
			if _, ok := (*boulder)[Point{point.x + direction.x, point.y + direction.y, point.z + direction.z}]; !ok {
				totalSides++
			}
		}
	}
	fmt.Println(totalSides)
}
