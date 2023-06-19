package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maja42/goval"
)

type Monke struct {
	items           *Queue
	operation       string // TODO Equation, how to encode in code ???
	test            int    // Divisible by test
	throwToIfTrue   int
	throwToIfFalse  int
	inspectionCount *int
}

type Queue []int

func (q *Queue) Push(b int) {
	*q = append(*q, b)
}

func (q *Queue) Pop() int {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue) Peek() int {
	return (*q)[0]
}

func Map[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	eval := goval.NewEvaluator()
	fileContents, err := os.ReadFile("input.txt")
	panicOnError(err)

	monkeInputs := strings.Split(string(fileContents), "\n\n")
	monkes := make([]Monke, 0)

	for _, monke := range monkeInputs {
		lines := strings.Split(monke, "\n")
		startingItems := Queue(Map(Map(strings.Split(strings.Split(lines[1], ":")[1], ","), func(s string) string {
			return strings.TrimSpace(s)
		}), func(s string) int {
			val, err := strconv.Atoi(s)
			panicOnError(err)
			return val
		}))

		operation := strings.TrimSpace(strings.Split(lines[2], "=")[1])

		test := strings.TrimSpace(strings.Split(lines[3], "by")[1])
		testValue, err := strconv.Atoi(test)
		panicOnError(err)

		throwToIfTrue := strings.TrimSpace(strings.Split(lines[4], "monkey")[1])
		throwToIfTrueValue, err := strconv.Atoi(throwToIfTrue)
		panicOnError(err)

		throwToIfFalse := strings.TrimSpace(strings.Split(lines[5], "monkey")[1])
		throwToIfFalseValue, err := strconv.Atoi(throwToIfFalse)
		panicOnError(err)

		inspectionCount := 0

		monkes = append(monkes, Monke{
			items:           &startingItems,
			operation:       operation,
			test:            testValue,
			throwToIfTrue:   throwToIfTrueValue,
			throwToIfFalse:  throwToIfFalseValue,
			inspectionCount: &inspectionCount,
		})
	}

	commonModulo := 1

	for _, monke := range monkes {
		commonModulo *= monke.test
	}

	for i := 0; i < 10000; i++ {
		for _, monke := range monkes {
			for len(*monke.items) > 0 {
				currentItem := monke.items.Pop()
				itemAfterCalculation, err := eval.Evaluate(monke.operation, map[string]interface{}{
					"old": currentItem,
				}, nil)
				panicOnError(err)
				intItemAfterCalculation := itemAfterCalculation.(int)
				*monke.inspectionCount += 1

				if intItemAfterCalculation%monke.test == 0 {
					monkes[monke.throwToIfTrue].items.Push(intItemAfterCalculation % commonModulo)
				} else {
					monkes[monke.throwToIfFalse].items.Push(intItemAfterCalculation % commonModulo)
				}
			}
		}

		if i%1000 == 999 {
			for _, monke := range monkes {
				fmt.Println(*monke.inspectionCount)
			}
			fmt.Println()
		}

	}

	maxOfMonkes := [...]int{0, 0}
	maxMonkeIndices := [...]int{-1, -1}

	for i, monke := range monkes {
		max := *monke.inspectionCount
		if max > maxOfMonkes[0] {
			temp := maxOfMonkes[0]
			tempIdx := maxMonkeIndices[0]
			maxOfMonkes[0] = max
			maxMonkeIndices[0] = i
			maxOfMonkes[1] = temp
			maxMonkeIndices[1] = tempIdx
		} else if max > maxOfMonkes[1] {
			maxOfMonkes[1] = max
			maxMonkeIndices[1] = i
		}
	}

	fmt.Println(maxOfMonkes[0] * maxOfMonkes[1])
}
