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
	Lose = 0
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

func charToResultType(c string) int {
	if c == "X" {
		return Lose
	}
	if c == "Y" {
		return Draw
	}
	if c == "Z" {
		return Win
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
		result := charToResultType(hands[1])

		switch opHand {
		case Rock:
			switch result {
			case Win:
				score += 8
			case Draw:
				score += 4
			case Lose:
				score += 3
			}

		case Paper:
			switch result {
			case Win:
				score += 9
			case Draw:
				score += 5
			case Lose:
				score += 1
			}
		case Scissors:
			switch result {
			case Win:
				score += 7
			case Draw:
				score += 6
			case Lose:
				score += 2
			}
		}
	}

	fmt.Println(score)
}
