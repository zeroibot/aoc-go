package aoc15

import (
	"fmt"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
)

func Day04() Solution {
	key := data04(true)

	// Part 1
	idx1 := findHash(key, 5)

	// Part 2
	idx2 := findHash(key, 6)

	return NewSolution(idx1, idx2)
}

func data04(full bool) string {
	return ReadFirstLine(15, 4, full)
}

func findHash(key string, numZeros int) int {
	goal := strings.Repeat("0", numZeros)
	i := 1
	for {
		word := fmt.Sprintf("%s%d", key, i)
		hash := MD5Hash(word)
		if strings.HasPrefix(hash, goal) {
			return i
		}
		i += 1
	}
}
