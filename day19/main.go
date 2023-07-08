package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Blueprint [4][4]int

// type Blueprint int

func parse() []Blueprint {
	file, _ := os.Open("smallinput.txt")
	scanner := bufio.NewScanner(file)
	blueprints := make([]Blueprint, 0)
	costIndices := [...]int{6, 12, 18, 21, 27, 30}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		costs := [6]int{}
		for i := 0; i < len(costIndices); i++ {
			parsed, _ := strconv.Atoi(tokens[costIndices[i]])
			costs[i] = parsed
		}

		blueprint := Blueprint{
			[4]int{costs[0], 0, 0, 0},
			[4]int{costs[2], 0, 0, 0},
			[4]int{costs[2], costs[3], 0, 0},
			[4]int{costs[4], 0, costs[5], 0},
		}
		blueprints = append(blueprints, blueprint)
	}

	return blueprints
}

func add(a [4]int, b [4]int) [4]int {
	sum := [4]int{}
	for i := 0; i < 4; i++ {
		sum[i] = a[i] + b[i]
	}
	return sum
}

func subtract(a [4]int, b [4]int) [4]int {
	diff := [4]int{}
	for i := 0; i < 4; i++ {
		diff[i] = a[i] - b[i]
	}
	return diff
}

func isGE(a [4]int, b [4]int) bool {
	for i := 0; i < 4; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

func max(values []int) int {
	max := 0
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func simulateBlueprint(bp Blueprint, robots [4]int, resources [4]int, rounds int, cache map[[9]int]int) int {
	if rounds == 0 {
		if resources[3] > 7 {
			fmt.Println(resources[3])
		}
		return resources[3]
	}

	// if cachedVal, ok := cache[[9]int{
	// 	robots[0],
	// 	robots[1],
	// 	robots[2],
	// 	robots[3],
	// 	resources[0],
	// 	resources[1],
	// 	resources[2],
	// 	resources[3],
	// 	rounds,
	// }]; ok {
	// 	return cachedVal
	// }

	results := make([]int, 0)
	resourcesToAdd := add(robots, [4]int{})
	// results = append(results, simulateBlueprint(bp, robots, resources, rounds-1, cache))
	for j := 3; j >= 0; j-- {
		cost := bp[j]
		for isGE(resources, cost) {
			results = append(results, simulateBlueprint(bp, robots, resources, rounds-1, cache))
			robots[j] += 1
			resources = subtract(resources, cost)
			results = append(results, simulateBlueprint(bp, robots, resources, rounds-1, cache))
		}
	}
	resources = add(resources, resourcesToAdd)
	results = append(results, simulateBlueprint(bp, robots, resources, rounds-1, cache))

	// Add to resources
	fmt.Println(robots)
	cache[[9]int{
		robots[0],
		robots[1],
		robots[2],
		robots[3],
		resources[0],
		resources[1],
		resources[2],
		resources[3],
		rounds,
	}] = max(results)
	return max(results)
}

func main() {
	blueprints := parse()
	resources := [4]int{}
	robots := [4]int{1, 0, 0, 0}
	resultCache := make(map[[9]int]int)
	r := simulateBlueprint(blueprints[1], robots, resources, 27, resultCache)
	fmt.Println(r)
}
