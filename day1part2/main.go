package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(file)
	var sum = 0
	maxes := []int{0, 0, 0}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			if sum > maxes[0] {
				temp0 := maxes[0]
				temp1 := maxes[1]
				maxes[0] = sum
				maxes[1] = temp0
				maxes[2] = temp1
			} else if sum > maxes[1] {
				temp1 := maxes[1]
				maxes[1] = sum
				maxes[2] = temp1
			} else if sum > maxes[2] {
				maxes[2] = sum
			}
			sum = 0
		} else {
			value, err := strconv.Atoi(line)
			check(err)
			sum += value
		}
	}

	fmt.Println(maxes)
}
