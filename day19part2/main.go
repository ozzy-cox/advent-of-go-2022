package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Blueprint [4][4]int

// type Blueprint int

func parse() (*[]Blueprint, [][3]int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	blueprints := make([]Blueprint, 0)
	costIndices := [...]int{6, 12, 18, 21, 27, 30}
	maxSpends := make([][3]int, 0)
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
			[4]int{costs[1], 0, 0, 0},
			[4]int{costs[2], costs[3], 0, 0},
			[4]int{costs[4], 0, costs[5], 0},
		}
		maxSpend := [3]int{}

		for _, recipe := range blueprint {
			for j, res := range recipe {
				if j <= 3 && res > 0 {
					maxSpend[j] = int(math.Max(float64(maxSpend[j]), float64(recipe[j])))
				}
			}
		}
		blueprints = append(blueprints, blueprint)
		maxSpends = append(maxSpends, maxSpend)
	}

	return &blueprints, maxSpends
}

type State struct {
	resources [4]int
	robots    [4]int
	time      int
}

func simulateBlueprint(bp Blueprint, maxSpend [3]int, robots [4]int, resources [4]int, rounds int, cache map[State]int) int {
	if rounds == 0 {
		return resources[3]
	}

	key := State{
		resources: resources,
		robots:    robots,
		time:      rounds,
	}
	if val, ok := cache[key]; ok {
		return val
	}

	maxVal := resources[3] + robots[3]*rounds
	if maxVal > 10 {
		fmt.Print()
	}

	for i, recipe := range bp {
		if i != 3 && robots[i] >= maxSpend[i] {
			continue
		}

		wait := 0

		broke := false
		for j, res := range recipe {
			if res > 0 {
				if robots[j] == 0 {
					broke = true
					break
				}
				necessaryRes := float64(res-resources[j]) / float64(robots[j])
				wait = int(math.Max(float64(wait), math.Ceil(necessaryRes)))
			}
		}
		if !broke {
			remtime := rounds - wait - 1
			if remtime <= 0 {
				// Maybe SUS !!!
				continue
			}

			resources_ := resources
			robots_ := robots
			for k := 0; k < 4; k++ {
				resources_[k] = resources_[k] + robots[k]*(wait+1)
			}

			for k := 0; k < 4; k++ {
				resources_[k] -= recipe[k]
			}
			robots_[i] += 1

			for k := 0; k < 3; k++ {
				resources_[k] = int(math.Min(float64(resources_[k]), float64(maxSpend[k]*remtime)))
			}

			maxVal = int(math.Max(float64(maxVal), float64(simulateBlueprint(bp, maxSpend, robots_, resources_, remtime, cache))))
		}

	}

	cache[key] = maxVal
	return maxVal
}

func main() {
	blueprints, maxSpend := parse()

	sum := 1
	for i, bp := range *blueprints {
		if i < 3 {
			resources := [4]int{}
			robots := [4]int{1, 0, 0, 0}
			resultCache := make(map[State]int)
			r := simulateBlueprint(bp, maxSpend[i], robots, resources, 32, resultCache)
			sum *= r
		}
	}
	fmt.Println(sum)
}
