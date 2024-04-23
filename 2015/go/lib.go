package aoc

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Solution func(io.Reader) string

func Setup() (io.ReadCloser, int) {
	part := flag.Int("part", 1, "advent of code day 'part' (1 or 2)")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("missing input file")
	}

	file, err := os.Open(args[0])
	Check(err)

	return file, *part
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(part int, one, two Solution, input io.Reader) {
	output := func() string {
		if part == 1 {
			return one(input)
		} else {
			return two(input)
		}
	}()

	fmt.Println(output)
}
