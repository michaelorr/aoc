package data

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Take lines, then split along comma, return each line as string slice along channel
func AsStrArr(strLines <-chan string) <-chan []string {
	resultLines := make(chan []string)
	go func() {
		defer close(resultLines)
		for str := range strLines {
			resultLines <- strings.Split(str, ",")
		}
	}()
	return resultLines
}

// Range over all lines, split everything on comma, convert each token to int
// and return all data as a slice of ints
func AsIntArr(strLines <-chan string) []int {
	strSlice := strings.Split(<-strLines, ",")
	var intSlice []int
	for _, item := range strSlice {
		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, i)
	}
	return intSlice
}

// Returns lines as ints via channel
func AsInts(strLines <-chan string) <-chan int {
	intLines := make(chan int)
	go func() {
		defer close(intLines)
		for str := range strLines {
			int, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			intLines <- int
		}
	}()
	return intLines
}

// Returns lines as strings via channel
func Day(day int) <-chan string {
	path := fmt.Sprintf("./data/input/%d.txt", day)

	lines := make(chan string)
	go func() {
		defer close(lines)

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return lines
}

// group the chunks of lines using a blank line as separator
func AsChunks(input <-chan string) <-chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		var b bytes.Buffer

		for line := range input {
			b.WriteString(line + " ")
			if line == "" {
				output <- b.String()
				b.Reset()
			}
		}
		output <- b.String()
	}()
	return output
}

// group the chunks of lines using a blank line as separator
func AsChunkedSlices(input <-chan string) <-chan []string {
	output := make(chan []string)
	go func() {
		defer close(output)
		var s []string

		for line := range input {
			if line != "" {
				s = append(s, line)
			} else {
				output <- s
				s = nil
			}
		}
		output <- s
	}()
	return output
}
