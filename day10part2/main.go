package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("File cant read")
	}

	scanner := bufio.NewScanner(file)

	register := 1
	cycles := 1

	instrCyclesRemaining := 0
	var currentInstr []string

	for {
		if instrCyclesRemaining == 0 {
			// Commit current instr changes
			if len(currentInstr) != 0 {
				if currentInstr[0] != "noop" {
					instrValue, err := strconv.Atoi(currentInstr[1])
					if err != nil {
						panic("error at converting int")
					}
					register += instrValue
				}
			}

			hasLine := scanner.Scan()
			if !hasLine {
				break
			}
			currentInstr = strings.Split(scanner.Text(), " ")
			if currentInstr[0] == "addx" {
				instrCyclesRemaining = 2
			} else {
				instrCyclesRemaining = 1
			}
		}
		if (cycles)%40 == 1 {
			println()
		}
		drawnPixel := (cycles - 1) % 40
		if drawnPixel >= register-1 && drawnPixel <= register+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		instrCyclesRemaining--
		cycles++
	}
}
