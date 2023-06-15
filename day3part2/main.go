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
		lines := [3]string{}

		lines[0] = fileScanner.Text()
		fileScanner.Scan()
		lines[1] = fileScanner.Text()
		fileScanner.Scan()
		lines[2] = fileScanner.Text()

		itemSet1 := make(map[byte]struct{})
		itemSet2 := make(map[byte]struct{})

		for i := range lines[0] {
			itemSet1[lines[0][i]] = struct{}{}
		}

		for i := range lines[1] {
			itemSet2[lines[1][i]] = struct{}{}
		}

		for i := range lines[2] {
			_, ok1 := itemSet1[lines[2][i]]
			_, ok2 := itemSet2[lines[2][i]]

			if ok1 && ok2 {
				char := lines[2][i]
				if char > 96 {
					fmt.Println(char - 96)
					sum += int(char) - 96
				} else {
					fmt.Println(char - 64 + 26)
					sum += int(char) - 64 + 26
				}
				fmt.Printf("%c\n", char)
				break
			}
		}
	}
	println(sum)
}
