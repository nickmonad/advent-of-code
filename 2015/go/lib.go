package aoc

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Solution func(io.Reader, bool) string

func Setup() (io.ReadCloser, int, bool) {
	part := flag.Int("part", 1, "advent of code day 'part' (1 or 2)")
	optimized := flag.Bool("optimized", false, "run the optimized version of the solution, if available")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("missing input file")
	}

	file, err := os.Open(args[0])
	Check(err)

	return file, *part, *optimized
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(one, two Solution, input io.Reader, part int, optimized bool) {
	start := time.Now()
	output := func() string {
		if part == 1 {
			return one(input, optimized)
		} else {
			return two(input, optimized)
		}
	}()

	end := time.Now()
	fmt.Println(output)
	fmt.Printf("time: %s\n", end.Sub(start))
}
