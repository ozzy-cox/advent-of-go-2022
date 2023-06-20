package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point [2]int

func (p Point) x() int {
	return p[0]
}

func (p Point) y() int {
	return p[1]
}

type Node struct {
	sensor  Point
	mradius int
}

func getMradius(n1 Point, n2 Point) int {
	return int(math.Abs(float64(n1.x()-n2.x())) + math.Abs(float64(n2.y()-n1.y())))
}

func getCoveredRange(y int, p Point, radius int) (int, int) {
	dist := int(math.Abs(float64(y - p.y())))
	remaining := radius - dist
	return p.x() - remaining, p.x() + remaining
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	nodes := make([]Node, 0)

	coveredPoints := make(map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")

		sensorParts := strings.Split(strings.Split(lineParts[0], "at ")[1], ", ")
		beaconParts := strings.Split(strings.Split(lineParts[1], "at ")[1], ", ")

		sensorX, _ := strconv.Atoi(sensorParts[0][2:])
		sensorY, _ := strconv.Atoi(sensorParts[1][2:])
		beaconX, _ := strconv.Atoi(beaconParts[0][2:])
		beaconY, _ := strconv.Atoi(beaconParts[1][2:])

		sensor := Point{sensorX, sensorY}
		beacon := Point{beaconX, beaconY}
		node := Node{
			sensor:  sensor,
			mradius: getMradius(sensor, beacon),
		}
		nodes = append(nodes, node)
		x1, x2 := getCoveredRange(2_000_000, sensor, node.mradius)
		for i := x1; i < x2; i++ {
			coveredPoints[i] = true
		}
	}
	fmt.Println(len(coveredPoints))
}
