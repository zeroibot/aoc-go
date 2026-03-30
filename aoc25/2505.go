package aoc25

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day05() Solution {
	ranges, ingredients := data05(true)
	ranges = MergeRanges(ranges)

	// Part 1
	boolToInt := map[bool]int{true: 1, false: 0}
	count1 := list.Sum(list.Map(ingredients, func(i int) int {
		inRange := list.Any(ranges, func(r Int2) bool {
			return r[0] <= i && i <= r[1]
		})
		return boolToInt[inRange]
	}))

	// Part 2
	count2 := list.Sum(list.Map(ranges, func(r Int2) int {
		return r[1] - r[0] + 1
	}))

	return NewSolution(count1, count2)
}

func data05(full bool) ([]Int2, []int) {
	ranges := make([]Int2, 0)
	ingredients := make([]int, 0)
	inPart2 := false
	for _, line := range ReadRawLines(25, 5, full, true) {
		if line == "" {
			inPart2 = true
		} else if inPart2 {
			ingredients = append(ingredients, number.ParseInt(line))
		} else {
			ranges = append(ranges, ToInt2(line, "-"))
		}
	}
	return ranges, ingredients

}
