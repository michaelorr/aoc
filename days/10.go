package days

import (
	"fmt"
	"sort"

	"orr.co/adventofcode/data"
)

func Day10() {
	finished := make(chan bool)

	var nums []int
	for num := range data.AsInts(data.Day(10)) {
		nums = append(nums, num)
	}

	sort.Ints(nums)
	go Day10Part1(nums, finished)
	go Day10Part2(nums, finished)

	<-finished
	<-finished
}

func Day10Part1(in []int, finished chan bool) {
	var out, ones, threes, jolts int
	threes = 1

	for _, i := range in {
		switch i - jolts {
		case 1:
			ones++
		case 3:
			threes++
		}
		jolts = i
	}
	out = ones * threes
	fmt.Printf("Part 1:\n\tOnes: %d\n\tThrees: %d\n\t%d\n", ones, threes, out)
	finished <- true
}

func Day10Part2(in []int, finished chan bool) {
	var out int

	paths := make(map[int]int)
	paths[0] = 1
	for _, j := range in {
		paths[j] = paths[j-1] + paths[j-2] + paths[j-3]
	}

	out = paths[in[len(in)-1]]
	fmt.Printf("Part 2:\n\t%d\n", out)
	finished <- true
}
