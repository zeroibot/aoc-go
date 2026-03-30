package aoc23

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

func Day08() Solution {
	moves, m := data08(true)
	limit := len(moves)
	start, goal := "AAA", "ZZZ"

	// Part 1
	count1, i := 0, 0
	curr := start
	for curr != goal {
		idx := moves[i]
		curr = m[curr][idx]
		count1 += 1
		i = (i + 1) % limit
	}

	// Part 2
	starts := list.Filter(dict.Keys(m), func(k string) bool {
		return strings.HasSuffix(k, "A")
	})
	counts := make([]int, 0)
	for _, start2 := range starts {
		count, i := 0, 0
		curr := start2
		for !strings.HasSuffix(curr, "Z") {
			idx := moves[i]
			curr = m[curr][idx]
			count += 1
			i = (i + 1) % limit
		}
		counts = append(counts, count)
	}
	count2 := LCM(counts[0], counts[1], counts[2:]...)

	return NewSolution(count1, count2)
}

func data08(full bool) ([]int, map[string]Str2) {
	lines := ReadRawLines(23, 8, full, true)
	T := map[rune]int{'L': 0, 'R': 1}
	moves := list.Translate([]rune(lines[0]), T)
	m := make(map[string]Str2)
	for _, line := range lines[2:] {
		p := str.CleanSplit(line, "=")
		key := p[0]
		tail := strings.Trim(p[1], "()")
		m[key] = ToStr2(tail, ",")
	}
	return moves, m
}
