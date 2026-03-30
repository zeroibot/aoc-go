package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

func Day08() Solution {
	grid := data08(true)
	rows, cols := GridBounds(grid).Tuple()

	visible := ds.NewSet[Coords]()
	maxScore := 0
	for _, row := range NumRange(1, rows-1) {
		for _, col := range NumRange(1, cols-1) {
			pt := Coords{row, col}
			checkRowVisible(grid, visible, pt)
			checkColVisible(grid, visible, pt)
			maxScore = max(maxScore, computeGridScore(grid, row, col))
		}
	}
	numEdges := (2 * cols) + (2 * (rows - 2))
	count := numEdges + visible.Len()
	return NewSolution(count, maxScore)
}

func data08(full bool) IntGrid {
	return list.Map(ReadLines(22, 8, full), ToIntLine)
}

func checkRowVisible(grid IntGrid, visible *ds.Set[Coords], c Coords) {
	if visible.Has(c) {
		return
	}
	row, col := c.Tuple()
	isLower := func(x int) bool {
		return x < grid[row][col]
	}
	if list.All(grid[row][:col], isLower) {
		visible.Add(c)
		return
	}
	if list.All(grid[row][col+1:], isLower) {
		visible.Add(c)
	}
}

func checkColVisible(grid IntGrid, visible *ds.Set[Coords], c Coords) {
	if visible.Has(c) {
		return
	}
	row, col := c.Tuple()
	isLower := func(x int) bool {
		return x < grid[row][col]
	}
	above := list.Map(NumRange(0, row), func(r int) int {
		return grid[r][col]
	})
	if list.All(above, isLower) {
		visible.Add(c)
		return
	}
	below := list.Map(NumRange(row+1, len(grid)), func(r int) int {
		return grid[r][col]
	})
	if list.All(below, isLower) {
		visible.Add(c)
	}
}

func computeGridScore(grid IntGrid, row int, col int) int {
	rows, cols := GridBounds(grid).Tuple()
	value := grid[row][col]
	n, e, w, s := 0, 0, 0, 0
	for r := row - 1; r >= 0; r-- {
		n += 1
		if grid[r][col] >= value {
			break
		}
	}
	for _, r := range NumRange(row+1, rows) {
		s += 1
		if grid[r][col] >= value {
			break
		}
	}
	for c := col - 1; c >= 0; c-- {
		w += 1
		if grid[row][c] >= value {
			break
		}
	}
	for _, c := range NumRange(col+1, cols) {
		e += 1
		if grid[row][c] >= value {
			break
		}
	}
	return n * e * w * s
}
