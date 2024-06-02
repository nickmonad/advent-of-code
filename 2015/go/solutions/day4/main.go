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

	prefix := 5 // "00000"

	if optimized {
		return optimized_f(input, prefix)
	} else {
		return naive(input, prefix)
	}
}

func Part2(r io.Reader, optimized bool) string {
	input, err := io.ReadAll(r)
	aoc.Check(err)

	prefix := 6 // "000000"

	if optimized {
		return optimized_f(input, prefix)
	} else {
		return naive(input, prefix)
	}
}

// "naive" solution tries every nonce in a single thread
func naive(input []byte, prefix int) string {
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
func optimized_f(input []byte, prefix int) string {
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

// Check if the given MD5 hash has the given leading number of zeros.
func check(hash [16]byte, leading int) bool {
	numBytes := leading / 2
	for i := 0; i < numBytes; i++ {
		if hash[i] != 0x00 {
			return false
		}
	}

	if leading%2 != 0 {
		// looking for odd number of zeros...
		// since each hex character represents 4 bits, we have to use a bit shift to check our last byte
		// in the hash if the desired number of leading zeros is odd. For example, if we are looking for
		// `000` (odd number of zeros), the byte sequence _could_ be `0x00 0x01`, or `0x00 0x02`, etc...
		// In this case, we bit shift the last byte to the right by 4, and check if that result is 0,
		// ignoring 4 least significant bits.

		// last index in hash to check is equal to `numBytes` because of the integer division above.
		// for example, if leading == 5, numBytes == 2 (due to integer division), and we need to
		// check the 3rd byte in the sequence (i.e. index 2)
		return (hash[numBytes] >> 4) == 0x00
	}

	return true
}

//func check(hash [16]byte, leading string) bool {
//	hex := fmt.Sprintf("%x", hash)
//	return strings.HasPrefix(hex, leading)
//}
