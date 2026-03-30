package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day02() Solution {
	numbersList := data02(true)

	count1, count2 := 0, 0
	for _, numbers := range numbersList {
		// Part 1
		if isSafe(numbers) {
			count1 += 1
		}

		// Part 2
		if isSafe2(numbers) {
			count2 += 1
		}
	}

	return NewSolution(count1, count2)
}

func data02(full bool) [][]int {
	return list.Map(ReadLines(24, 2, full), func(line string) []int {
		return ToIntList(line, " ")
	})
}

func isSafe(numbers []int) bool {
	diffs := list.Map(NumRange(1, len(numbers)), func(i int) int {
		return numbers[i] - numbers[i-1]
	})
	safeInc := list.All(diffs, func(d int) bool {
		return 1 <= d && d <= 3
	})
	safeDec := list.All(diffs, func(d int) bool {
		return -3 <= d && d <= -1
	})
	return safeInc || safeDec
}

func isSafe2(numbers []int) bool {
	if isSafe(numbers) {
		return true
	}
	for idx := range len(numbers) {
		numbers2 := make([]int, 0)
		numbers2 = append(numbers2, numbers[:idx]...)
		numbers2 = append(numbers2, numbers[idx+1:]...)
		if isSafe(numbers2) {
			return true
		}
	}
	return false
}
