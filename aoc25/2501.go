package aoc25

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/number"
)

func Day01() Solution {
	moves := data01(true)

	count1, count2 := 0, 0
	curr := 50
	for _, r := range moves {
		sign, repeat := r[0], r[1]
		for range repeat {
			curr = (curr + sign) % 100
			if curr == 0 {
				count2 += 1
			}
		}
		if curr == 0 {
			count1 += 1
		}
	}

	return NewSolution(count1, count2)
}

func data01(full bool) [][2]int {
	moves := make([][2]int, 0)
	for _, line := range ReadLines(25, 1, full) {
		sign := lang.Ternary(line[0] == 'R', 1, -1)
		moves = append(moves, [2]int{sign, number.ParseInt(line[1:])})
	}
	return moves
}
