package days

import (
	"fmt"

	"orr.co/adventofcode/data"
)

const (
	F0 uint8 = 0b1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
)

func Day5() {
	c1 := make(chan int)
	c2 := make(chan int)
	finished := make(chan bool)
	go Day5Part1(c1, finished)
	go Day5Part2(c2, finished)

	for boardingPass := range data.Day(5) {
		id := SeatID(parseRow(boardingPass), parseColumn(boardingPass))
		c1 <- id
		c2 <- id
	}

	close(c1)
	close(c2)
	<-finished
	<-finished
}

func Day5Part1(ids chan int, finished chan bool) {
	var largestID int

	for id := range ids {
		if id > largestID {
			largestID = id
		}
	}
	fmt.Printf("Part 1:\n\tLargest ID: %d\n", largestID)
	finished <- true
}

func Day5Part2(ids chan int, finished chan bool) {
	seatIDs := make(map[int]bool)

	for id := range ids {
		seatIDs[id] = true
	}

	var i int
	for i = 0; i < len(seatIDs); i++ {
		if seatIDs[i] == false && seatIDs[i+1] == true && seatIDs[i-1] == true {
			fmt.Printf("Part 2:\n\tSeat ID: %d\n", i)
		}
	}

	finished <- true
}

func SeatID(row, column int) int {
	return row*8 + column
}

func parseRow(boardingPass string) int {
	var row uint8

	for i, flag := range []uint8{F6, F5, F4, F3, F2, F1, F0} {
		if string(boardingPass[i]) == "B" {
			row = row | flag
		}
	}

	return int(row)
}

func parseColumn(boardingPass string) int {
	var col uint8

	for i, flag := range []uint8{F2, F1, F0} {
		if string(boardingPass[i+7]) == "R" {
			col = col | flag
		}
	}

	return int(col)
}
