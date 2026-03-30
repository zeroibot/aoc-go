package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day03() Solution {
	words := data03(true)

	// Part 1
	total1 := 0
	for _, word := range words {
		total1 += commonScore(word)
	}

	// Part 2
	total2, limit := 0, len(words)
	for i := 0; i < limit; i += 3 {
		total2 += badgeScore(words[i : i+3])
	}

	return NewSolution(total1, total2)
}

func data03(full bool) []string {
	return ReadLines(22, 3, full)
}

func score(char rune) int {
	v := int(char)
	if 97 <= v && v <= 122 {
		return v - 96
	} else if 65 <= v && v <= 90 {
		return v - 38
	}
	return 0
}

func commonScore(word string) int {
	mid := len(word) / 2
	chars := ds.NewSet[rune]()
	for i, char := range word {
		if i < mid {
			chars.Add(char)
		} else if chars.Has(char) {
			return score(char)
		}
	}
	return 0
}

func badgeScore(words []string) int {
	common := ds.SetFrom([]rune(words[0]))
	for i := 1; i < len(words); i++ {
		uncommon := ds.SetFrom(common.Items())
		for _, char := range words[i] {
			if uncommon.Has(char) {
				uncommon.Delete(char)
			}
		}
		for _, char := range uncommon.Items() {
			common.Delete(char)
		}
	}
	badge := common.Items()[0]
	return score(badge)
}
