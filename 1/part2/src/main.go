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
	var reached = map[int]bool{0: true}
	var frequency int
	for scanner.Scan() {
		delta, err := strconv.Atoi(scanner.Text())
		if err == nil {
			frequency += delta
			if _, ok := reached[frequency]; ok {
				fmt.Println(frequency)
				os.Exit(0)
			}
			reached[frequency] = true
		}
	}
	os.Exit(1)
}
