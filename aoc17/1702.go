package aoc17

import (
	"slices"
	"sort"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day02() Solution {
	numbersList := data02(true)
	total1, total2 := 0, 0
	for _, numbers := range numbersList {
		// Part 1
		total1 += slices.Max(numbers) - slices.Min(numbers)

		// Part 2
		for _, pair := range Combinations(numbers, 2) {
			sort.Ints(pair)
			a, b := pair[0], pair[1]
			if b%a == 0 {
				total2 += b / a
				break
			}
		}
	}
	return NewSolution(total1, total2)
}

func data02(full bool) [][]int {
	return list.Map(ReadLines(17, 2, full), func(line string) []int {
		return ToIntList(line, " ")
	})
}
