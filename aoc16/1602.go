package aoc16

import (
	. "github.com/zeroibot/aoc-go/aoc"
)

func Day02() Solution {
	movesList := data02(true)

	// Part 1
	pad1 := Keypad{
		Grid:   []string{"123", "456", "789"},
		Start:  Coords{1, 1},
		Bounds: Dims2{3, 3},
	}
	code1 := solveCode(pad1, movesList)

	// Part 2
	pad2 := Keypad{
		Grid:   []string{"00100", "02340", "56789", "0ABC0", "00D00"},
		Start:  Coords{2, 0},
		Bounds: Dims2{5, 5},
	}
	code2 := solveCode(pad2, movesList)

	return NewSolution(code1, code2)
}

func data02(full bool) []string {
	return ReadLines(16, 2, full)
}

type Keypad struct {
	Grid   []string
	Start  Coords
	Bounds Dims2
}

func solveCode(pad Keypad, movesList []string) string {
	T := map[rune]Delta{
		'U': U,
		'D': D,
		'L': L,
		'R': R,
	}
	code := make([]byte, len(movesList))
	curr := pad.Start
	for i, moves := range movesList {
		for _, m := range moves {
			nxt := Move(curr, T[m])
			r, c := nxt.Tuple()
			if InsideBounds(nxt, pad.Bounds) && pad.Grid[r][c] != '0' {
				curr = nxt
			}
		}
		row, col := curr.Tuple()
		code[i] = pad.Grid[row][col]
	}
	return string(code)
}
