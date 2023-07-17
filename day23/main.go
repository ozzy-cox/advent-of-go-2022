package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Elf struct {
	id  int
	pos [2]int
}

var dirVecs = map[string][2]int{
	"N":  {-1, 0},
	"S":  {1, 0},
	"W":  {0, -1},
	"E":  {0, 1},
	"NW": {-1, -1},
	"NE": {-1, 1},
	"SW": {1, -1},
	"SE": {1, 1},
}

var dir2CheckMap = map[string][3]string{
	"N": {"N", "NE", "NW"},
	"S": {"S", "SE", "SW"},
	"W": {"W", "SW", "NW"},
	"E": {"E", "SE", "NE"},
}

func canMove(grid [][]byte, elf Elf, dir string) bool {
	hasAnyNeighbors := false
	for _, v := range dirVecs {
		if grid[v[0]+elf.pos[0]][v[1]+elf.pos[1]] != '.' {
			hasAnyNeighbors = true
		}
	}
	if !hasAnyNeighbors {
		return false
	}

	for _, dir2Check := range dir2CheckMap[dir] {
		dirVec := dirVecs[dir2Check]
		if grid[dirVec[0]+elf.pos[0]][dirVec[1]+elf.pos[1]] != '.' {
			return false
		}
	}
	return true
}

func print(grid [][]byte) {
	for i := 10; i < len(grid); i++ {
		for j := 10; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args

	file, _ := os.ReadFile(args[1])

	lines := strings.Split(string(file), "\n")

	baseLen := len(lines[0])
	grid := make([][]byte, 0)
	for i := 0; i < 20; i++ {
		grid = append(grid, make([]byte, baseLen+40))
	}
	for _, line := range lines {
		grid = append(grid, make([]byte, baseLen+40))
		for i, char := range line {
			grid[len(grid)-1][20+i] = byte(char)
		}
	}
	for i := 0; i < 20; i++ {
		grid = append(grid, make([]byte, baseLen+40))
	}

	elves := make([]Elf, 0)

	for i, line := range grid {
		for j, char := range line {
			if char == '#' {
				elves = append(elves, Elf{
					id:  len(elves),
					pos: [2]int{i, j},
				})
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = '.'
		}
	}
	for _, elf := range elves {
		grid[elf.pos[0]][elf.pos[1]] = '#'
	}
	print(grid)

	rounds := 10

	proposals := make(map[Elf][2]int)
	positionRequests := make(map[[2]int]int)

	dirIdx := 0
	dirs := []string{"N", "S", "W", "E"}
	for i := 0; i < rounds; i++ {
		// Find proposed moves
		for _, elf := range elves {
			for j := 0; j < len(dirs); j++ {
				dir := dirs[(dirIdx+j)%len(dirs)]
				if canMove(grid, elf, dir) {
					dirVec := dirVecs[dir]
					nextPos := [2]int{elf.pos[0] + dirVec[0], elf.pos[1] + dirVec[1]}
					proposals[elf] = nextPos
					_, ok := positionRequests[nextPos]
					if !ok {
						positionRequests[nextPos] = 1
					} else {
						positionRequests[nextPos]++
					}
					break
				}
			}
		}
		dirIdx = (dirIdx + 1) % 4

		// Update the grid
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j] = '.'
			}
		}

		// Make proposed moves
		for elf, nextPos := range proposals {
			if val, ok := positionRequests[nextPos]; ok && val == 1 {
				elves[elf.id].pos = nextPos
			}
		}

		for _, elf := range elves {
			grid[elf.pos[0]][elf.pos[1]] = '#'
		}

		proposals = make(map[Elf][2]int)
		positionRequests = make(map[[2]int]int)
		print(grid)
	}

	minX := math.MaxInt
	maxX := 0
	minY := math.MaxInt
	maxY := 0
	for _, elf := range elves {
		if elf.pos[0] < int(minX) {
			minX = elf.pos[0]
		}
		if elf.pos[0] > maxX {
			maxX = elf.pos[0]
		}
		if elf.pos[1] < int(minY) {
			minY = elf.pos[1]
		}
		if elf.pos[1] > int(maxY) {
			maxY = elf.pos[1]
		}
	}

	total := 0
	for i := minX; i < maxX+1; i++ {
		for j := minY; j < maxY+1; j++ {
			if grid[i][j] != '#' {
				total++
			}
		}
	}
	fmt.Println(total)
}
