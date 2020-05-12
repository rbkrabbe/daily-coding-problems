// Good morning! Here's your coding interview problem for today.

// This problem was asked by Snapchat.

// Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.

// For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

package twentyone

import "sort"

// Slot represents a timeslot with start and end time.
type Slot struct {
	Start int
	End   int
}

// Twentyone computes the maximum number of simultaneously running slots.
func Twentyone(slots []Slot) int {

	sort.Slice(slots, func(i, j int) bool { return slots[i].Start < slots[i].End })

	ends := make([]int, 0)
	ongoing := 1
	ongoingMax := 1
	ends = append(ends, slots[0].End)

	for _, e := range slots[1:] {
		ongoing++

		tmp := make([]int, len(ends))
		for _, end := range ends {
			if end <= e.Start {
				ongoing--
			} else {
				tmp = append(tmp, end)
			}
		}
		ends = append(tmp, e.End)

		if ongoing > ongoingMax {
			ongoingMax++
		}
	}

	return ongoingMax
}
