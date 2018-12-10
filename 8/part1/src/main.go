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
	for i := 0; i < nodeCount; i++ {
		sum += readNode(messages)
	}
	for i := 0; i < metadataCount; i++ {
		sum += <-messages
	}
	return sum
}
