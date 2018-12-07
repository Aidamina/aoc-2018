package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type Step struct {
	id   string
	todo int
	next []*Step
	prev []*Step
}

func taskLength(id string) int {
	return int([]rune(id)[0]) - 4
}

func main() {
	r := regexp.MustCompile(`Step (.+) must be finished before step (.+) can begin.`)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var steps = map[string]*Step{}

	var getStep = func(id string) *Step {
		var step = steps[id]
		if step == nil {
			step = &Step{
				id:   id,
				todo: taskLength(id),
			}
			steps[id] = step
		}
		return step
	}

	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindSubmatch([]byte(line))
		firstId := string(match[1])
		secondId := string(match[2])
		first := getStep(firstId)
		second := getStep(secondId)
		first.next = append(first.next, second)
		second.prev = append(second.prev, first)
	}
	var second int
	for {
		var ready []string
		for id, step := range steps {
			if step.todo > 0 {
				var preqMet = true
				for _, p := range step.prev {
					if p.todo > 0 {
						preqMet = false
					}
				}
				if preqMet {
					ready = append(ready, id)
				}
			}
		}
		if len(ready) > 0 {
			sort.Strings(ready)
			for i := 0; i < 5 && i < len(ready); i++ {
				steps[ready[i]].todo--
			}
			second++
		} else {
			break
		}

	}

	fmt.Printf("%d\n", second)
}
