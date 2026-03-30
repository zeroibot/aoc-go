package aoc15

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

func Day06() Solution {
	side := 1000

	// Part 1
	mask := map[string]int{
		on:     1,
		off:    0,
		toggle: -1,
	}
	commands := data06(mask, true)
	grid := NewIntGrid(side, side, 0)
	for _, cmd := range commands {
		b, y1, x1, y2, x2 := cmd.Tuple()
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				if b == -1 {
					grid[y][x] ^= 1 // flip bit
				} else {
					grid[y][x] = b
				}
			}
		}
	}
	part1 := GridSum(grid)

	// Part 2
	mask = map[string]int{
		on:     1,
		off:    -1,
		toggle: 2,
	}
	commands = data06(mask, true)
	grid = NewIntGrid(side, side, 0)
	for _, cmd := range commands {
		b, y1, x1, y2, x2 := cmd.Tuple()
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				value := grid[y][x] + b
				grid[y][x] = max(value, 0)
			}
		}
	}
	part2 := GridSum(grid)

	return NewSolution(part1, part2)
}

var (
	on     string = "turn on"
	off    string = "turn off"
	toggle string = "toggle"
)

type command [5]int

func (cmd command) Tuple() (int, int, int, int, int) {
	return cmd[0], cmd[1], cmd[2], cmd[3], cmd[4]
}

func data06(mask map[string]int, full bool) []command {
	return list.Map(ReadLines(15, 6, full), func(line string) command {
		b := 0
		if strings.HasPrefix(line, on) {
			b = mask[on]
		} else if strings.HasPrefix(line, off) {
			b = mask[off]
		} else if strings.HasPrefix(line, toggle) {
			b = mask[toggle]
		}
		p := str.CleanSplit(line, "through")
		head := Last(str.SpaceSplit(p[0]), 1)
		tail := p[1]
		c1 := ToInt2(head, ",")
		c2 := ToInt2(tail, ",")
		x1, y1 := c1.Tuple()
		x2, y2 := c2.Tuple()
		return command{b, y1, x1, y2 + 1, x2 + 1}
	})
}
