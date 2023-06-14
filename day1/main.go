package main

import (
	"bufio"
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
	var max = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			value, err := strconv.Atoi(line)
			check(err)
			sum += value
		}
	}
	println(max)
}
