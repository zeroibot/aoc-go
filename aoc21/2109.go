package aoc21

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

func Day09() Solution {
	grid := data09(true)

	// Part 1
	low := findLowPoints(grid)
	total := list.Sum(list.Map(low, func(pt Int3) int {
		return pt[2] + 1
	}))

	// Part 2
	basins := list.Map(low, func(t Int3) int {
		y, x, _ := t.Tuple()
		return basinSize(grid, Coords{y, x})
	})
	slices.Sort(basins)
	b1, b2, b3 := Last(basins, 1), Last(basins, 2), Last(basins, 3)
	product := b1 * b2 * b3

	return NewSolution(total, product)
}

func data09(full bool) IntGrid {
	return list.Map(ReadLines(21, 9, full), ToIntLine)
}

func findLowPoints(grid IntGrid) []Int3 {
	low := make([]Int3, 0)
	for row, line := range grid {
		for col, height := range line {
			allHigher := list.All(surrounding(grid, Coords{row, col}), func(t Int3) bool {
				return t[2] > height // compare height of surrounding to current
			})
			if allHigher {
				low = append(low, Int3{row, col, height})
			}
		}
	}
	return low
}

func surrounding(grid IntGrid, center Coords) []Int3 {
	bounds := GridBounds(grid)
	near := make([]Int3, 0)
	for _, nxt := range Surround4(center) {
		if !InsideBounds(nxt, bounds) {
			continue
		}
		y, x := nxt.Tuple()
		near = append(near, Int3{y, x, grid[y][x]})
	}
	return near
}

func basinSize(grid IntGrid, center Coords) int {
	visited := ds.NewSet[Coords]()
	stack := ds.NewStack[Coords]()
	stack.Push(center)
	for stack.Len() > 0 {
		c, _ := stack.Pop()
		visited.Add(c)
		for _, t := range surrounding(grid, c) {
			y, x, h := t.Tuple()
			if h == 9 || visited.Has(Coords{y, x}) {
				continue
			}
			stack.Push(Coords{y, x})
		}
	}
	return visited.Len()
}
