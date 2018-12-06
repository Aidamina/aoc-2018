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

func NewPoint(x int, y int) Point {
	return [2]int{x, y}
}

func (point Point) coords() (int, int) {
	return point[0], point[1]
}
func (point Point) x() int {
	return point[0]
}
func (point Point) y() int {
	return point[1]
}
func (point Point) String() string {
	return strconv.Itoa(point[0]) + ", " + strconv.Itoa(point[1])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var points []Point
	for scanner.Scan() {
		var split = strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points = append(points, NewPoint(x, y))
	}
	var minX, minY, maxX, maxY int
	for i, point := range points {
		x, y := point.coords()
		if x < minX || i == 0 {
			minX = x
		}
		if y < minY || i == 0 {
			minY = y
		}
		if x > maxX || i == 0 {
			maxX = x
		}
		if y > maxY || i == 0 {
			maxY = y
		}
	}
	var disqualified []Point
	var count = map[Point]int{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			point1 := NewPoint(x, y)
			var distances = map[Point]int{}
			for _, point2 := range points {
				distances[point2] = distance(point1, point2)
			}
			var closestPoint Point
			var smallestDistance = math.MaxInt32
			var conflict bool
			for point, distance := range distances {
				if distance < smallestDistance {
					smallestDistance = distance
					closestPoint = point
					conflict = false
				} else if distance == smallestDistance {
					conflict = true
				}
			}
			if conflict {
				continue
			}
			if x == minX || x == maxX || y == minY || y == maxY {
				disqualified = append(disqualified, closestPoint)
			}
			count[closestPoint]++
		}
	}
	for _, disq := range disqualified {
		delete(count, disq)
	}
	var largestCount int
	for _, c := range count {
		if c > largestCount {
			largestCount = c
		}
	}
	fmt.Printf("%d\n", largestCount)

}
func distance(point1 Point, point2 Point) int {
	return abs(point1.x()-point2.x()) + abs(point1.y()-point2.y())
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
