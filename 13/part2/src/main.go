package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Direction rune

const (
	Up    Direction = '^'
	Down  Direction = 'v'
	Left            = '<'
	Right           = '>'
)

type Cart struct {
	direction Direction
	x         int
	y         int
	state     int
	crashed   bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var world = map[int]map[int]rune{}
	var carts []*Cart
	var addCart = func(direction Direction, x int, y int) {
		carts = append(carts, &Cart{
			direction: direction,
			x:         x,
			y:         y,
		})
	}

	var y int
	for scanner.Scan() {
		world[y] = map[int]rune{}
		line := scanner.Text()
		for x, r := range line {
			if r == rune(Right) || r == rune(Left) || r == rune(Up) || r == rune(Down) {
				addCart(Direction(r), x, y)
				if r == rune(Up) || r == rune(Down) {
					r = '|'
				}
				if r == rune(Right) || r == rune(Left) {
					r = '-'
				}
			}
			if r != ' ' {
				world[y][x] = r
			}
		}
		y++
	}
	var instructions = map[rune]map[Direction]Direction{}
	instructions['\\'] = map[Direction]Direction{}
	instructions['\\'][Up] = Left
	instructions['\\'][Down] = Right
	instructions['\\'][Left] = Up
	instructions['\\'][Right] = Down
	instructions['/'] = map[Direction]Direction{}
	instructions['/'][Up] = Right
	instructions['/'][Down] = Left
	instructions['/'][Left] = Down
	instructions['/'][Right] = Up
	var instructionsCrossing = map[int]map[Direction]Direction{}
	// Left the first time
	instructionsCrossing[0] = map[Direction]Direction{}
	instructionsCrossing[0][Up] = Left
	instructionsCrossing[0][Down] = Right
	instructionsCrossing[0][Left] = Down
	instructionsCrossing[0][Right] = Up
	// Right the third time
	instructionsCrossing[2] = map[Direction]Direction{}
	instructionsCrossing[2][Up] = Right
	instructionsCrossing[2][Down] = Left
	instructionsCrossing[2][Left] = Up
	instructionsCrossing[2][Right] = Down

	var findNextDirection = func(cart *Cart) Direction {
		// Default to keep current direction
		var direction = cart.direction
		r := world[cart.y][cart.x]
		if in, ok := instructions[r]; ok {
			direction = in[cart.direction]
		}

		if r == '+' {
			state := cart.state % 3
			// Straight the second time
			if state != 1 {
				direction = instructionsCrossing[state][cart.direction]
			}
			cart.state++
		}

		return direction

	}
	var findNext = func(cart *Cart) (int, int) {
		var x, y int
		if cart.direction == Up {
			x = cart.x
			y = cart.y - 1
		}
		if cart.direction == Down {
			x = cart.x
			y = cart.y + 1
		}
		if cart.direction == Left {
			x = cart.x - 1
			y = cart.y
		}
		if cart.direction == Right {
			x = cart.x + 1
			y = cart.y
		}
		return x, y
	}

	var findCart = func(x int, y int) *Cart {
		for _, cart := range carts {
			if !cart.crashed && cart.x == x && cart.y == y {
				return cart
			}
		}
		return nil
	}
	var cartLeft *Cart
	var tick int
	for {
		// Sort cart list in turn order
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y < carts[j].y {
				return true
			}
			if carts[i].y == carts[j].y && carts[i].x < carts[j].x {
				return true
			}
			return false
		})

		for _, cart := range carts {
			if cart.crashed {
				continue
			}
			var nextX, nextY = findNext(cart)
			other := findCart(nextX, nextY)
			if other != nil {
				fmt.Printf("crashed %d,%d\n", nextX, nextY)
				other.crashed = true
				cart.crashed = true
			}
			cart.x = nextX
			cart.y = nextY
			cart.direction = findNextDirection(cart)
		}
		var cartsLeft int
		for _, cart := range carts {
			if !cart.crashed {
				cartLeft = cart
				cartsLeft++
			}
		}
		if cartsLeft <= 1 {
			break
		}

		tick++
	}
	fmt.Printf("%d,%d\n", cartLeft.x, cartLeft.y)

}
