package aoc21

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

type Line = [2]Coords

func Day05() Solution {
	lines, bounds := data05(true)

	// Part 1
	count1 := countIntersection(lines, bounds, false)

	// Part 2
	count2 := countIntersection(lines, bounds, true)

	return NewSolution(count1, count2)
}

func data05(full bool) ([]Line, Dims2) {
	rows, cols := 0, 0
	lines := make([]Line, 0)
	for _, text := range ReadLines(21, 5, full) {
		p := str.CleanSplit(text, "->")
		x1, y1 := ToInt2(p[0], ",").Tuple()
		x2, y2 := ToInt2(p[1], ",").Tuple()
		rows = max(rows, y1, y2)
		cols = max(cols, x1, x2)
		lines = append(lines, Line{{y1, x1}, {y2, x2}})
	}
	return lines, Dims2{rows + 1, cols + 1}
}

func countIntersection(lines []Line, bounds Dims2, withDiagonal bool) int {
	rows, cols := bounds.Tuple()
	g := NewIntGrid(rows, cols, 0)
	for _, line := range lines {
		y1, x1 := line[0].Tuple()
		y2, x2 := line[1].Tuple()
		if x1 == x2 {
			addVertical(g, y1, y2, x1)
		} else if y1 == y2 {
			addHorizontal(g, x1, x2, y1)
		} else if withDiagonal && number.Abs(y2-y1) == number.Abs(x2-x1) {
			addDiagonal(g, x1, y1, x2, y2)
		}
	}
	total := 0
	for r := range rows {
		for c := range cols {
			if g[r][c] > 1 {
				total += 1
			}
		}
	}
	return total
}

func addVertical(g IntGrid, y1, y2, x int) {
	ys := []int{y1, y2}
	slices.Sort(ys)
	y1, y2 = ys[0], ys[1]
	for y := y1; y <= y2; y++ {
		g[y][x] += 1
	}
}

func addHorizontal(g IntGrid, x1, x2, y int) {
	xs := []int{x1, x2}
	slices.Sort(xs)
	x1, x2 = xs[0], xs[1]
	for x := x1; x <= x2; x++ {
		g[y][x] += 1
	}
}

func addDiagonal(g IntGrid, x1, y1, x2, y2 int) {
	xs := lang.Ternary(x1 < x2, NumRange(x1, x2+1), RevRange(x1, x2-1))
	ys := lang.Ternary(y1 < y2, NumRange(y1, y2+1), RevRange(y1, y2-1))
	for i := range len(xs) {
		y, x := ys[i], xs[i]
		g[y][x] += 1
	}
}
