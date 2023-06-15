package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)
	sum := 0

	for fileScanner.Scan() {
		itemSet := make(map[byte]struct{})
		line := fileScanner.Text()
		for i := 0; i < len(line)/2; i++ {
			itemSet[line[i]] = struct{}{}
		}
		for i := len(line) / 2; i < len(line); i++ {
			if _, ok := itemSet[line[i]]; ok {
				char := line[i]
				fmt.Printf("%c\n", line[i])
				fmt.Println(line)

				if char > 96 {
					fmt.Println(char - 96)
					sum += int(char) - 96
				} else {
					fmt.Println(char - 64 + 26)
					sum += int(char) - 64 + 26
				}
				break

			}
		}
	}
	println(sum)
}
