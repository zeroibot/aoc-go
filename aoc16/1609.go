package aoc16

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day09() Solution {
	word := data09(true)

	// Part 1
	size1 := decompressLength(word, true)

	// Part 2
	size2 := decompressLength(word, false)

	return NewSolution(size1, size2)
}

func data09(full bool) string {
	return ReadFirstLine(16, 9, full)
}

func decompressLength(word string, skip bool) int {
	wordLen := len(word)
	count := make([]int, wordLen)
	for i := range wordLen {
		count[i] = 1
	}
	i := 0
	for i < wordLen {
		if word[i] != '(' {
			i += 1
			continue
		}
		end := strings.Index(word[i:], ")") + i
		p := ToInt2(word[i+1:end], "x")
		size, repeat := p.Tuple()
		start := end + 1
		for j := i; j < start; j++ {
			count[j] = 0
		}
		for j := start; j < start+size; j++ {
			count[j] *= repeat
		}
		if skip {
			i = start + size
		} else {
			i = end + 1
		}
	}
	return list.Sum(count)
}
