package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yourbasic/bit"
)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)

	count := 0
	for reader.Scan() {
		pairRanges := strings.Split(reader.Text(), ",")

		rangeStr1 := strings.Split(pairRanges[0], "-")

		range11, _ := strconv.Atoi(rangeStr1[0])
		range12, _ := strconv.Atoi(rangeStr1[1])

		rangeStr2 := strings.Split(pairRanges[1], "-")

		range21, _ := strconv.Atoi(rangeStr2[0])
		range22, _ := strconv.Atoi(rangeStr2[1])

		bitSet1 := new(bit.Set).AddRange(range11, range12+1) // {0..99}
		bitSet2 := new(bit.Set).AddRange(range21, range22+1) // {0..99}

		if tempbs := bitSet1.Or(bitSet2); tempbs.Equal(bitSet1) || tempbs.Equal(bitSet2) {

			fmt.Println(tempbs)
			fmt.Println(bitSet1)
			fmt.Println(bitSet2)
			count += 1
		}
	}
	fmt.Println(count)

}
