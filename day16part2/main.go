package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func getKeys[T any](_map map[string]T) []string {
	keys := make([]string, len(_map))

	i := 0
	for k := range _map {
		keys[i] = k
		i++
	}
	return keys
}

func solve(time int, curr Node, sum int, openedValves map[string]bool, visited map[string]bool, dists map[string]map[string]int, nodes map[string]Node, results map[string]int, path []string) (int, []string) {
	maxVal := 0
	maxOfPaths := make([]int, 0)
	maxPath := make([]string, 0)
	paths := make([][]string, 0)
	visited[curr.id] = true
	if curr.id != "AA" {
		path = append(path, curr.id)
	}
	if time > 0 {
		for k := range dists[curr.id] {
			if _, ok := visited[k]; !ok {
				//  Do not open the valve
				resOpen := 0
				pathOpen := make([]string, 0)
				remTime := time - dists[curr.id][k]
				resClose, pathClose := solve(remTime, nodes[k], sum, openedValves, visited, dists, nodes, results, path)

				// Open the valve
				if _, ok := openedValves[k]; !ok && remTime > 0 {
					openedValves[k] = true
					resOpen, pathOpen = solve(remTime-1, nodes[k], sum+((remTime-1)*nodes[k].rate), openedValves, visited, dists, nodes, results, path)
					resOpen = resOpen + ((remTime - 1) * nodes[k].rate)
					pathOpen = append(pathOpen, nodes[k].id)
					delete(openedValves, k)
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

	results[strings.Join(path, ",")] = sum

	delete(visited, curr.id)
	return maxVal, maxPath
}

func isDisjoint(set1 []string, set2 []string) bool {
	for _, i := range set1 {
		for _, j := range set2 {
			if i == j {
				return false
			}
		}
	}
	return true
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
		fmt.Println(dists)
	}

	openedValves := make(map[string]bool)
	visited := make(map[string]bool)

	results := make(map[string]int)
	start := time.Now()

	path := make([]string, 0)
	fmt.Println(solve(26, nodes["AA"], 0, visited, openedValves, dists, nodes, results, path))
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	fmt.Println(path)
	keys := getKeys(results)
	fmt.Println(len(results))

	for k, v := range results {
		if v == 0 {
			delete(results, k)
		}
	}
	// sort.SliceStable(keys, func(i int, j int) bool {
	// 	return results[keys[i]] < results[keys[j]]
	// })

	// Find disjoint sets

	fmt.Println(len(keys))
	max := 0
	var seta []string
	var setb []string
	for i, key := range keys {
		set1 := strings.Split(key, ",")
		fmt.Println(i)
		for _, key2 := range keys {
			set2 := strings.Split(key2, ",")
			tempMax := results[key] + results[key2]
			if isDisjoint(set1, set2) && max < tempMax {
				max = tempMax
				seta = set1
				setb = set2

			}
		}
	}

	fmt.Println(max)
	fmt.Println(seta)
	fmt.Println(setb)
}
