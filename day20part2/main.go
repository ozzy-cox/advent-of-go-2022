package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int64
}

func findIndex(arr []*Node, needle *Node) int {
	for i, item := range arr {
		if item == needle {
			return i
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	nodeOrder := make([]*Node, 0)

	items := make([]*Node, 0)

	decryptionKey := int64(811589153)

	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.ParseInt(line, 10, 64)
		node := Node{
			value: value * decryptionKey,
		}
		items = append(items, &node)
		nodeOrder = append(nodeOrder, &node)
	}

	for i := 0; i < 10; i++ {
		for _, node := range nodeOrder {
			idx := findIndex(items, node)

			move := node.value % int64(len(nodeOrder)-1)

			endPos := idx + int(move)
			// qu := (idx + node.value) / len(nodeOrder)
			if endPos >= len(nodeOrder) {
				endPos = endPos - len(nodeOrder) + 1
			} else if endPos <= 0 {
				endPos = len(nodeOrder) + (endPos - 1)
			}

			items = append(items[:idx], items[idx+1:]...)
			items = append(items[:endPos+1], items[endPos:]...)
			items[endPos] = node
		}
	}

	sum := int64(0)
	zeroNodeIdx := 0
	for i, node := range items {
		if node.value == 0 {
			zeroNodeIdx = i
		}
	}
	for i := 1000; i <= 3000; i += 1000 {
		fmt.Println("value", items[(zeroNodeIdx+i)%len(items)].value)
		sum += items[(zeroNodeIdx+i)%len(items)].value
	}

	fmt.Println(sum)
}
