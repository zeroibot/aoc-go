package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day06() Solution {
	line := data06(true)

	// Part 1
	marker1 := findMarker(line, 4)

	// Part 2
	marker2 := findMarker(line, 14)

	return NewSolution(marker1, marker2)
}

func data06(full bool) string {
	return ReadFirstLine(22, 6, full)
}

func findMarker(line string, length int) int {
	for n := length; n <= len(line); n++ {
		if allUnique(line[n-length : n]) {
			return n
		}
	}
	return -1
}

func allUnique(text string) bool {
	unique := ds.SetFrom([]rune(text))
	return len(text) == unique.Len()
}
