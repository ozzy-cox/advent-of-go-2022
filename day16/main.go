package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id         string
	gid        uint
	adjs       []string
	rate       int
	isReleased bool
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

func solve(time int, curr Node, openedValves map[string]bool, dists map[string]map[string]int, nodes map[string]Node) int {
	maxVal := 0
	fmt.Println(time)
	for k := range dists[curr.id] {

		// fmt.Println(k, v, maxVal)
		//  Do not open the valve
		resOpen := 0
		remTime := time - dists[curr.id][k]
		resClose := solve(remTime, nodes[k], openedValves, dists, nodes)

		// Open the valve
		if _, ok := openedValves[k]; !ok {
			fmt.Println("enterereered")
			// fmt.Println(k, remTime, nodes[k].rate)
			openedValves[k] = true
			maxVal += (remTime - 1) * nodes[k].rate
			resOpen = solve(remTime-1, nodes[k], openedValves, dists, nodes)
		}
		// fmt.Println(resClose, resOpen)
		maxVal += int(math.Max(float64(resClose), float64(resOpen)))
	}
	return maxVal
}

func main() {
	file, _ := os.Open("smallinput.txt")

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
	fmt.Println(nodes)

	dists := make(map[string]map[string]int)

	for _, node := range nodes {
		if node.id != "AA" && node.rate != 0 {
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
						fmt.Println(curr.node.id, adjNode)
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

	openedValves := make(map[string]bool)
	fmt.Println(solve(30, nodes["AA"], openedValves, dists, nodes))
}
