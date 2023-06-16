package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() byte {
	return s[len(s)-1]
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	tableLines := []string{}

	for scanner.Scan() {
		// Read until the empty line
		line := scanner.Text()
		if line == "" {
			break
		}
		tableLines = append(tableLines, line)
	}

	// Create stacks

	stackCount := (len(tableLines[len(tableLines)-1]) + 1) / 4

	stacks := make([]stack, stackCount)

	tableValues := tableLines[:len(tableLines)-1]

	for i := len(tableValues) - 1; i >= 0; i-- {
		for j := 1; j < len(tableValues[i]); j += 4 {
			if tableValues[i][j-1] == '[' && tableValues[i][j+1] == ']' {
				// fmt.Printf("%c %d\n", tableValues[i][j], tableValues[i][j])
				stacks[(j-1)/4] = stacks[(j-1)/4].Push(tableValues[i][j])
			}
		}
	}

	for _, stack := range stacks {
		fmt.Println(stack)
	}

	// Parse operation information
	for scanner.Scan() {
		line := scanner.Text()
		// seperate the move part

		operationSplit := strings.Split(line, " ")

		amount, _ := strconv.Atoi(operationSplit[1])
		from, _ := strconv.Atoi(operationSplit[3])
		to, _ := strconv.Atoi(operationSplit[5])

		for i := 0; i < amount; i++ {
			fmt.Printf("%d %d %d\n", amount, from, to)
			stak, temp := stacks[from-1].Pop()
			stacks[from-1] = stak

			stacks[to-1] = stacks[to-1].Push(temp)
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Peek())
	}
}
