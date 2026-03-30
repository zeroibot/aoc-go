package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day04() Solution {
	grid := data04(true)
	bounds := StringGridBounds(grid)

	// Part 1
	points := make([]Coords, 0)
	for row, line := range grid {
		for col, tile := range line {
			if tile == 'X' {
				points = append(points, Coords{row, col})
			}
		}
	}
	vectors := make([]Vector, 0)
	for _, center := range points {
		for _, pt := range Surround8(center) {
			if !InsideBounds(pt, bounds) {
				continue
			}
			row, col := pt.Tuple()
			if grid[row][col] == 'M' {
				vector := Vector{Coords: pt, Delta: GetDelta(center, pt)}
				vectors = append(vectors, vector)
			}
		}
	}
	for _, letter := range []byte("AS") {
		vectors = findNextPositions(grid, bounds, vectors, letter)
	}
	count1 := len(vectors)

	// Part 2
	rows, cols := bounds.Tuple()
	minBounds := Dims2{1, 1}
	maxBounds := Dims2{rows - 1, cols - 1}
	points = make([]Coords, 0)
	for row, line := range grid {
		for col, tile := range line {
			pt := Coords{row, col}
			if tile == 'A' && InsideBoundsFull(pt, maxBounds, minBounds) {
				points = append(points, pt)
			}
		}
	}
	count2 := 0
	for _, pt := range points {
		row, col := pt.Tuple()
		// Left diagonal
		tl := grid[row-1][col-1]
		br := grid[row+1][col+1]
		ldiag := string([]byte{tl, 'A', br})
		// Right diagonal
		tr := grid[row-1][col+1]
		bl := grid[row+1][col-1]
		rdiag := string([]byte{tr, 'A', bl})
		if isXMAS(ldiag, rdiag) {
			count2 += 1
		}
	}

	return NewSolution(count1, count2)
}

func data04(full bool) []string {
	return ReadLines(24, 4, full)
}

func findNextPositions(grid []string, bounds Dims2, vectors []Vector, letter byte) []Vector {
	vectors2 := make([]Vector, 0)
	for _, vector := range vectors {
		c, d := vector.Tuple()
		c = Move(c, d)
		if !InsideBounds(c, bounds) {
			continue
		}
		row, col := c.Tuple()
		if grid[row][col] == letter {
			vectors2 = append(vectors2, Vector{Coords: c, Delta: d})
		}
	}
	return vectors2
}

func isXMAS(diag1 string, diag2 string) bool {
	diags := []string{diag1, diag2}
	return list.All(diags, func(diag string) bool {
		return diag == "MAS" || diag == "SAM"
	})
}
