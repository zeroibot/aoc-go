package aoc16

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day03() Solution {
	triples := data03(true)

	// Part 1
	count1 := countValidTriples(triples)

	// Part 2
	triples = readVertical(triples)
	count2 := countValidTriples(triples)

	return NewSolution(count1, count2)
}

func data03(full bool) []Dims3 {
	return list.Map(ReadLines(16, 3, full), func(line string) Dims3 {
		return ToDims3(line, " ")
	})
}

func readVertical(t []Dims3) []Dims3 {
	t2 := make([]Dims3, 0)
	for r := 0; r < len(t); r += 3 {
		for c := range 3 {
			t2 = append(t2, Dims3{
				t[r][c],
				t[r+1][c],
				t[r+2][c],
			})
		}
	}
	return t2
}

func countValidTriples(triples []Dims3) int {
	count := 0
	for _, triple := range triples {
		if isValidTriple(triple) {
			count += 1
		}
	}
	return count
}

func isValidTriple(t Dims3) bool {
	a, b, c := t.Tuple()
	return (a+b > c) && (b+c > a) && (a+c > b)
}
