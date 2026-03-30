package aoc25

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
)

func Day07() Solution {
	grid := data07(true)
	start := list.Filter(list.NumRange(0, len(grid[0])), func(col int) bool {
		return grid[0][col] == 'S'
	})[0]
	curr := make(dict.IntCounter)
	curr[start] = 1
	numSplit := 0
	for row := 1; row < len(grid); row++ {
		next := make(dict.IntCounter)
		for col, count := range curr {
			if grid[row][col] == '^' {
				numSplit += 1
				next[col-1] += count
				next[col+1] += count
			} else {
				next[col] += count
			}
		}
		curr = next
	}
	numPath := list.Sum(dict.Values(curr))
	return NewSolution(numSplit, numPath)
}

func data07(full bool) []string {
	return ReadLines(25, 7, full)
}
