package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

var (
	rock     int = 1
	paper    int = 2
	scissors int = 3
	lose     int = 0
	draw     int = 3
	win      int = 6
)

func Day02() Solution {
	// Part 1
	T1 := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	games := data02(T1, true)
	total1 := 0
	for _, game := range games {
		total1 += computeScore(game)
	}

	// Part 2
	T2 := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": lose,
		"Y": draw,
		"Z": win,
	}
	games = data02(T2, true)
	total2 := 0
	for _, game := range games {
		total2 += coerceScore(game)
	}

	return NewSolution(total1, total2)
}

func data02(T map[string]int, full bool) []Int2 {
	return list.Map(ReadLines(22, 2, full), func(line string) Int2 {
		p := str.CleanSplit(line, " ")
		return Int2{T[p[0]], T[p[1]]}
	})
}

var (
	winsOver = map[int]int{rock: scissors, paper: rock, scissors: paper}
	losesTo  = map[int]int{scissors: rock, rock: paper, paper: scissors}
)

func computeScore(game Int2) int {
	opp, you := game.Tuple()
	score := you
	if opp == you {
		score += draw
	} else if winsOver[you] == opp {
		score += win
	}
	return score
}

func coerceScore(game Int2) int {
	opp, result := game.Tuple()
	you := 0
	switch result {
	case win:
		you = losesTo[opp]
	case lose:
		you = winsOver[opp]
	case draw:
		you = opp
	}
	return computeScore(Int2{opp, you})
}
