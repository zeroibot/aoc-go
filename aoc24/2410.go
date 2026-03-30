package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"

	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

func Day10() Solution {
	grid := data10(true)
	start := make([]Coords, 0)
	for row, line := range grid {
		for col, x := range line {
			if x == 0 {
				start = append(start, Coords{row, col})
			}
		}
	}

	score := make([]int, 0)
	rating := make([]int, 0)
	for _, c := range start {
		reached := count9(c, grid)
		// Part 1
		score = append(score, len(reached))
		// Part 2
		rating = append(rating, list.Sum(dict.Values(reached)))
	}

	return NewSolution(list.Sum(score), list.Sum(rating))
}

func data10(full bool) [][]int {
	return list.Map(ReadLines(24, 10, full), ToIntLine)
}

func count9(start Coords, grid [][]int) map[Coords]int {
	bounds := GridBounds(grid)
	goals := make(map[Coords]int)
	q := ds.NewQueue[Coords]()
	q.Enqueue(start)
	for !q.IsEmpty() {
		c, _ := q.Dequeue()
		row, col := c.Tuple()
		value := grid[row][col]
		if value == 9 {
			goals[c] += 1
		} else {
			for _, nxt := range Surround4(c) {
				if !InsideBounds(nxt, bounds) {
					continue
				}
				row, col := nxt.Tuple()
				if grid[row][col] == value+1 {
					q.Enqueue(Coords{row, col})
				}
			}
		}
	}
	return goals
}
