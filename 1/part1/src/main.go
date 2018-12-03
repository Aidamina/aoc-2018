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
	var total int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			total += i
		}
	}
	fmt.Println(total)
}
