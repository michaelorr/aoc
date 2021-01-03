package days

import (
	"fmt"

	"orr.co/adventofcode/data"
)

func Day6() {
	c1 := make(chan []string)
	c2 := make(chan []string)
	finished := make(chan bool)
	go Day6Part1(c1, finished)
	go Day6Part2(c2, finished)

	for group := range data.AsChunkedSlices(data.Day(6)) {
		c1 <- group
		c2 <- group
	}

	close(c1)
	close(c2)
	<-finished
	<-finished
}

func Day6Part1(groups chan []string, finished chan bool) {
	var runningTotal int

	for group := range groups {
		groupAnswers := make(map[rune]bool)
		for _, member := range group {
			for _, char := range member {
				groupAnswers[char] = true
			}
		}
		runningTotal += len(groupAnswers)
		groupAnswers = make(map[rune]bool)
	}
	fmt.Printf("Part 1:\n\t%d\n", runningTotal)
	finished <- true
}

func Day6Part2(groups chan []string, finished chan bool) {
	var runningTotal int

	for group := range groups {
		groupAnswers := make(map[rune]int)
		var memberCount int
		for _, member := range group {
			for _, r := range member {
				var cur int
				var ok bool
				if cur, ok = groupAnswers[r]; !ok {
					cur = 0
				}
				groupAnswers[r] = cur + 1
			}
			memberCount++
		}
		for _, v := range groupAnswers {
			if v == memberCount {
				runningTotal++
			}
		}

		groupAnswers = make(map[rune]int)
		memberCount = 0
	}
	fmt.Printf("Part 2:\n\t%d\n", runningTotal)
	finished <- true
}
