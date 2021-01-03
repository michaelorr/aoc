package days

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"orr.co/adventofcode/data"
)

func Day7() {
	c1 := make(chan string)
	c2 := make(chan string)
	finished := make(chan bool)
	go Day7Part1(c1, finished)
	go Day7Part2(c2, finished)

	for line := range data.Day(7) {
		c1 <- line
		c2 <- line
	}

	close(c1)
	close(c2)
	<-finished
	<-finished
}

func Day7Part1(lines chan string, finished chan bool) {
	mapping := make(map[string][]string)
	bagParser := regexp.MustCompile("(.*) bags contain")
	contentsParser := regexp.MustCompile("([[:digit:]] ([a-z]+ [a-z]+))+")
	for line := range lines {
		bag := bagParser.FindStringSubmatch(line)[1]
		contents := contentsParser.FindAllStringSubmatch(line, -1)
		for _, j := range contents {
			mapping[j[2]] = append(mapping[j[2]], bag)
		}
	}
	// fmt.Printf("Part 1:\n\t%#v\n", uniqueCount(parents("shiny gold", mapping)))
	finished <- true
}

func uniqueCount(bags []string) int {
	res := make(map[string]bool)
	for _, v := range bags {
		res[v] = true
	}
	return len(res)
}

func parents(bag string, mapping map[string][]string) []string {
	res := mapping[bag]
	for _, v := range mapping[bag] {
		res = append(res, parents(v, mapping)...)
	}
	return res
}

func Day7Part2(lines chan string, finished chan bool) {
	mapping := make(map[string]map[string]int)

	bagParser := regexp.MustCompile("(.*) bags contain")
	contentsParser := regexp.MustCompile("(([[:digit:]]) ([a-z]+ [a-z]+))+")

	for line := range lines {
		parent := bagParser.FindStringSubmatch(line)[1]
		contents := contentsParser.FindAllStringSubmatch(line, -1)

		for _, j := range contents {
			if _, ok := mapping[parent]; !ok {
				mapping[parent] = make(map[string]int)
			}

			color := j[3]
			count, _ := strconv.Atoi(j[2])
			mapping[parent][color] = count
		}
	}
	b, _ := json.MarshalIndent(mapping, "", "    ")
	fmt.Printf("%s\n", string(b))
	fmt.Printf("Part 2:\n\t%#v\n", children("shiny gold", mapping)-1)
	finished <- true
}

func children(bag string, mapping map[string]map[string]int) int {
	res := 1
	for child, count := range mapping[bag] {
		for i := 0; i < count; i++ {
			childCount := children(child, mapping)
			res += childCount
			fmt.Printf("Child %s Adding %d\n", child, childCount)
		}
	}
	fmt.Printf("Bag %s has %d children\n", bag, res)
	return res
}
