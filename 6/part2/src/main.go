package main

import (
	"bufio"
	"fmt"
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
	var total = 10000
	var buffer = total / len(points)
	var count int
	for x := minX - buffer; x <= maxX+buffer; x++ {
		for y := minY - buffer; y <= maxY+buffer; y++ {
			point1 := NewPoint(x, y)
			var sum int
			for _, point2 := range points {
				sum += distance(point1, point2)
			}
			if sum < total {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
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
