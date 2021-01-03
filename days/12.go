package days

import (
	"fmt"
	"regexp"
	"strconv"

	"orr.co/adventofcode/data"
)

type Ship struct {
	Heading int
	X       int
	Y       int
}

type Waypoint struct {
	X int
	Y int
}

type Instruction struct {
	Command string
	Amount  int
}

func Day12() {
	finished := make(chan bool)

	var lines []Instruction
	for line := range data.Day(12) {
		r := regexp.MustCompile("([NSEWLRF])([[:digit:]]*)")
		data := r.FindStringSubmatch(line)
		amt, _ := strconv.Atoi(data[2])
		i := Instruction{Command: data[1], Amount: amt}
		lines = append(lines, i)
	}

	go Day12Part1(lines, finished)
	go Day12Part2(lines, finished)

	<-finished
	<-finished
}

func Day12Part1(in []Instruction, finished chan bool) {
	var out int

	s := Ship{Heading: 90}

	for _, inst := range in {
		s.ApplyInst1(inst)
	}

	a := s.X
	if a < 0 {
		a *= -1
	}

	b := s.Y
	if b < 0 {
		b *= -1
	}
	out = a + b

	fmt.Printf("Part 1:\n\t%d\n", out)
	finished <- true
}

func (s *Ship) ApplyInst1(i Instruction) {
	switch i.Command {
	case "R":
		s.Heading = (s.Heading + i.Amount) % 360
	case "L":
		s.Heading = (s.Heading - i.Amount + 360) % 360

	case "F":
		switch s.Heading {
		case 0:
			s.Y += i.Amount
		case 90:
			s.X += i.Amount
		case 180:
			s.Y -= i.Amount
		case 270:
			s.X -= i.Amount
		}
	case "E":
		s.X += i.Amount
	case "N":
		s.Y += i.Amount
	case "S":
		s.Y -= i.Amount
	case "W":
		s.X -= i.Amount
	}
}

func Day12Part2(in []Instruction, finished chan bool) {
	var out int

	s := Ship{}
	waypoint := Waypoint{X: 10, Y: 1}

	for _, inst := range in {
		ApplyInst2(&s, &waypoint, inst)
	}

	a := s.X
	if a < 0 {
		a *= -1
	}

	b := s.Y
	if b < 0 {
		b *= -1
	}
	out = a + b

	fmt.Printf("Part 2:\n\t%d\n", out)
	finished <- true
}

func ApplyInst2(s *Ship, w *Waypoint, i Instruction) {
	switch i.Command {
	case "R":
		if i.Amount == 90 {
			w.X, w.Y = w.Y, -w.X
		} else if i.Amount == 180 {
			w.X, w.Y = -w.X, -w.Y
		} else if i.Amount == 270 {
			w.X, w.Y = -w.Y, w.X
		}
	case "L":
		if i.Amount == 90 {
			w.X, w.Y = -w.Y, w.X
		} else if i.Amount == 180 {
			w.X, w.Y = -w.X, -w.Y
		} else if i.Amount == 270 {
			w.X, w.Y = w.Y, -w.X
		}
	case "F":
		s.X += i.Amount * w.X
		s.Y += i.Amount * w.Y
	case "E":
		w.X += i.Amount
	case "N":
		w.Y += i.Amount
	case "S":
		w.Y -= i.Amount
	case "W":
		w.X -= i.Amount
	}
}
