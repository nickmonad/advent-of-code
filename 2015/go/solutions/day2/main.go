package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	aoc "github.com/nickmonad/advent-of-code"
)

// https://adventofcode.com/2015/day/2
func main() {
	input, part, _ := aoc.Setup()
	defer input.Close()

	aoc.Run(Part1, Part2, input, part, false)
}

func Part1(r io.Reader, _ bool) string {
	buffer := bufio.NewReader(r)
	total := 0

	for {
		// read all input, line by line
		// updating total as we go along
		line, err := buffer.ReadString('\n')
		if err != nil { // io.EOF, most likely
			break
		}

		// assuming input is in correct format, not checking array bound
		dimensions := strings.Split(strings.TrimSpace(line), "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		s1 := l * w
		s2 := w * h
		s3 := h * l
		area := (2 * s1) + (2 * s2) + (2 * s3)
		smallest := min(s1, s2, s3)

		total = total + area + smallest
	}

	return fmt.Sprintf("%d", total)
}

func Part2(r io.Reader, _ bool) string {
	buffer := bufio.NewReader(r)
	total := 0

	for {
		// read all input, line by line
		// updating total as we go along
		line, err := buffer.ReadString('\n')
		if err != nil { // io.EOF, most likely
			break
		}

		// assuming input is in correct format, not checking array bound
		dimensions := strings.Split(strings.TrimSpace(line), "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		p1 := (2 * l) + (2 * w)
		p2 := (2 * w) + (2 * h)
		p3 := (2 * h) + (2 * l)
		smallest := min(p1, p2, p3)
		volume := l * w * h

		total = total + smallest + volume
	}

	return fmt.Sprintf("%d", total)
}
