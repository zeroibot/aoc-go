package aoc17

import (
	"fmt"
	"slices"

	"github.com/zeroibot/aoc-go/aoc"
)

const knotHashLimit int = 256

func knotHash(lengths []int, rounds int) []int {
	lengths = append(lengths, 17, 31, 73, 47, 23)
	numbers := make([]int, knotHashLimit)
	for i := range knotHashLimit {
		numbers[i] = i
	}
	i, skip := 0, 0
	for range rounds {
		for _, length := range lengths {
			if length > knotHashLimit {
				continue
			}

			j := i + length
			if j <= knotHashLimit {
				slices.Reverse(numbers[i:j])
			} else {
				s := knotHashLimit - i
				j = length - s
				chunk := aoc.JoinLists(numbers[i:], numbers[:j])
				slices.Reverse(chunk)
				for x := range s {
					numbers[i+x] = chunk[x]
				}
				for x := range j {
					numbers[x] = chunk[s+x]
				}
			}
			i = (i + length + skip) % knotHashLimit
			skip += 1
		}
	}
	return numbers
}

func hexCode(x int) string {
	return fmt.Sprintf("%0.2x", x)
}
