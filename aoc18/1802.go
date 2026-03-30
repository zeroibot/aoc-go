package aoc18

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
)

func Day02() Solution {
	words := data02(true)

	// Part 1
	count2, count3 := 0, 0
	for _, word := range words {
		freq := dict.Values(CharFreq(word, nil))
		if slices.Contains(freq, 2) {
			count2 += 1
		}
		if slices.Contains(freq, 3) {
			count3 += 1
		}
	}
	part1 := count2 * count3

	// Part 2
	var word string
	for _, pair := range Combinations(words, 2) {
		word1, word2 := pair[0], pair[1]
		diff := strDiff(word1, word2)
		if len(diff) == 1 {
			idx := diff[0]
			word = word1[:idx] + word1[idx+1:]
			break
		}
	}

	return NewSolution(part1, word)
}

func data02(full bool) []string {
	return ReadLines(18, 2, full)
}

func strDiff(word1, word2 string) []int {
	diff := make([]int, 0)
	for i := range len(word1) {
		if word1[i] != word2[i] {
			diff = append(diff, i)
		}
	}
	return diff
}
