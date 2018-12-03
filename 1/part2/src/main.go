package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var deltas []int
	for scanner.Scan() {
		delta, err := strconv.Atoi(scanner.Text())
		if err == nil {
			deltas = append(deltas, delta)
		}
	}
	var seen = map[int]bool{0: true}
	var frequency int
	for {
		for _, delta := range deltas {
			frequency += delta
			if _, ok := seen[frequency]; ok {
				fmt.Println(frequency)
				os.Exit(0)
			}
			seen[frequency] = true
		}
	}
}
