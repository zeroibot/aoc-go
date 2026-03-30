package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

type Pair = [2]Int2

func Day04() Solution {
	pairs := data04(true)

	count1, count2 := 0, 0
	for _, pair := range pairs {
		// Part 1
		if isSupersetPair(pair) {
			count1 += 1
		}

		// Part 2
		if isOverlappingPair(pair) {
			count2 += 1
		}
	}

	return NewSolution(count1, count2)
}

func data04(full bool) []Pair {
	return list.Map(ReadLines(22, 4, full), func(line string) Pair {
		p := str.CleanSplit(line, ",")
		p1 := ToInt2(p[0], "-")
		p2 := ToInt2(p[1], "-")
		return Pair{p1, p2}
	})
}

func isSupersetRange(r1 Int2, r2 Int2) bool {
	s1, e1 := r1.Tuple()
	s2, e2 := r2.Tuple()
	return s1 <= s2 && e2 <= e1
}

func isSupersetPair(pair Pair) bool {
	r1, r2 := pair[0], pair[1]
	return isSupersetRange(r1, r2) || isSupersetRange(r2, r1)
}

func isOverlappingPair(pair Pair) bool {
	s1, e1 := pair[0].Tuple()
	s2, e2 := pair[1].Tuple()
	return lang.Ternary(s1 < s2, s2 <= e1, s1 <= e2)
}
