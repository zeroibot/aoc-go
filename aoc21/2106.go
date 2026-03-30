package aoc21

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
)

func Day06() Solution {
	fish := data06(true)

	// Part 1
	count1 := simulateFish(fish, 80)

	// Part 2
	count2 := simulateFish(fish, 256)

	return NewSolution(count1, count2)
}

func data06(full bool) []int {
	line := ReadFirstLine(21, 6, full)
	return ToIntList(line, ",")
}

func simulateFish(fish []int, days int) int {
	groups := make(map[int]int)
	for _, f := range fish {
		groups[f] += 1
	}

	for range days {
		groups2 := make(map[int]int)
		for timer, count := range groups {
			if timer == 0 {
				timer = 6
				groups2[8] = count
			} else {
				timer -= 1
			}
			groups2[timer] += count
		}
		groups = groups2
	}
	return list.Sum(dict.Values(groups))
}
