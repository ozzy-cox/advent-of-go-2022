package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bits-and-blooms/bitset"
)

type Node struct {
	id   string
	gid  uint
	adjs []string
	rate int
	// isReleased bool
}
type Queue[T any] []T

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

type QItem struct {
	dist int
	node Node
}
type Result struct {
	bits  bitset.BitSet
	value int
}

func max(arr []int) int {
	maxIdx := 0
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			maxIdx = i
			max = arr[i]
		}
	}
	return maxIdx
}

func solve(time int, curr Node, sum int, visited *bitset.BitSet, openedValves *bitset.BitSet, dists map[string]map[string]int, nodes map[string]Node, results map[string]Result, path []string) (int, []string) {
	maxVal := 0
	maxOfPaths := make([]int, 0)
	var maxPath []string
	paths := make([][]string, 0)
	visited.Set(curr.gid)
	if curr.id != "AA" {
		path = append(path, curr.id)
	}
	if time > 0 {
		for k, dist := range dists[curr.id] {
			node := nodes[k]
			remTime := time - dist
			if !visited.Test(node.gid) && remTime > 0 {
				//  Do not open the valve
				resOpen := 0
				pathOpen := make([]string, 0)
				resClose, pathClose := solve(remTime, node, sum, visited, openedValves, dists, nodes, results, path)

				// Open the valve
				if !openedValves.Test(node.gid) {
					openedValves.Set(node.gid)
					resOpen, pathOpen = solve(remTime-1, node, sum+((remTime-1)*node.rate), visited, openedValves, dists, nodes, results, path)
					resOpen = resOpen + ((remTime - 1) * node.rate)
					pathOpen = append(pathOpen, node.id)
					openedValves.Clear(node.gid)
				}
				if resClose > resOpen {
					paths = append(paths, pathClose)
					maxOfPaths = append(maxOfPaths, resClose)
				} else {
					paths = append(paths, pathOpen)
					maxOfPaths = append(maxOfPaths, resOpen)
				}
			}
		}
	}

	if len(maxOfPaths) > 0 {
		maxIdx := max(maxOfPaths)
		maxVal = maxOfPaths[maxIdx]
		maxPath = paths[maxIdx]
	}

	resultKey := openedValves.String()
	if result, ok := results[resultKey]; ok {
		if result.value < sum {
			results[resultKey] = Result{
				value: sum,
				bits:  *openedValves.Clone(),
			}
		}
	} else {
		results[resultKey] = Result{
			value: sum,
			bits:  *openedValves.Clone(),
		}
	}

	visited.Clear(curr.gid)
	return maxVal, maxPath
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	nodes := make(map[string]Node)

	gid := uint(0)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		scTokens := strings.Split(line, ";")
		firstHalfTokens := strings.Split(scTokens[0], " ")
		secondHalfTokens := strings.Split(scTokens[1], " ")
		id := firstHalfTokens[1]
		rate, _ := strconv.Atoi(strings.Split(firstHalfTokens[4], "=")[1])
		adjs := strings.Split(strings.Join(secondHalfTokens[5:], " "), ", ")

		node := Node{
			id:   id,
			gid:  gid,
			adjs: adjs,
			rate: rate,
		}
		nodes[id] = node
		gid++
	}

	dists := make(map[string]map[string]int)

	for _, node := range nodes {
		if node.rate == 0 && node.id != "AA" {
			continue
		}
		visited := make(map[string]bool)
		visited[node.id] = true

		dists[node.id] = make(map[string]int)
		dists[node.id]["AA"] = 0
		dists[node.id][node.id] = 0

		q := make(Queue[QItem], 0)
		q.Push(QItem{
			dist: 0,
			node: node,
		})
		for len(q) > 0 {
			curr := q.Pop()
			dist := curr.dist

			for _, adjNode := range curr.node.adjs {
				if _, ok := visited[adjNode]; !ok {
					visited[adjNode] = true
					if nodes[adjNode].rate != 0 {
						dists[node.id][adjNode] = dist + 1
					}
					q.Push(QItem{
						dist: dist + 1,
						node: nodes[adjNode],
					})
				}
			}
		}
		delete(dists[node.id], node.id)
		if node.id != "AA" {
			delete(dists[node.id], "AA")
		}
	}

	openedValves := bitset.New(uint(len(dists)))
	visited := bitset.New(uint(len(dists)))

	results := make(map[string]Result)

	path := make([]string, 0)
	start := time.Now()

	fmt.Println(solve(26, nodes["AA"], 0, visited, openedValves, dists, nodes, results, path))
	elapsed := time.Since(start)
	fmt.Printf("Elapsed %s\n", elapsed)
	// Find disjoint sets
	max := 0
	for _, value1 := range results {
		for _, value2 := range results {
			tempMax := value2.value + value1.value
			if !value2.bits.Intersection(&value1.bits).Any() && max < tempMax {
				max = tempMax
			}
		}
	}
	fmt.Println(max)
}
