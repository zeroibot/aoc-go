package aoc18

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day01() Solution {
	numbers := data01(true)

	// Part 1
	total := list.Sum(numbers)

	// Part 2
	limit := len(numbers)
	done := ds.NewSet[int]()
	i, curr := 0, 0
	for {
		curr += numbers[i]
		if done.Has(curr) {
			break
		}
		done.Add(curr)
		i = (i + 1) % limit
	}

	return NewSolution(total, curr)
}

func data01(full bool) []int {
	return list.Map(ReadLines(18, 1, full), number.ParseInt)
}
