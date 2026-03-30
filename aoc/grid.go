package aoc

import (
	"math"

	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

var (
	U Delta = [2]int{-1, 0}
	D Delta = [2]int{1, 0}
	L Delta = [2]int{0, -1}
	R Delta = [2]int{0, 1}
	X Delta = [2]int{0, 0}
)

var (
	LeftOf  = map[Delta]Delta{U: L, L: D, D: R, R: U}
	RightOf = map[Delta]Delta{U: R, R: D, D: L, L: U}
)

func NewIntGrid(rows, cols, initial int) IntGrid {
	grid := make(IntGrid, rows)
	for row := range rows {
		line := make([]int, cols)
		for col := range cols {
			line[col] = initial
		}
		grid[row] = line
	}
	return grid
}

func Move(c Coords, d Delta) Coords {
	row, col := c.Tuple()
	dy, dx := d.Tuple()
	return Coords{row + dy, col + dx}
}

func GridSum(grid IntGrid) int {
	return list.Sum(list.Map(grid, func(line []int) int {
		return list.Sum(line)
	}))
}

func GridBounds[T any](items [][]T) Dims2 {
	rows, cols := len(items), len(items[0])
	return Dims2{rows, cols}
}

func StringGridBounds(grid []string) Dims2 {
	rows, cols := len(grid), len(grid[0])
	return Dims2{rows, cols}
}

func Manhattan(c1 Coords, c2 Coords) int {
	y1, x1 := c1.Tuple()
	y2, x2 := c2.Tuple()
	return number.Abs(y2-y1) + number.Abs(x2-x1)
}

func ManhattanOrigin(c Coords) int {
	return Manhattan(c, Coords{0, 0})
}

func Euclidean3(p1 Int3, p2 Int3) float64 {
	return math.Sqrt(list.Sum(list.Map(list.NumRange(0, 3), func(i int) float64 {
		return math.Pow(float64(p2[i])-float64(p1[i]), 2)
	})))
}

func InsideBounds(c Coords, maxBounds Dims2) bool {
	return InsideBoundsFull(c, maxBounds, Dims2{0, 0})
}

func InsideBoundsFull(c Coords, maxBounds Dims2, minBounds Dims2) bool {
	row, col := c.Tuple()
	minRows, minCols := minBounds.Tuple()
	numRows, numCols := maxBounds.Tuple()
	return minRows <= row && row < numRows && minCols <= col && col < numCols
}

func Surround4(c Coords) []Coords {
	y, x := c.Tuple()
	return []Coords{{y - 1, x}, {y, x - 1}, {y, x + 1}, {y + 1, x}}
}

func Surround8(c Coords) []Coords {
	y, x := c.Tuple()
	return []Coords{
		{y - 1, x - 1}, {y - 1, x}, {y - 1, x + 1},
		{y - 0, x - 1}, {y - 0, x + 1},
		{y + 1, x - 1}, {y + 1, x}, {y + 1, x + 1},
	}
}

func GetCoordsY(c Coords) int {
	return c[0]
}

func GetCoordsX(c Coords) int {
	return c[1]
}

func GetDelta(c1 Coords, c2 Coords) Delta {
	y1, x1 := c1.Tuple()
	y2, x2 := c2.Tuple()
	return Delta{y2 - y1, x2 - x1}
}
