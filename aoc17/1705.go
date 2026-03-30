package aoc17

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day05() Solution {
	full := true

	// Part 1
	count1 := countJumps(full, false)

	// Part 2
	count2 := countJumps(full, true)

	return NewSolution(count1, count2)
}

func data05(full bool) []int {
	return list.Map(ReadLines(17, 5, full), number.ParseInt)
}

func countJumps(full bool, clip bool) int {
	jumps := data05(full)
	limit := len(jumps)
	i, count := 0, 0
	for 0 <= i && i < limit {
		jump := jumps[i]
		increment := 1
		if clip && jump >= 3 {
			increment = -1
		}
		jumps[i] += increment
		i += jump
		count += 1
	}
	return count
}
