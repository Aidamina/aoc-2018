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
	var fuelCells = map[int]map[int]int{}
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			if _, ok := fuelCells[y]; !ok {
				fuelCells[y] = map[int]int{}
			}
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
	var highestX, highestY, highestOutput = -1, -1, -1
	for y := 1; y <= 300-2; y++ {
		for x := 1; x <= 300-2; x++ {
			var output int
			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					output += fuelCells[i][j]
				}
			}
			if output > highestOutput {
				highestX = x
				highestY = y
				highestOutput = output
			}
		}
	}
	fmt.Printf("%d,%d\n", highestX, highestY)
}
