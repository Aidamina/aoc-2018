package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	r := regexp.MustCompile(`initial state: ([#\.]+)`)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var currentGeneration = map[int]bool{}
	// Get initial
	scanner.Scan()
	line := scanner.Text()
	match := r.FindSubmatch([]byte(line))
	initialStr := string(match[1])
	var min, max = 0, len(initialStr) - 1
	for index, r := range initialStr {
		if r == '#' {
			currentGeneration[index] = true
		}
	}
	r = regexp.MustCompile(`([#\.]{5}) => ([#\.])`)
	var operations = map[string]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindSubmatch([]byte(line))
		if match == nil {
			continue
		}
		input := string(match[1])
		output := string(match[2])
		if output == "#" {
			operations[input] = true
		}
	}
	for generation := 0; generation < 20; generation++ {
		var nextGeneration = map[int]bool{}
		var newMin, newMax = math.MaxInt32, math.MinInt32
		for i := min - 2; i <= max+2; i++ {
			var sb strings.Builder
			for j := i - 2; j <= i+2; j++ {
				if currentGeneration[j] {
					sb.WriteRune('#')
				} else {
					sb.WriteRune('.')
				}
			}
			state := sb.String()
			if operations[state] {
				nextGeneration[i] = true
				if i < newMin {
					newMin = i
				}
				if i > newMax {
					newMax = i
				}
			}
		}
		max = newMax
		min = newMin
		currentGeneration = nextGeneration
	}
	var total int
	for i, val := range currentGeneration {
		if val {
			total += i
		}

	}
	fmt.Printf("%d\n", total)

}
