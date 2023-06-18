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
}

func (p1 Point) isAdjacent(p2 Point) bool {
	if p1.x > p2.x {
		if dist := p1.x - p2.x; dist > 1 {
			return false
		}
	}

	if p1.x < p2.x {
		if dist := p2.x - p1.x; dist > 1 {
			return false
		}
	}

	if p1.y > p2.y {
		if dist := p1.y - p2.y; dist > 1 {
			return false
		}
	}

	if p1.y < p2.y {
		if dist := p2.y - p1.y; dist > 1 {
			return false
		}
	}

	return true
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	tailHistory := make(map[string]bool)
	tailHistory["0-0"] = true

	tail := Point{0, 0}
	head := Point{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		dir := lineParts[0]
		steps, err := strconv.Atoi(lineParts[1])
		if err != nil {
			panic("Cant convert to int")
		}

		for i := 0; i < steps; i++ {
			headPrev := head

			switch dir {
			case "U":
				head.y += 1
			case "R":
				head.x += 1
			case "D":
				head.y += -1
			case "L":
				head.x += -1
			}

			if !tail.isAdjacent(head) {
				tail = headPrev
				tailHistory[fmt.Sprintf("%d-%d", tail.x, tail.y)] = true
			}
		}
	}
	count := 0

	for range tailHistory {
		count++
	}
	println(count)
}
