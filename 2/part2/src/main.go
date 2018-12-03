package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var ids = []string{}
	for scanner.Scan() {
		id := scanner.Text()
		ids = append(ids, id)
	}

	size := len(ids[0])
	for i := range make([]int, size) {
		var seen = map[string]bool{}
		for _, id := range ids {
			fid := id[0:i] + id[i+1:size]
			fmt.Println(id + " => " + fid)
			if _, ok := seen[fid]; ok {
				fmt.Println(fid)
				os.Exit(0)
			}
			seen[fid] = true
		}
	}
}
