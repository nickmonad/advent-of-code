package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"

	aoc "github.com/nickmonad/advent-of-code"
)

// https://adventofcode.com/2015/day/4
func main() {
	input, part, optimized := aoc.Setup()
	defer input.Close()

	aoc.Run(Part1, Part2, input, part, optimized)
}

func Part1(r io.Reader, optimized bool) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	prefix := "00000" // 5 zeros

	if optimized {
		return optimized_f(input, prefix)
	} else {
		return naive(input, prefix)
	}
}

func Part2(r io.Reader, optimized bool) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	prefix := "000000" // 6 zeros

	if optimized {
		return optimized_f(input, prefix)
	} else {
		return naive(input, prefix)
	}
}

// "naive" solution tries every nonce in a single thread
func naive(input []byte, prefix string) string {
	seed := strings.TrimSpace(string(input))
	nonce := 1

	for {
		preimage := []byte(fmt.Sprintf("%s%d", seed, nonce))
		hash := md5.Sum(preimage)

		if check(hash, prefix) {
			break // found nonce
		}

		nonce += 1
	}

	return fmt.Sprintf("%d", nonce)
}

// Optimized solution uses multiple go-routines to converge on a nonce.
// Each worker is given some base from which the next nonce it should try is calculated,
// which ensures no two workers attempt the same nonce.
func optimized_f(input []byte, prefix string) string {
	seed := strings.TrimSpace(string(input))
	numWorkers := runtime.NumCPU()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	solution := make(chan int, 1)

	for b := 0; b < numWorkers; b++ {
		// spawn worker
		wg.Add(1)
		go func(base, n int) {
			defer wg.Done()

			i := 0
			for {
				select {
				case <-ctx.Done():
					return

				default:
					nonce := base + (i * n)
					preimage := []byte(fmt.Sprintf("%s%d", seed, nonce))
					hash := md5.Sum(preimage)

					if check(hash, prefix) {
						solution <- nonce
						return
					}

					i = i + 1
				}
			}
		}(b, numWorkers)
	}

	winner := <-solution
	cancel()

	wg.Wait()
	return fmt.Sprintf("%d", winner)
}

func check(hash [16]byte, prefix string) bool {
	hex := fmt.Sprintf("%x", hash)
	return strings.HasPrefix(hex, prefix)
}
