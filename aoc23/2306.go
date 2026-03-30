package aoc23

import (
	"strconv"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day06() Solution {
	times, bests := data06(true)

	// Part 1
	total1 := 1
	for i, limit := range times {
		breakers := list.Filter(computeOutcomes(limit), func(d int) bool {
			return d > bests[i]
		})
		total1 *= len(breakers)
	}

	// Part 2
	newLimit := strings.Join(list.Map(times, strconv.Itoa), "")
	newBest := strings.Join(list.Map(bests, strconv.Itoa), "")
	limit, best := number.ParseInt(newLimit), number.ParseInt(newBest)
	breakers := list.Filter(computeOutcomes(limit), func(d int) bool {
		return d > best
	})
	total2 := len(breakers)

	return NewSolution(total1, total2)
}

func data06(full bool) ([]int, []int) {
	lines := ReadLines(23, 6, full)
	times := list.Map(str.SpaceSplit(lines[0])[1:], number.ParseInt)
	bests := list.Map(str.SpaceSplit(lines[1])[1:], number.ParseInt)
	return times, bests
}

func computeOutcomes(limit int) []int {
	return list.Map(NumRange(0, limit+1), func(hold int) int {
		return hold * (limit - hold)
	})
}
