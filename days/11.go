package days

import (
	"fmt"
	"strings"

	"orr.co/adventofcode/data"
)

func Day11() {
	finished := make(chan bool)

	var lines [][]string
	for line := range data.Day(11) {
		lines = append(lines, strings.Split(line, ""))
	}

	go Day11Part1(lines, finished)
	go Day11Part2(lines, finished)

	<-finished
	<-finished
}

func Day11Part1(in [][]string, finished chan bool) {
	var out int

	seats := in
	done := false

	for !done {
		seats, done = applyRules(seats)
	}

	for _, row := range seats {
		for _, seat := range row {
			if seat == "#" {
				out++
			}
		}
	}
	fmt.Printf("Part 1:\n\t%d\n", out)
	finished <- true
}

func applyRules(seats [][]string) ([][]string, bool) {
	out := make([][]string, len(seats))
	for i := range out {
		out[i] = make([]string, len(seats[0]))
	}
	done := true

	for i, row := range seats {
		for j, space := range row {
			switch space {
			case ".":
				out[i][j] = "."
			case "L":
				if countAdjacent(seats, i, j) == 0 {
					out[i][j] = "#"
					done = false
				} else {
					out[i][j] = "L"
				}
			case "#":
				if countAdjacent(seats, i, j) >= 4 {
					out[i][j] = "L"
					done = false
				} else {
					out[i][j] = "#"
				}
			}
		}
	}
	return out, done
}

func countAdjacent(seats [][]string, i, j int) int {
	var count int
	if seatExists(seats, i-1, j-1) && seatTaken(seats, i-1, j-1) {
		count++
	}
	if seatExists(seats, i-1, j) && seatTaken(seats, i-1, j) {
		count++
	}
	if seatExists(seats, i-1, j+1) && seatTaken(seats, i-1, j+1) {
		count++
	}
	if seatExists(seats, i, j-1) && seatTaken(seats, i, j-1) {
		count++
	}
	if seatExists(seats, i, j+1) && seatTaken(seats, i, j+1) {
		count++
	}
	if seatExists(seats, i+1, j-1) && seatTaken(seats, i+1, j-1) {
		count++
	}
	if seatExists(seats, i+1, j) && seatTaken(seats, i+1, j) {
		count++
	}
	if seatExists(seats, i+1, j+1) && seatTaken(seats, i+1, j+1) {
		count++
	}
	return count
}

func seatExists(seats [][]string, i, j int) bool {
	return (i >= 0) && (i < len(seats)) && (j >= 0) && (j < len(seats[i]))
}

func seatTaken(seats [][]string, i, j int) bool {
	return seats[i][j] == "#"
}

func Day11Part2(in [][]string, finished chan bool) {
	var out int

	seats := in
	done := false

	viz := make([][]int, len(seats))
	for i, row := range seats {
		viz[i] = make([]int, len(row))
		for j := range row {
			viz[i][j] = countVisible(seats, i, j)
		}
	}

	// fmt.Println(seats)
	// seats, done = applyRules2(seats)
	// fmt.Println(seats)
	// fmt.Println(viz)

	for !done {
		seats, done = applyRules2(seats)
	}

	for _, row := range seats {
		for _, seat := range row {
			if seat == "#" {
				out++
			}
		}
	}

	fmt.Printf("Part 2:\n\t%d\n", out)
	finished <- true
}

func applyRules2(seats [][]string) ([][]string, bool) {
	out := make([][]string, len(seats))
	for i := range out {
		out[i] = make([]string, len(seats[0]))
	}
	done := true

	for i, row := range seats {
		for j, space := range row {
			switch space {
			case ".":
				out[i][j] = "."
			case "L":
				out[i][j] = "L"
				if countVisible(seats, i, j) == 0 {
					out[i][j] = "#"
					done = false
				}
			case "#":
				out[i][j] = "#"
				if countVisible(seats, i, j) >= 5 {
					out[i][j] = "L"
					done = false
				}
			}
		}
	}
	return out, done
}

func countVisible(seats [][]string, i, j int) int {
	var i2, j2, count int

	i2, j2 = i-1, j
	for i2 >= 0 {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2--
	}

	i2, j2 = i-1, j-1
	for (i2 >= 0) && (j2 >= 0) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2--
		j2--
	}

	i2, j2 = i-1, j+1
	for (i2 >= 0) && (j2 < len(seats[i2])) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2--
		j2++
	}

	i2, j2 = i, j-1
	for j2 >= 0 {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		j2--
	}

	i2, j2 = i, j+1
	for j2 < len(seats[i2]) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		j2++
	}

	i2, j2 = i+1, j
	for i2 < len(seats) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2++
	}

	i2, j2 = i+1, j-1
	for (i2 < len(seats)) && (j2 >= 0) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2++
		j2--
	}

	i2, j2 = i+1, j+1
	for (i2 < len(seats)) && (j2 < len(seats[i2])) {
		if seats[i2][j2] == "#" {
			count++
			break
		}
		if seats[i2][j2] == "L" {
			break
		}
		i2++
		j2++
	}

	return count
}
