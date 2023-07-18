package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var charIntMap = map[byte]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(file), "\n")

	longestLine := 0
	for _, line := range lines {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}
	numbers := make([][]int, 0)

	for _, line := range lines {
		numbers = append(numbers, make([]int, longestLine))
		for i, char := range line {
			number := charIntMap[byte(char)]
			numbers[len(numbers)-1][i+(longestLine-len(line))] = number
		}
	}

	numbersSum := make([]int, longestLine)

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < longestLine; i++ {
			number := numbers[j]
			numbersSum[i] += number[i]
		}
	}

	decimal := 0
	for i := range numbersSum {
		val := numbersSum[len(numbersSum)-1-i]
		decimal += val * int(math.Pow(float64(5), float64(i)))
	}

	snafuSum := make([]int, 0)

	take := 0
	for i := len(numbersSum) - 1; i >= 0; i-- {
		dig := numbersSum[i]
		tot := dig + take
		quot := tot / 5
		rem := tot % 5
		if rem == 4 {
			quot++
			rem = -1
		} else if rem == 3 {
			quot++
			rem = -2
		} else if rem == -3 {
			quot--
			rem = 2
		} else if rem == -4 {
			quot--
			rem = 1
		}
		snafuSum = append(snafuSum, rem)
		take = quot
	}
	reverseMap := make(map[int]byte)
	for k, v := range charIntMap {
		reverseMap[v] = k
	}
	for i := range snafuSum {
		fmt.Printf("%c",
			reverseMap[snafuSum[len(snafuSum)-1-i]])
	}
}
