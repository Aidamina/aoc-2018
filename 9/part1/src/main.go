package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	r := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	line := scanner.Text()
	match := r.FindSubmatch([]byte(line))
	players, _ := strconv.Atoi(string(match[1]))
	marbles, _ := strconv.Atoi(string(match[2]))
	var score = map[int]int{}
	score[0] = 1
	var state []int
	var current = 1
	var remove = func(index int) {
		copy(state[index:], state[index+1:])
		state[len(state)-1] = 0
		state = state[:len(state)-1]

	}
	var insert = func(index int, value int) {

		state = append(state, 0)
		copy(state[index+1:], state[index:])
		state[index] = value
	}
	// Insert the first marbles manually
	insert(0, 0)
	insert(1, 1)
	var find = func(clockwise bool, distance int) int {
		var modifier int
		if clockwise {
			modifier = 1
		} else {
			modifier = -1
		}
		index := (current + (distance * modifier)) % len(state)
		for index < 0 {
			index += len(state)
		}
		return index
	}
	var marble = 2
	var player = 1
	for marble <= marbles {
		if marble%23 == 0 {
			score[player] += marble
			var index = find(false, 7)
			score[player] += state[index]
			remove(index)
			current = index
		} else {
			var index = find(true, 2)
			if index == 0 {
				index = len(state)
			}
			insert(index, marble)
			current = index
		}
		player++
		player %= players
		marble++
	}
	var highScore int
	for _, val := range score {
		if highScore < val {
			highScore = val
		}
	}
	fmt.Printf("%d\n", highScore)

}
