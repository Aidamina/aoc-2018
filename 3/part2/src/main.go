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
	var lines []string
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var maxWidth, maxHeight int
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
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
	for _, line := range lines {
		match := r.FindSubmatch([]byte(line))
		id, _ := strconv.Atoi(string(match[1]))
		x, _ := strconv.Atoi(string(match[2]))
		y, _ := strconv.Atoi(string(match[3]))
		width, _ := strconv.Atoi(string(match[4]))
		height, _ := strconv.Atoi(string(match[5]))
		var conflict bool
		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				if state[i][j] > 1 {
					conflict = true
				}
			}
		}
		if !conflict {
			fmt.Printf("%d\n", id)
		}
	}

}
