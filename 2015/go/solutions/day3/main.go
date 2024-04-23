package main

import (
	"fmt"
	"io"

	aoc "github.com/nickmonad/advent-of-code"
)

// https://adventofcode.com/2015/day/3
func main() {
	input, part := aoc.Setup()
	defer input.Close()

	aoc.Run(part, Part1, Part2, input)
}

type location struct {
	x int
	y int
}

func Part1(r io.Reader) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	world := make(map[location]bool)
	start := location{x: 0, y: 0}

	world[start] = true
	loc := start

	for _, b := range input {
		switch rune(b) {
		case '^':
			loc.y += 1
		case 'v':
			loc.y -= 1
		case '>':
			loc.x += 1
		case '<':
			loc.x -= 1
		}

		world[loc] = true
	}

	answer := len(world)
	return fmt.Sprintf("%d", answer)
}

func Part2(r io.Reader) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	world := make(map[location]bool)
	start := location{x: 0, y: 0}

	world[start] = true

	santa := start
	robot := start

	isSanta := true

	for _, b := range input {
		var loc *location
		if isSanta {
			loc = &santa
		} else {
			loc = &robot
		}

		switch rune(b) {
		case '^':
			loc.y += 1
		case 'v':
			loc.y -= 1
		case '>':
			loc.x += 1
		case '<':
			loc.x -= 1
		}

		world[*loc] = true
		isSanta = !isSanta
	}

	answer := len(world)
	return fmt.Sprintf("%d", answer)
}
