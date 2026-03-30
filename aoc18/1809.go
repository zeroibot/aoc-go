package aoc18

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day09() Solution {
	numPlayers, lastMarble := data09(true)

	// Part 1
	score1 := maxScore(numPlayers, lastMarble)

	// Part 2
	score2 := maxScore(numPlayers, lastMarble*100)

	return NewSolution(score1, score2)
}

func data09(full bool) (int, int) {
	line := ReadFirstLine(18, 9, full)
	p := str.SpaceSplit(line)
	numPlayers := number.ParseInt(p[0])
	lastMarble := number.ParseInt(Last(p, 2))
	return numPlayers, lastMarble
}

func maxScore(numPlayers, lastMarble int) int {
	score := make([]int, numPlayers)
	player := 0
	curr := ds.NewDLLNode(0)

	for m := 1; m <= lastMarble; m++ {
		if m%23 == 0 {
			prev7 := curr
			for range 7 {
				prev7 = prev7.Prev
			}
			score[player] += m + prev7.Value
			curr = prev7.Next
			prev7.Remove()
		} else {
			next1 := curr.Next
			curr = next1.AddAfter(m)
		}
		player = (player + 1) % numPlayers
	}
	return slices.Max(score)
}
