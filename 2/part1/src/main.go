package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var twos int
	var threes int
	for scanner.Scan() {
		text := scanner.Text()
		var chars = map[rune]int{}
		for _, rune := range text {
			chars[rune]++
		}
		var foundTwo, foundThree bool
		for _, val := range chars {
			if val == 2 {
				foundTwo = true
			}
			if val == 3 {
				foundThree = true
			}
		}
		if foundTwo {
			twos++
		}
		if foundThree {
			threes++
		}
	}
	fmt.Println(twos * threes)

}
