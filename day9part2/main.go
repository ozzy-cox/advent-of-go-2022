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

func draw(head Point, snake []Point) {
	for y := 10; y > -10; y-- {
		for x := -10; x < 10; x++ {
			p := true
			if x == head.x && y == head.y {
				fmt.Print("H")
				p = false
			}
			for i, v := range snake {
				if v.x == x && v.y == y {
					fmt.Print(i + 1)
					p = false
					break
				}
			}
			if p {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
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
	tailHistory["0x0"] = true

	head := Point{0, 0}

	snake := make([]Point, 9)
	for i := range snake {
		snake[i] = Point{0, 0}
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		dir := lineParts[0]
		steps, err := strconv.Atoi(lineParts[1])
		if err != nil {
			panic("Cant convert to int")
		}

		for i := 0; i < steps; i++ {

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
			next := head

			// TODO maybe dont check the last item
			for i := 0; i < len(snake); i++ {
				if !snake[i].isAdjacent(next) {

					if next.x != snake[i].x && next.y != snake[i].y {
						// Go diagonally
						if next.x > snake[i].x {
							snake[i].x += 1
						} else {
							snake[i].x -= 1
						}

						if next.y > snake[i].y {
							snake[i].y += 1
						} else {
							snake[i].y -= 1
						}
					} else {
						if next.x > snake[i].x {
							snake[i].x += 1
						} else if next.x < snake[i].x {
							snake[i].x -= 1
						}

						if next.y > snake[i].y {
							snake[i].y += 1
						} else if next.y < snake[i].y {
							snake[i].y -= 1
						}
					}
					next = snake[i]

					if i == len(snake)-1 {
						tailHistory[fmt.Sprintf("%dx%d", snake[len(snake)-1].x, snake[len(snake)-1].y)] = true
					}
				} else {
					break
				}
			}
		}
		draw(head, snake)
	}
	count := 0

	for range tailHistory {
		count++
	}
	println(count)
}
