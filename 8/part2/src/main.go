package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var inputs []int
	for scanner.Scan() {
		for _, in := range strings.Split(scanner.Text(), " ") {
			value, err := strconv.Atoi(in)
			if err == nil {
				inputs = append(inputs, value)
			}
		}
	}
	messages := make(chan int)
	go func() {
		for _, input := range inputs {
			messages <- input
		}
	}()
	var total = readNode(messages)
	fmt.Println(total)
}

func readNode(messages chan int) int {
	var sum int
	var nodeCount = <-messages
	var metadataCount = <-messages
	var nodes []int
	for i := 0; i < nodeCount; i++ {
		val := readNode(messages)
		nodes = append(nodes, val)
	}
	for i := 0; i < metadataCount; i++ {
		val := <-messages
		if nodeCount > 0 {
			index := val - 1
			if index >= 0 && index < nodeCount {
				sum += nodes[index]
			}
		} else {
			sum += val
		}
	}
	return sum
}
