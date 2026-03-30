package aoc16

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day01() Solution {
	moves := data01(true)

	// Part 1
	hq := findHQ(moves, false)
	dist1 := ManhattanOrigin(hq)

	// Part 2
	hq = findHQ(moves, true)
	dist2 := ManhattanOrigin(hq)

	return NewSolution(dist1, dist2)
}

const (
	left  int = -1
	right int = 1
)

func data01(full bool) []Int2 {
	T := map[byte]int{
		'L': left,
		'R': right,
	}
	line := ReadFirstLine(16, 1, full)
	return list.Map(str.CleanSplit(line, ","), func(move string) Int2 {
		turn := T[move[0]]
		steps := number.ParseInt(move[1:])
		return Int2{turn, steps}
	})
}

func findHQ(moves []Int2, atVisitedTwice bool) Coords {
	curr := Coords{0, 0}
	d := X
	visited := ds.NewSet[Coords]()
	for _, move := range moves {
		turn, steps := move.Tuple()
		if d == X {
			d = lang.Ternary(turn == left, L, R)
		} else if turn == left {
			d = LeftOf[d]
		} else if turn == right {
			d = RightOf[d]
		}
		for range steps {
			curr = Move(curr, d)
			if atVisitedTwice && visited.Has(curr) {
				return curr
			}
			visited.Add(curr)
		}
	}
	return curr
}
