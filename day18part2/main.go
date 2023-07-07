package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Point        [3]int
	Queue[T any] []T
)

func (q *Queue[T]) Push(b T) {
	*q = append(*q, b)
}

func (q *Queue[T]) Pop() T {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
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
		points[indices] = true
	}

	return &points
}

func canGoToNextEdge(curr Point, dir Point, boulder map[Point]bool) bool {
	obstacles := 0

	for i, n := range dir {
		rock := curr
		rock[i] = curr[i] + n
		_, isRock := boulder[rock]

		if isRock {
			obstacles++
		}
	}

	return obstacles != 2
}

func main() {
	boulder := parse()
	dirs := [6]Point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}

	curr := Point{0, 0, 0}
	queue := Queue[Point]{}
	queue.Push(curr)
	seen := make(map[Point]bool)
	for {
		curr = queue.Pop()
		foundEdge := false
		seen[curr] = true
		for _, dir := range dirs[:3] {
			next := Point{curr[0] + dir[0], curr[1] + dir[1], curr[2] + dir[2]}
			if _, ok := (*boulder)[next]; ok {
				foundEdge = true
			} else {
				queue.Push(next)
			}
		}
		if foundEdge {
			break
		}
	}

	dirs2 := []Point{
		{1, 1, 0},
		{1, -1, 0},
		{-1, 1, 0},
		{-1, -1, 0},
		{0, 1, 1},
		{0, 1, -1},
		{0, -1, 1},
		{0, -1, -1},
		{1, 0, 1},
		{-1, 0, 1},
		{1, 0, -1},
		{-1, 0, -1},
	}

	// dirs3 := []Point{
	// 	{1, 1, 1},
	// 	{1, 1, -1},
	// 	{1, -1, 1},
	// 	{1, -1, -1},
	// 	{-1, 1, 1},
	// 	{-1, 1, -1},
	// 	{-1, -1, 1},
	// 	{-1, -1, -1},
	// }

	count := 0

	queue = Queue[Point]{}
	queue.Push(curr)
	seen = make(map[Point]bool)
	edges := make(map[[6]int]bool)
	totalSides := 0
	for len(queue) > 0 {
		count++
		curr = queue.Pop()
		seen[curr] = true

		hasAnyEdge := false
		for _, dir := range dirs {
			next := Point{curr[0] + dir[0], curr[1] + dir[1], curr[2] + dir[2]}
			_, hasEdge := (*boulder)[next]
			_, edgeSeen := edges[[6]int{curr[0], curr[1], curr[2], next[0], next[1], next[2]}]
			if hasEdge && !edgeSeen {
				edges[[6]int{curr[0], curr[1], curr[2], next[0], next[1], next[2]}] = true
				totalSides++
				hasAnyEdge = true
			}
		}
		if !hasAnyEdge {
			continue
		}
		for _, dir := range dirs {
			next := Point{curr[0] + dir[0], curr[1] + dir[1], curr[2] + dir[2]}
			_, isRock := (*boulder)[next]
			if _, ok := seen[next]; !ok && !isRock {
				queue.Push(next)
			}
		}
		for _, dir := range dirs2 {
			next := Point{curr[0] + dir[0], curr[1] + dir[1], curr[2] + dir[2]}
			_, edgeSeen := seen[next]
			_, isRock := (*boulder)[next]

			if !edgeSeen && canGoToNextEdge(curr, dir, *boulder) && !isRock {
				queue.Push(next)
			}
		}
	}

	fmt.Println("totalSides", totalSides)
}
