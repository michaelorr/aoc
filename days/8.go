package days

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"orr.co/adventofcode/data"
)

func Day8() {
	var ops1 []Operation
	var ops2 []Operation
	finished := make(chan bool)

	for line := range data.Day(8) {
		ops1 = append(ops1, parseOperation(line))
		ops2 = append(ops2, parseOperation(line))
	}
	go Day8Part1(ops1, finished)
	go Day8Part2(ops2, finished)

	<-finished
	<-finished
}

const NOP = "nop"
const ACC = "acc"
const JMP = "jmp"

type Operation struct {
	Instruction string
	Argument    int
	Executed    bool
}

func parseOperation(in string) Operation {
	r := regexp.MustCompile("([a-z]{3}) ([+-][[:digit:]]+)")
	parts := r.FindStringSubmatch(in)
	i, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(i)
	}

	return Operation{Instruction: parts[1], Argument: i}
}

func Day8Part1(ops []Operation, finished chan bool) {
	ptr := 0
	acc := 0

	for !ops[ptr].Executed {
		op := &ops[ptr]
		if op.Executed {
			break
		}

		switch op.Instruction {
		case ACC:
			acc += op.Argument
		case JMP:
			ptr += op.Argument - 1
		}
		// fmt.Printf("Executing: %#v\nACC: %d\n", op, acc)
		ptr++
		op.Executed = true
	}
	fmt.Printf("Part 1:\n\t%d\n", acc)
	finished <- true
}

func (o *Operation) String() string {
	var ex string
	if o.Executed {
		ex = " x"
	}
	return fmt.Sprintf("%s %d%s", o.Instruction, o.Argument, ex)
}

func Day8Part2(ops []Operation, finished chan bool) {
	var nops []int
	var jmps []int
	for i, op := range ops {
		if op.Instruction == NOP {
			nops = append(nops, i)
		} else if op.Instruction == JMP {
			jmps = append(jmps, i)
		}
	}

	var res int
	// fmt.Println("Original:")
	// fmt.Println(ops)
	for _, i := range nops {
		ops[i].Instruction = JMP
		// fmt.Printf("NOP->JMP[%d]\n", i)
		res = findHalt(ops)
		for i, _ := range ops {
			ops[i].Executed = false
		}
		ops[i].Instruction = NOP
		if res != 0 {
			fmt.Printf("Part 2:\n\t%d\n", res)
			finished <- true
			return
		}
	}

	for _, i := range jmps {
		ops[i].Instruction = NOP
		// fmt.Printf("JMP->NOP[%d]\n", i)
		res = findHalt(ops)
		for i, _ := range ops {
			ops[i].Executed = false
		}
		ops[i].Instruction = JMP

		if res != 0 {
			fmt.Printf("Part 2:\n\t%d\n", res)
			finished <- true
			return
		}
	}

	finished <- true

}

func findHalt(ops []Operation) int {
	ptr := 0
	acc := 0

	op := &ops[ptr]
	for !op.Executed {
		switch op.Instruction {
		case ACC:
			acc += op.Argument
		case JMP:
			ptr += op.Argument - 1
		}
		ptr++
		op.Executed = true
		if ptr >= len(ops) {
			return acc
		}
		op = &ops[ptr]
	}
	return 0
}
