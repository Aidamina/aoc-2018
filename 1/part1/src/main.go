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
	var frequency int
	for scanner.Scan() {
		delta, err := strconv.Atoi(scanner.Text())
		if err == nil {
			frequency += delta
		}
	}
	fmt.Println(frequency)
}
