package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type (
	Point [2]int
	Jet   struct {
		pos Point
		dir Point
	}
	Grid         [][]byte
	Queue[T any] []T
	State        struct {
		pos      Point
		round    int
		index    int
		priority int
	}
)

type PriorityQueue []*State

func calcManDist(x Point, y Point) int {
	return (y[0] - x[0]) + (y[1] - x[1])
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *State, pos Point, priority int) {
	item.pos = pos
	item.priority = priority
	heap.Fix(pq, item.index)
}

var dirMap = map[byte]Point{
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
	'^': {-1, 0},
}

func calculateJets(jets []Jet, dimensions [2]int) []Jet {
	newJets := make([]Jet, 0)
	for _, jet := range jets {
		newPos := Point{(jet.pos[0] + jet.dir[0] + dimensions[0]) % dimensions[0], (jet.pos[1] + jet.dir[1] + dimensions[1]) % dimensions[1]}
		newJets = append(newJets, Jet{
			pos: newPos,
			dir: jet.dir,
		})
	}
	return newJets
}

func createGridFromJets(jets []Jet, dimensions Point) Grid {
	grid := make(Grid, dimensions[0])
	for i := 0; i < dimensions[0]; i++ {
		grid[i] = make([]byte, dimensions[1])
	}

	for _, jet := range jets {
		grid[jet.pos[0]][jet.pos[1]] = '#'
	}
	return grid
}

func isPointInside(point Point, dimensions Point) bool {
	return point[0] >= 0 && point[1] >= 0 && point[0] < dimensions[0] && point[1] < dimensions[1]
}

func main() {
	file, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(file)

	grid := make([][]byte, 0)

	start := Point{0, 0}
	var end Point
	jets := make([]Jet, 0)

	var line string
	i := 0
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, []byte(line[1:len(line)-1]))
		i++
	}

	grid = grid[1 : len(grid)-1]
	for i, line := range grid {
		for j, char := range line {
			if char != '.' && char != '#' {
				jets = append(jets, Jet{
					Point{i, j},
					dirMap[byte(char)],
				})
			}
		}
	}

	for i, char := range line {
		if char == '.' {
			end = Point{len(grid) - 1, i - 1}
		}
	}
	min := math.MaxInt

	dimensions := Point{len(grid), len(grid[0])}
	jetsCache := make([][]Jet, 0)
	gridCache := make([]Grid, 0)

	jetsCache = append(jetsCache, calculateJets(jets, dimensions))
	gridCache = append(gridCache, createGridFromJets(jets, dimensions))
	seen := make(map[State]bool)

	q := make(PriorityQueue, 0)
	state := &State{
		pos:      start,
		priority: calcManDist(start, end),
		round:    1,
	}
	heap.Init(&q)
	heap.Push(&q, state)
	q.update(state, state.pos, calcManDist(start, end))

	for q.Len() > 0 {
		state := heap.Pop(&q).(*State)
		round := state.round
		curr := state.pos
		if _, ok := seen[*state]; ok {
			continue
		}

		seen[*state] = true
		if round > min {
			continue
		}

		if curr == end {
			if round < min {
				min = round
				fmt.Println("found", min)
			}
			continue
		}

		if len(jetsCache) <= round {
			jetsCache = append(jetsCache, calculateJets(jetsCache[round-1], dimensions))
			gridCache = append(gridCache, createGridFromJets(jetsCache[round-1], dimensions))
		}

		nextGrid := gridCache[round]
		for _, dir := range dirMap {
			nextPos := Point{curr[0] + dir[0], curr[1] + dir[1]}
			if isPointInside(nextPos, dimensions) {
				if nextGrid[nextPos[0]][nextPos[1]] != '#' {
					state := &State{
						pos:      nextPos,
						round:    round + 1,
						priority: calcManDist(nextPos, end),
					}
					heap.Push(&q, state)
					// q.update(state, state.pos, calcManDist(start, end))
				}
			}
		}
		if nextGrid[curr[0]][curr[1]] != '#' || round == 1 {
			state := &State{
				pos:      curr,
				round:    round + 1,
				priority: calcManDist(curr, end),
			}
			heap.Push(&q, state)
		}
	}
	// TODO Dont forget +1
	fmt.Println(min)
}
