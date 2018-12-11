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
	scanner.Scan()
	serial, _ := strconv.Atoi(scanner.Text())
	var fuelCells [301][301]int
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			rackID := x + 10
			powerLevel := rackID * y
			powerLevel += serial
			powerLevel *= rackID
			strPowerLevel := strconv.Itoa(powerLevel)
			digit := 0
			if len(strPowerLevel) >= 3 {
				digit, _ = strconv.Atoi(string(strPowerLevel[len(strPowerLevel)-3]))
			}
			digit -= 5
			fuelCells[y][x] = digit
		}
	}
	var highestX, highestY, highestOutput, highestSize = -1, -1, -1, -1
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			for size := 1; size <= 300; size++ {
				if x+size > 301 || y+size > 301 {
					continue
				}
				var output int
				for i := y; i < y+size; i++ {
					for j := x; j < x+size; j++ {
						output += fuelCells[i][j]
					}
				}
				if output > highestOutput {
					highestX = x
					highestY = y
					highestOutput = output
					highestSize = size
				}
			}
		}
	}

	fmt.Printf("%d,%d,%d\n", highestX, highestY, highestSize)
}
