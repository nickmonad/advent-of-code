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
	input, part := aoc.Setup()
	defer input.Close()

	aoc.Run(part, Part1, Part2, input)
}

func Part1(r io.Reader) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	return naive(input, "00000")
}

func Part2(r io.Reader) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	// return optimized(input, "000000")
	return naive(input, "000000")
}

// "naive" solution tries every nonce in a single thread
func naive(input []byte, prefix string) string {
	seed := strings.TrimSpace(string(input))
	nonce := 1
	for {
		preimage := []byte(fmt.Sprintf("%s%d", seed, nonce))
		hash := md5.Sum(preimage)

		hex := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hex, prefix) {
			break // found nonce
		}

		nonce += 1
	}

	return fmt.Sprintf("%d", nonce)
}

// "optimized" solution uses multiple go-routines to converge on a nonce
// NOTE: This actually makes it run slower, lol. Most likely due to the extra blocking and context switching.
// Even utilizing all available CPU cores requires extra coordination and the "warm up" is not likely worth the
// effort for this heavy CPU bound task.
func optimized(input []byte, prefix string) string {
	seed := strings.TrimSpace(string(input))
	numWorkers := runtime.NumCPU()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// NOTE: Even making the nonce channel buffered doesn't help significantly
	nonces := make(chan int, numWorkers)
	solution := make(chan int)

	// "manager" routine spawns workers and hands out nonce values
	wg.Add(1)
	go func() {
		defer wg.Done()

		for n := 0; n < numWorkers; n++ {
			// spawn worker
			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					select {
					case <-ctx.Done():
						return

					case nonce := <-nonces:
						preimage := []byte(fmt.Sprintf("%s%d", seed, nonce))
						hash := md5.Sum(preimage)

						hex := fmt.Sprintf("%x", hash)
						if strings.HasPrefix(hex, prefix) {
							solution <- nonce
							return
						}
					}
				}
			}()
		}

		nonce := 1
		for {
			select {
			case <-ctx.Done():
				return

			default:
				nonces <- nonce
				nonce += 1
			}
		}
	}()

	winner := <-solution
	cancel()

	wg.Wait()
	return fmt.Sprintf("%d", winner)
}
