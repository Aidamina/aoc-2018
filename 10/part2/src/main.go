package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Vec2 [2]int

func NewVec2(x int, y int) Vec2 {
	return [2]int{x, y}
}

func (v Vec2) coords() (int, int) {
	return v[0], v[1]
}
func (v Vec2) x() int {
	return v[0]
}
func (v Vec2) y() int {
	return v[1]
}
func (v Vec2) String() string {
	return strconv.Itoa(v[0]) + ", " + strconv.Itoa(v[1])
}

type Star struct {
	position Vec2
	velocity Vec2
	steps    int
}

func (star *Star) Steps() int {
	return star.steps

}

func (star *Star) Step(steps int) {
	star.position = NewVec2(star.position.x()+(steps*star.velocity.x()), star.position.y()+(steps*star.velocity.y()))
	star.steps += steps
}

func main() {
	var stars []*Star
	r := regexp.MustCompile(`position=<\s?(-?\d+), \s?(-?\d+)> velocity=<\s?(-?\d+), \s?(-?\d+)>`)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	// Make a list of events
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindSubmatch([]byte(line))
		posX, _ := strconv.Atoi(string(match[1]))
		posY, _ := strconv.Atoi(string(match[2]))
		velX, _ := strconv.Atoi(string(match[3]))
		velY, _ := strconv.Atoi(string(match[4]))
		star := &Star{
			position: NewVec2(posX, posY),
			velocity: NewVec2(velX, velY),
		}
		stars = append(stars, star)
	}
	var smallestSize = 1000000
	var second = 0
	for {
		var minX, minY, maxX, maxY int
		for i, star := range stars {
			star.Step(1)
			second++
			x, y := star.position.coords()
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
		size := (maxY - minY) + (maxX - minX)
		if size < smallestSize {
			smallestSize = size
		} else {
			break
		}
		for _, star := range stars {
			star.Step(1)
			second++
		}
	}
	for _, star := range stars {
		star.Step(-1)
	}
	fmt.Printf("%d\n", stars[0].Steps())
}
