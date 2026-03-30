package aoc15

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day05() Solution {
	words := data05(true)
	count1, count2 := 0, 0
	for _, word := range words {
		// Part 1
		if isNice(word) {
			count1 += 1
		}

		// Part 2
		if isNice2(word) {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data05(full bool) []string {
	return ReadLines(15, 5, full)
}

var (
	invalids = []string{"ab", "cd", "pq", "xy"}
	vowels   = []rune("aeiou")
)

func isNice(word string) bool {
	hasInvalid := list.Any(invalids, func(invalid string) bool {
		return strings.Contains(word, invalid)
	})
	if hasInvalid {
		return false
	}

	if !HasTwins(word, 0) {
		return false
	}

	freq := CharFreq(word, nil)
	numVowels := list.Sum(list.Map(vowels, func(vowel rune) int {
		return freq[vowel]
	}))
	return numVowels >= 3
}

func isNice2(word string) bool {
	if !HasTwins(word, 1) {
		return false
	}

	pairs := substringGroups(word, 2)
	for _, idxs := range pairs {
		if len(idxs) >= 3 {
			return true
		} else if len(idxs) == 2 && number.Abs(idxs[0]-idxs[1]) >= 2 {
			return true
		}
	}
	return false
}

func substringGroups(word string, length int) [][]int {
	at := make(map[string][]int)
	limit := len(word) - (length - 1)
	for i := range limit {
		sub := word[i : i+length]
		at[sub] = append(at[sub], i)
	}
	return dict.Values(at)
}
