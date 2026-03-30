package aoc25

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/check"
	"github.com/zeroibot/fn/list"
)

func Day04() Solution {
	grid := data04(true)
	bounds := GridBounds(grid)
	total1, total2 := 0, 0
	for {
		grid2 := make([][]bool, len(grid))
		count := 0
		for row, line := range grid {
			line2 := list.Copy(line)
			for col, paper := range line {
				if !paper {
					continue
				}
				paperNeighbors := list.Filter(Surround8(Coords{row, col}), func(coords Coords) bool {
					r, c := coords.Tuple()
					return InsideBounds(coords, bounds) && grid[r][c]
				})
				if len(paperNeighbors) < 4 {
					line2[col] = false
					count += 1
				}
			}
			grid2[row] = line2
		}
		grid = grid2
		if total1 == 0 {
			total1 = count
		}
		total2 += count
		if count == 0 {
			break
		}
	}
	return NewSolution(total1, total2)
}

func data04(full bool) [][]bool {
	var paper byte = '@'
	return list.Map(ReadLines(25, 4, full), func(line string) []bool {
		return list.Map([]byte(line), check.IsEqual(paper))
	})
}
