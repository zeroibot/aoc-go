package aoc20

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day09() Solution {
	numbers := data09(true)
	limit := len(numbers)

	// Part 1
	window := 25
	target := 0
	for i := window; i < limit; i++ {
		if !hasPairSum(numbers[i], numbers[i-window:i]) {
			target = numbers[i]
			break
		}
	}

	// Part 2
	part2 := 0
mainLoop:
	for i := range limit {
		j := i
		total := numbers[i]
		for total < target {
			j += 1
			total += numbers[j]
			if total == target {
				options := numbers[i : j+1]
				part2 = slices.Min(options) + slices.Max(options)
				break mainLoop
			}
		}
	}

	return NewSolution(target, part2)
}

func data09(full bool) []int {
	return list.Map(ReadLines(20, 9, full), number.ParseInt)
}

func hasPairSum(target int, numbers []int) bool {
	for _, p := range Combinations(numbers, 2) {
		if list.Sum(p) == target {
			return true
		}
	}
	return false
}
