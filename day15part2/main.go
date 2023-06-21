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

type Line struct {
	m int
	q int
}

func getMdistance(n1 Point, n2 Point) int {
	return int(math.Abs(float64(n1.x()-n2.x())) + math.Abs(float64(n1.y()-n2.y())))
}

func getCoveredRange(y int, p Point, radius int) (int, int) {
	dist := int(math.Abs(float64(y - p.y())))
	remaining := radius - dist
	return p.x() - remaining, p.x() + remaining
}

func isEmpty(p Point, nodes []Node) bool {
	for _, node := range nodes {
		if node.mradius >= getMdistance(node.sensor, p) {
			return false
		}
	}

	return true
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
			mradius: getMdistance(sensor, beacon),
		}
		nodes = append(nodes, node)
		x1, x2 := getCoveredRange(2_000_000, sensor, node.mradius)
		for i := x1; i < x2; i++ {
			coveredPoints[i] = true
		}
	}

	lines := make([]Line, 0)
	lineCounts := make(map[Line]int)

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		lines = append(lines, Line{m: 1, q: node.sensor.y() - node.mradius - 1 - node.sensor.x()})
		lines = append(lines, Line{m: -1, q: node.sensor.y() - node.mradius - 1 + node.sensor.x()})
		lines = append(lines, Line{m: 1, q: node.sensor.y() + node.mradius + 1 - node.sensor.x()})
		lines = append(lines, Line{m: -1, q: node.sensor.y() + node.mradius + 1 + node.sensor.x()})

		for j := 0; j < 4; j++ {
			line := lines[len(lines)-j-1]
			_, ok := lineCounts[line]
			if ok {
				lineCounts[line]++
			} else {
				lineCounts[line] = 1
			}
		}
	}

	risingLines := make([]Line, 0)
	descLines := make([]Line, 0)

	for v, count := range lineCounts {
		if count > 1 {
			if v.m < 0 {
				risingLines = append(risingLines, v)
			} else {
				descLines = append(descLines, v)
			}
		}
	}

	intersectionPoints := make([]Point, 0)

	for _, rl := range risingLines {
		for _, dl := range descLines {
			x := (rl.q - dl.q) / 2
			y := x + dl.q
			intersectionPoints = append(intersectionPoints, Point{x, y})
		}
	}

	limit := 4_000_000

	for _, point := range intersectionPoints {
		if isEmpty(point, nodes) {
			fmt.Println(point.x()*limit + point.y())
			break
		}
	}
}
