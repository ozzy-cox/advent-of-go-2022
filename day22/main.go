package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func print(dirCharMap map[int]byte, grid [][]byte, pos Point, dirType int) {
	fmt.Println()
	for i := 0; i < len(grid); i++ {
		for j, char := range grid[i] {
			if pos[0] == i && pos[1] == j {
				switch dirType {
				case 0:
					fmt.Printf("%c", '>')
				case 1:
					fmt.Printf("%c", 'v')
				case 2:
					fmt.Printf("%c", '<')
				case 3:
					fmt.Printf("%c", '^')
				}
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Point [2]int

func getNextPosition(grid [][]byte, curr Point, dir Point) Point {
	to := curr
	for i := 0; i < len(curr); i++ {
		to[i] += dir[i]
	}

	isOutOfBounds := false
	if to[0] == len(grid) || to[0] < 0 || to[1] == len(grid[0]) || to[1] < 0 {
		isOutOfBounds = true
	}

	currValue := byte(0)
	if isOutOfBounds || grid[to[0]][to[1]] == ' ' {
		// Adjust
		for currValue == 0 || currValue == ' ' {
			to[0] = to[0] % len(grid)
			if to[0] < 0 {
				to[0] = len(grid) + to[0]
			}
			to[1] = to[1] % len(grid[0])
			if to[1] < 0 {
				to[1] = len(grid[0]) + to[1]
			}
			currValue = grid[to[0]][to[1]]
			if currValue != ' ' {
				break
			}
			to[0] += dir[0]
			to[1] += dir[1]
		}
	}

	return to
}

func canMove(grid [][]byte, pos Point) bool {
	return grid[pos[0]][pos[1]] != '#'
}

func main() {
	file, _ := os.ReadFile(os.Args[1])

	tokens := strings.Split(string(file), "\n")

	instructions := tokens[len(tokens)-1]
	var grid [][]byte

	longestLine := 0
	for _, st := range tokens[:len(tokens)-2] {
		if len(st) > longestLine {
			longestLine = len(st)
		}
	}

	for _, st := range tokens[:len(tokens)-2] {
		line := make([]byte, 0)
		for i := 0; i < longestLine; i++ {
			if i < len(st) {
				line = append(line, byte(st[i]))
			} else {
				line = append(line, byte(' '))
			}
		}
		grid = append(grid, line)
	}

	buff := make([]byte, 0)

	dirs := []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	dirCharMap := map[int]byte{
		0: '>',
		1: 'v',
		2: '<',
		3: '^',
	}
	var pos Point
	// Find starting position

	for i, char := range grid[0] {
		if char == '.' {
			pos = Point{0, i}
			break
		}
	}

	dirIdx := 0

	idx := 0
	for {
		if idx >= len(instructions) {
			break
		}
		for idx < len(instructions) && unicode.IsDigit(rune(instructions[idx])) {
			buff = append(buff, []byte(instructions)[idx])
			idx++
		}
		if idx >= len(instructions) {
			idx--

			// buff = append(buff, []byte(instructions)[idx])
			moveAmt, _ := strconv.Atoi(string(buff[:]))
			for i := 0; i < moveAmt; i++ {
				nextPos := getNextPosition(grid, pos, dirs[dirIdx])

				if canMove(grid, nextPos) {
					grid[pos[0]][pos[1]] = dirCharMap[dirIdx]
					pos = nextPos
					// Update position
				} else {
					break
				}
			}

			// execute last instr here
			break
		}
		buff = append(buff, []byte(instructions)[idx])
		idx++

		moveAmt, _ := strconv.Atoi(string(buff[:len(buff)-1]))
		dirChange := buff[len(buff)-1]

		for i := 0; i < moveAmt; i++ {
			nextPos := getNextPosition(grid, pos, dirs[dirIdx])

			if canMove(grid, nextPos) {
				grid[pos[0]][pos[1]] = dirCharMap[dirIdx]
				pos = nextPos
				// Update position
			} else {
				break
			}
		}

		nextDirIdx := dirIdx
		if dirChange == 'R' {
			nextDirIdx = (dirIdx + 1) % 4
		} else {
			nextDirIdx--
		}

		if nextDirIdx < 0 {
			nextDirIdx = 4 + nextDirIdx
		}
		dirIdx = nextDirIdx
		buff = make([]byte, 0)
	}

	fmt.Println(((pos[0] + 1) * 1000) + ((pos[1] + 1) * 4) + dirIdx)
}
