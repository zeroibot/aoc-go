package aoc16

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
)

func Day06() Solution {
	words := data06(true)
	numCols := len(words[0])
	freq := columnFrequency(words, numCols)
	maxMsg := make([]rune, numCols)
	minMsg := make([]rune, numCols)
	for col := range numCols {
		colFreq := make([]CharInt, 0)
		for k, v := range freq[col] {
			colFreq = append(colFreq, CharInt{Char: k, Int: v})
		}
		slices.SortFunc(colFreq, SortCharIntAsc)

		// Part 1
		minPair := colFreq[0]
		minMsg[col] = minPair.Char

		// Part 2
		maxPair := Last(colFreq, 1)
		maxMsg[col] = maxPair.Char
	}
	return NewSolution(string(maxMsg), string(minMsg))
}

func data06(full bool) []string {
	return ReadLines(16, 6, full)
}

func columnFrequency(words []string, numCols int) map[int]map[rune]int {
	freq := make(map[int]map[rune]int)
	for i := range numCols {
		freq[i] = make(map[rune]int)
	}
	for _, word := range words {
		for col, char := range word {
			freq[col][char] += 1
		}
	}
	return freq
}
