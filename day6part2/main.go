package main

import (
	"fmt"
	"os"

	"github.com/willf/bitset"
)

type queue []byte

func (q *queue) Push(b byte) {
	*q = append(*q, b)
}

func (q *queue) Pop() byte {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *queue) Peek() byte {
	return (*q)[0]
}

func main() {
	line, _ := os.ReadFile("input.txt")

	queue := make(queue, 0)
	b := bitset.New(256)

	for i := 0; i < 14; i++ {
		queue.Push(line[i])
		b.Set(uint(line[i]))
	}

	for i, c := range line[14:] {
		if b.Count() == 14 {
			fmt.Println(i + 14)
			break
		}
		b.ClearAll()
		queue.Pop()
		queue.Push(c)
		for j := 0; j < 14; j++ {
			b.Set(uint(queue[j]))
		}
	}
}
