package minclsr

import "sort"

// Event type
type Event struct {
	time  int
	start bool
}

// MinClassrooms function
func MinClassrooms(schedule [][]int) int {
	events := make([]Event, 0, len(schedule)*2)
	for _, inter := range schedule {
		events = append(events, Event{inter[0], true})
		events = append(events, Event{inter[1], false})
	}

	sort.SliceStable(events, func(i, j int) bool {
		return events[i].time < events[j].time
	})

	n := 0
	m := 0

	for _, e := range events {
		if e.start {
			n++
			m = max(m, n)
		} else {
			n--
		}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
