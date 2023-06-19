package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	Point [2]int
	Queue []Point
	Grid  [][]byte
)

func (p *Point) x() int {
	return p[0]
}

func (p *Point) y() int {
	return p[1]
}

func (q *Queue) Push(b Point) {
	*q = append(*q, b)
}

func (q *Queue) Pop() Point {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue) Peek() Point {
	return (*q)[0]
}

var dirs = [...]Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getReachableNeighbors(p Point, g Grid) []Point {
	neighbors := make([]Point, 0)
	for _, dir := range dirs {
		neighbor := Point{p.x() + dir.x(), p.y() + dir.y()}
		neighborIsInsideGrid := neighbor.x() >= 0 && neighbor.x() < len(g) && neighbor.y() >= 0 && neighbor.y() < len(g[0])
		if neighborIsInsideGrid {
			neighborValue := g[neighbor.x()][neighbor.y()]
			currValue := g[p.x()][p.y()]
			neighborIsReachable := currValue-1 <= neighborValue
			if neighborIsReachable {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func main() {
	visited := make(map[string]bool)

	distances := make([][]int, 0)
	grid := make(Grid, 0)

	// Read map
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	q := make(Queue, 0)

	lineCount := 0

	// var end Point
	for scanner.Scan() {
		line := scanner.Text()

		distances = append(distances, make([]int, len(line)))
		grid = append(grid, make([]byte, len(line)))
		for i, char := range line {
			if char == 'E' {
				visited[fmt.Sprintf("%d:%d", lineCount, i)] = true
				distances[lineCount][i] = 0
				q.Push(Point{lineCount, i})
				grid[lineCount][i] = byte('z')
			} else {
				if char == 'S' {
					grid[lineCount][i] = byte('a')
					distances[lineCount][i] = -1
					// end = Point{lineCount, i}
				} else {
					grid[lineCount][i] = byte(char)
					distances[lineCount][i] = -1
				}
			}
		}
		lineCount++
	}

	for len(q) > 0 {
		curr := q.Pop()

		currDistance := distances[curr.x()][curr.y()]

		neighbors := getReachableNeighbors(curr, grid)

		for _, neighbor := range neighbors {
			if _, ok := visited[fmt.Sprintf("%d:%d", neighbor.x(), neighbor.y())]; !ok {
				distanceToNeighbor := distances[neighbor.x()][neighbor.y()]
				if distanceToNeighbor == -1 || currDistance+1 < distanceToNeighbor {
					distances[neighbor.x()][neighbor.y()] = currDistance + 1
					q.Push(neighbor)
				}

				currValue := grid[curr.x()][curr.y()]
				if currValue == 'a' {
					fmt.Println(distances[neighbor.x()][neighbor.y()] + 1)
					return
				}
			}
		}
	}
}
