package days

import (
	"fmt"

	"orr.co/adventofcode/data"
)

func Day9() {
	finished := make(chan bool)

	var nums []int
	for num := range data.AsInts(data.Day(9)) {
		nums = append(nums, num)
	}
	go Day9Part1(nums, finished)
	go Day9Part2(nums, 552655238, finished)

	<-finished
	<-finished
}

func Day9Part1(in []int, finished chan bool) {
	window := 25
	ptr := 25

	for {
		matchFound := false
		window := in[ptr-window : ptr]
		target := in[ptr]

	Loop:
		for i := range window {
			for j := i + 1; j < len(window); j++ {
				if window[i]+window[j] == target {
					matchFound = true
					break Loop
				}
			}
		}
		if !matchFound {
			break
		}

		ptr++
	}

	fmt.Printf("Part 1:\n\t%d\n", in[ptr])
	finished <- true
}

func Day9Part2(in []int, target int, finished chan bool) {
	var out int
Loop:
	for start := range in {
		for end := start + 1; end < len(in); end++ {
			sum := Sum(in[start:end])
			if sum > target {
				continue Loop
			}

			if sum == target {
				out = Answer(in[start:end])
				break Loop
			}
		}

	}

	fmt.Printf("Part 2:\n\t%d\n", out)
	finished <- true
}

func Answer(data []int) int {
	smallest := data[0]
	largest := data[0]

	for _, n := range data {
		if n < smallest {
			smallest = n
		}
		if n > largest {
			largest = n
		}
	}
	return smallest + largest
}

func Sum(data []int) (total int) {
	for _, n := range data {
		total += n
	}
	return
}
