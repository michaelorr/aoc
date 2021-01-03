package days

import (
	"fmt"

	"orr.co/adventofcode/data"
)

const TREE = "#"
const OPEN = "."

func Day3() {
	Day3Part1()
}

func Day3Part1() {
	defaultMap := buildDefaultMap(data.Day(3))
	a := walkMap(defaultMap, 1, 1)
	b := walkMap(defaultMap, 3, 1)
	c := walkMap(defaultMap, 5, 1)
	d := walkMap(defaultMap, 7, 1)
	e := walkMap(defaultMap, 1, 2)
	fmt.Println(a * b * c * d * e)

}

func walkMap(forest [][]string, xStep int, yStep int) int {
	var trees, x, y int
	for true {
		if y >= len(forest) {
			break
		}
		if x >= len(forest[y]) {
			forest = extendMap(forest)
		}

		if forest[y][x] == TREE {
			trees++
		}
		x += xStep
		y += yStep
	}
	return trees
}

func buildDefaultMap(data <-chan string) [][]string {
	var result [][]string
	x := 0
	for line := range data {
		if len(result) < x+1 {
			result = append(result, []string{})
		}
		for _, by := range line {
			result[x] = append(result[x], string(by))
		}
		x++
	}
	return result
}

func extendMap(data [][]string) [][]string {
	for i, _ := range data {
		data[i] = append(data[i], data[i]...)
	}
	return data
}

func Print(data [][]string) {
	for _, row := range data {
		for y, _ := range row {
			fmt.Print(row[y])
		}
		fmt.Println("")
	}
}
