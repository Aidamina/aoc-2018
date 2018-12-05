package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EventType int

const (
	Duty EventType = 0 + iota
	Sleep
	Awake
)

type Event struct {
	id        int
	eventType EventType
	time      time.Time
}

func main() {
	r := regexp.MustCompile(`\[(\d+)-(\d+)-(\d+) (\d+):(\d+)\] (.+#(\d+).+|.+)`)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var events []*Event

	// Make a list of events
	for scanner.Scan() {
		var id int
		var eventType EventType
		line := scanner.Text()
		match := r.FindSubmatch([]byte(line))
		year, _ := strconv.Atoi(string(match[1]))
		month, _ := strconv.Atoi(string(match[2]))
		day, _ := strconv.Atoi(string(match[3]))
		hour, _ := strconv.Atoi(string(match[4]))
		minute, _ := strconv.Atoi(string(match[5]))
		data := string(match[6])
		if strings.HasPrefix(data, "Guard") {
			id, _ = strconv.Atoi(string(match[7]))
			eventType = Duty
		} else if data == "falls asleep" {
			eventType = Sleep
		} else if data == "wakes up" {
			eventType = Awake
		}
		event := &Event{
			id:        id,
			eventType: eventType,
			time:      time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
		}
		events = append(events, event)
	}
	sort.Slice(events, func(i, j int) bool { return events[i].time.Before(events[j].time) })

	// Fill in the missing ids for events
	var lastDutyId int
	for _, event := range events {
		if event.eventType == Duty {
			lastDutyId = event.id
		} else {
			event.id = lastDutyId
		}
	}

	// Map minutes to guards
	var state = map[int]map[int]int{}
	var sleepTime time.Time
	for _, event := range events {
		if _, ok := state[event.id]; !ok {
			state[event.id] = map[int]int{}
		}
		if event.eventType == Awake {
			wakeTime := event.time.Minute()
			for i := sleepTime.Minute(); i < wakeTime; i++ {
				state[event.id][i]++
			}
		}
		if event.eventType == Sleep {
			sleepTime = event.time
		}
	}

	// Find the guard and minute with the highest count
	var highestMinuteCount int
	var highestMinute int
	var highestId int
	for id, minutes := range state {
		for minute, count := range minutes {
			if count > highestMinuteCount {
				highestMinuteCount = count
				highestMinute = minute
				highestId = id
			}
		}
	}
	fmt.Println(highestId * highestMinute)
}
