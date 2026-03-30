package aoc20

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

type Group = []string

func Day06() Solution {
	groups := data06(true)
	total1, total2 := 0, 0
	for _, group := range groups {
		// Part 1
		total1 += countYes(group)

		// Part 2
		total2 += countAllYes(group)
	}

	return NewSolution(total1, total2)
}

func data06(full bool) []Group {
	groups := make([]Group, 0)
	curr := make(Group, 0)
	for _, line := range ReadRawLines(20, 6, full, true) {
		if line == "" {
			groups = append(groups, curr)
			curr = make(Group, 0)
		} else {
			curr = append(curr, line)
		}
	}
	groups = append(groups, curr)
	return groups
}

func countYes(group Group) int {
	qs := ds.NewSet[rune]()
	for _, questions := range group {
		for _, q := range questions {
			qs.Add(q)
		}
	}
	return qs.Len()
}

func countAllYes(group Group) int {
	count := make(map[rune]int)
	for _, questions := range group {
		for _, q := range questions {
			count[q] += 1
		}
	}
	groupSize := len(group)
	allYes := list.Filter(dict.Keys(count), func(k rune) bool {
		return count[k] == groupSize
	})
	return len(allYes)
}
