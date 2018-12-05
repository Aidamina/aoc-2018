package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var runes []rune
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			runes = append(runes, c)
		}
	}
	var smallestSequence = -1
	for i := int(rune("a"[0])); i <= int(rune("z"[0])); i++ {
		var input []rune
		for _, r := range runes {
			if r != rune(i) && r != rune(i-32) {
				input = append(input, r)
			}
		}
		output := collapse(input)
		if smallestSequence == -1 || smallestSequence > len(output) {
			smallestSequence = len(output)
		}
	}
	fmt.Println(smallestSequence)
}

func collapse(input []rune) []rune {
	var output []rune
	var last rune = -1
	for _, r := range input {
		if last != -1 && math.Abs(float64(int(r)-int(last))) == 32 {
			output = output[:len(output)-1]
			if len(output) > 0 {
				last = output[len(output)-1]
			} else {
				last = -1
			}
		} else {
			output = append(output, r)
			last = r
		}
	}
	return output
}
