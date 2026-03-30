package aoc24

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"

	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
)

type Rules = map[int]*ds.Set[int]

func Day05() Solution {
	rules, pages := data05(true)

	total1, total2 := 0, 0
	for _, page := range pages {
		// Part 1
		if isValidPage(page, rules) {
			total1 += page[len(page)/2]
		}

		// Part 2
		total2 += correctOrderMid(page, rules)
	}

	return NewSolution(total1, total2)
}

func data05(full bool) (Rules, [][]int) {
	rules := make([]Int2, 0)
	pages := make([][]int, 0)
	part2 := false
	for _, line := range ReadRawLines(24, 5, full, true) {
		if part2 {
			pages = append(pages, ToIntList(line, ","))
		} else if line == "" {
			part2 = true
		} else {
			rules = append(rules, ToInt2(line, "|"))
		}
	}
	book := make(Rules)
	for _, rule := range rules {
		before, after := rule.Tuple()
		if _, ok := book[after]; !ok {
			book[after] = ds.NewSet[int]()
		}
		book[after].Add(before)
	}
	return book, pages
}

func isValidPage(page []int, rules Rules) bool {
	for i := range len(page) - 1 {
		after := ds.SetFrom(page[i+1:])
		blacklist := rules[page[i]]
		common := after.Intersection(blacklist)
		if common.Len() > 0 {
			return false
		}
	}
	return true
}

func correctOrderMid(page []int, rules Rules) int {
	valid := true
	idx, limit := 0, len(page)-1
	for idx < limit {
		curr := page[idx]
		after := ds.SetFrom(page[idx+1:])
		blacklist := rules[curr]
		common := after.Intersection(blacklist)
		if common.Len() == 0 {
			idx += 1
		} else {
			valid = false
			indexes := list.Map(common.Items(), func(x int) int {
				return slices.Index(page, x)
			})
			insert := slices.Max(indexes) + 1
			page[idx] = 0
			page2 := make([]int, 0)
			page2 = append(page2, page[:insert]...)
			page2 = append(page2, curr)
			page2 = append(page2, page[insert:]...)
			page = list.Filter(page2, func(x int) bool {
				return x != 0
			})
		}
	}
	return lang.Ternary(valid, 0, page[len(page)/2])
}
