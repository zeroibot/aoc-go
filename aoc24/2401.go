package aoc24

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/number"
)

func Day01() Solution {
	col1, col2 := data01(true)

	// Part 1
	slices.Sort(col1)
	slices.Sort(col2)
	total1 := 0
	for i := range len(col1) {
		total1 += number.Abs(col1[i] - col2[i])
	}

	// Part 2
	freq := CountFreq(col2)
	total2 := 0
	for _, x := range col1 {
		total2 += x * freq[x]
	}

	return NewSolution(total1, total2)
}

func data01(full bool) ([]int, []int) {
	col1 := make([]int, 0)
	col2 := make([]int, 0)
	for _, line := range ReadLines(24, 1, full) {
		a, b := ToInt2(line, " ").Tuple()
		col1 = append(col1, a)
		col2 = append(col2, b)
	}
	return col1, col2
}
