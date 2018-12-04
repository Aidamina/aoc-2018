package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	r := regexp.MustCompile("#(?P<id>\\d+) @ (?P<x>\\d+),(?P<y>\\d+): (?P<width>\\d+)x(?P<height>\\d+)")
	var state = map[int]map[int]int{}
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var maxWidth, maxHeight int
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindSubmatch([]byte(line))
		x, _ := strconv.Atoi(string(match[2]))
		y, _ := strconv.Atoi(string(match[3]))
		width, _ := strconv.Atoi(string(match[4]))
		height, _ := strconv.Atoi(string(match[5]))
		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				if _, ok := state[i]; !ok {
					state[i] = map[int]int{}
				}
				if i > maxWidth {
					maxWidth = i
				}
				if j > maxHeight {
					maxHeight = j
				}
				state[i][j]++
			}
		}
	}
	var overlapCount int64
	for i := 0; i <= maxWidth; i++ {
		for j := 0; j <= maxHeight; j++ {
			if _, ok := state[i]; ok {
				if state[i][j] > 1 {
					overlapCount++
				}
			}
		}
	}
	fmt.Printf("%d\n", overlapCount)
}
