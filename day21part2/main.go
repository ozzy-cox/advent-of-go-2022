package main

import (
	"bufio"
	"errors"
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

func dfs(monkes map[string]Monke, curr string) (int, error) {
	currMonke := monkes[curr]
	if currMonke.id == "humn" {
		return 0, errors.New("Error")
	}
	if !currMonke.hasOperation {
		return currMonke.value, nil
	}
	value := 0
	subValue1, err1 := dfs(monkes, currMonke.children[0])
	subValue2, err2 := dfs(monkes, currMonke.children[1])

	if err1 != nil || err2 != nil {
		return 0, errors.New("Error")
	}

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

	currMonke.value = value
	currMonke.hasOperation = false

	monkes[currMonke.id] = currMonke
	return value, nil
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

	rootValue, _ := dfs(monkes, "root")
	fmt.Println(rootValue)
	rootEq := 0
	isLeftMonke := false

	curr := "root"
	rootChildren := monkes["root"].children
	for _, child := range rootChildren {
		if val, ok := monkes[child]; ok {
			if !val.hasOperation {
				rootEq = val.value
			} else {
				curr = child
			}
		}
	}
	fmt.Println(rootEq)

	acc := rootEq
	// Traverse to "humn" node while accumulating the value in reverse operations

	for curr != "humn" {
		currMonke := monkes[curr]
		var monkeWithoutValue Monke
		var monkeWithValue Monke
		for i, monkeId := range currMonke.children {
			if val, ok := monkes[monkeId]; ok {
				if val.hasOperation || monkeId == "humn" {
					monkeWithoutValue = val
					isLeftMonke = i == 0
				} else {
					monkeWithValue = val
				}
			}
		}
		value := monkeWithValue.value
		if curr != "root" {
			if !isLeftMonke {
				switch currMonke.operationType {
				case add:
					acc = acc - value
				case subtract:
					acc = value - acc
				case multiply:
					acc = acc / value
				case divide:
					acc = value / acc
				}
			} else {
				switch currMonke.operationType {
				case add:
					acc = acc - value
				case subtract:
					acc = acc + value
				case multiply:
					acc = acc / value
				case divide:
					acc = value * acc
				}
			}
		}

		curr = monkeWithoutValue.id
	}

	fmt.Println(acc)
}
