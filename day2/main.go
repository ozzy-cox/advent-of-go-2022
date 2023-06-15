package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Win  = 6
	Draw = 3
	Loss = 0
)

func charToOpType(c string) int {
	if c == "A" {
		return Rock
	}
	if c == "B" {
		return Paper
	}
	if c == "C" {
		return Scissors
	}
	return -1
}

func charToOwnType(c string) int {
	if c == "X" {
		return Rock
	}
	if c == "Y" {
		return Paper
	}
	if c == "Z" {
		return Scissors
	}
	return -1
}
func main() {

	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	score := 0
	for fileScanner.Scan() {
		hands := strings.Split(fileScanner.Text(), " ")

		opHand := charToOpType(hands[0])
		ownHand := charToOwnType(hands[1])

		switch opHand {
		case Rock:
			switch ownHand {
			case Rock:
				score += 4
			case Paper:
				score += 8
			case Scissors:
				score += 3
			}

		case Paper:
			switch ownHand {
			case Rock:
				score += 1
			case Paper:
				score += 5
			case Scissors:
				score += 9
			}
		case Scissors:
			switch ownHand {
			case Rock:
				score += 7
			case Paper:
				score += 2
			case Scissors:
				score += 6
			}
		}
	}

	fmt.Println(score)
}
