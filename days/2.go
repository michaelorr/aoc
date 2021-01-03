package days

import (
	"fmt"

	"orr.co/adventofcode/data"
	"orr.co/adventofcode/passwords"
)

func Day2() {
	Day2Part2()
}

func Day2Part1() {
	var numValid int
	for line := range data.Day(2) {
		passPolicy := passwords.ParsePolicy(line)
		if passPolicy.OldValid() {
			numValid++
		}
	}
	fmt.Printf("The total number of valid passwords: %d\n", numValid)
}

func Day2Part2() {
	var numValid int
	for line := range data.Day(2) {
		passPolicy := passwords.ParsePolicy(line)
		if passPolicy.NewValid() {
			numValid++
		}
	}
	fmt.Printf("The total number of valid passwords: %d\n", numValid)
}
