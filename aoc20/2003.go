package aoc20

import (
	"math"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day03() Solution {
	// Part 1
	count := countSlope(Delta{1, 3}, true)

	// Part 2
	deltas := []Delta{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	product := 1
	for _, d := range deltas {
		product *= countSlope(d, true)
	}

	return NewSolution(count, product)
}

func data03(full bool, d Delta) []string {
	lines := ReadLines(20, 3, full)
	dy, dx := d.Tuple()
	h, w := len(lines), len(lines[0])
	needW := float64((1 + dx) * numSteps(h, dy))
	repeat := int(math.Ceil(needW / float64(w)))
	return list.Map(lines, func(line string) string {
		return strings.Repeat(line, repeat)
	})
}

func numSteps(height int, dy int) int {
	return (height - 1) / dy
}

func countSlope(d Delta, full bool) int {
	g := data03(full, d)
	curr := Coords{0, 0}
	count := 0
	height, dy := len(g), d[0]
	steps := numSteps(height, dy)
	for range steps {
		curr = Move(curr, d)
		row, col := curr.Tuple()
		if g[row][col] == '#' {
			count += 1
		}
	}
	return count
}
