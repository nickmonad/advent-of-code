package main

import (
	"fmt"
	"io"

	aoc "github.com/nickmonad/advent-of-code"
)

// https://adventofcode.com/2015/day/1
func main() {
	input, part, _ := aoc.Setup()
	defer input.Close()

	aoc.Run(Part1, Part2, input, part, false)
}

func Part1(r io.Reader, _ bool) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	floor := 0
	for _, b := range input {
		switch rune(b) {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
		}
	}

	return fmt.Sprintf("%d", floor)
}

func Part2(r io.Reader, _ bool) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	position := func() int {
		floor := 0
		for i, b := range input {
			switch rune(b) {
			case '(':
				floor += 1
			case ')':
				floor -= 1
			default:
			}

			if floor < 0 {
				return i + 1
			}
		}

		return 0
	}()

	return fmt.Sprintf("%d", position)
}
