package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type Step struct {
	id       string
	complete bool
	next     []*Step
	prev     []*Step
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
				id: id,
			}
			steps[id] = step
		}
		return step
	}

	// Make a list of events
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

	for {
		var ready []string
		for id, step := range steps {
			if !step.complete {
				var preqMet = true
				for _, p := range step.prev {
					if !p.complete {
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
			//fmt.Printf("|")
			for _, id := range ready {
				steps[id].complete = true
				fmt.Printf("%s", id)
				break
			}
		} else {
			break
		}
	}

	fmt.Println()
}
