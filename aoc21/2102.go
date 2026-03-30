package aoc21

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day02() Solution {
	moves := data02(true)

	// Part 1
	curr := Coords{0, 0}
	for _, d := range moves {
		curr = Move(curr, d)
	}
	y, x := curr.Tuple()
	part1 := y * x

	// Part 2
	y, x, a := 0, 0, 0
	for _, d := range moves {
		dy, dx := d.Tuple()
		if dy == 0 {
			x += dx
			y += a * dx
		} else {
			a += dy
		}
	}
	part2 := y * x

	return NewSolution(part1, part2)
}

func data02(full bool) []Delta {
	return list.Map(ReadLines(21, 2, full), func(line string) Delta {
		p := str.SpaceSplit(line)
		cmd, x := p[0], number.ParseInt(p[1])
		switch cmd {
		case "forward":
			return Delta{0, x}
		case "up":
			return Delta{-x, 0}
		case "down":
			return Delta{x, 0}
		default:
			return X
		}
	})
}
