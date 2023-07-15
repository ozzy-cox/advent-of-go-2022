package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	add Operation = iota
	subtract
	multiply
	divide
)

type Monke struct {
	id            string
	hasOperation  bool
	operationType Operation
	value         int
	children      [2]string
}

var opMap = map[string]Operation{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func dfs(monkes map[string]Monke, curr string) int {
	currMonke := monkes[curr]
	if !currMonke.hasOperation {
		return currMonke.value
	}
	value := 0
	subValue1 := dfs(monkes, currMonke.children[0])
	subValue2 := dfs(monkes, currMonke.children[1])

	switch currMonke.operationType {
	case add:
		value = subValue1 + subValue2
	case subtract:
		value = subValue1 - subValue2
	case multiply:
		value = subValue1 * subValue2
	case divide:
		value = subValue1 / subValue2
	}

	return value
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	monkes := make(map[string]Monke)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ": ")
		var monke Monke
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			nextTokens := strings.Split(tokens[1], " ")
			children := [2]string{nextTokens[0], nextTokens[2]}
			monke = Monke{
				id:            tokens[0],
				hasOperation:  true,
				operationType: opMap[nextTokens[1]],
				children:      children,
			}
		} else {
			monke = Monke{
				id:           tokens[0],
				hasOperation: false,
				value:        value,
			}
		}

		monkes[tokens[0]] = monke
	}

	// TODO result cache may be useful

	rootValue := dfs(monkes, "root")
	fmt.Println(rootValue)
}
