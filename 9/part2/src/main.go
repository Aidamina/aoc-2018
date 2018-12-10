package main

import (
	"bufio"
	"container/ring"
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
	marbles *= 100
	var score = map[int]int{}
	score[0] = 1
	var state = ring.New(1)
	state.Value = 0

	var insert = func(value int) {
		r := ring.New(1)
		r.Value = value
		state.Link(r)
		state = state.Move(1)
	}
	insert(1)

	var marble = 2
	var player = 1
	for marble <= marbles {
		if marble%23 == 0 {
			score[player] += marble
			state = state.Move(-8)
			score[player] += state.Next().Value.(int)
			state.Unlink(1)
			state = state.Move(1)

		} else {
			state = state.Move(1)
			insert(marble)
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
