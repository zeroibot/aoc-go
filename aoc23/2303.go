package aoc23

import (
	"regexp"
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day03() Solution {
	grid := data03(true)

	// Part 1
	symbols := findSymbols(grid)
	total1 := list.Sum(findValidNumbers(grid, symbols))

	// Part 2
	gears := findGears(grid)
	total2 := list.Sum(findGearRatios(grid, gears))

	return NewSolution(total1, total2)
}

func data03(full bool) []string {
	return ReadLines(23, 3, full)
}

func findSymbols(grid []string) []Coords {
	symbols := ds.NewSet[Coords]()
	nonSymbol := []rune("0123456789.")
	for row, line := range grid {
		for col, char := range line {
			if !slices.Contains(nonSymbol, char) {
				symbols.Add(Coords{row, col})
			}
		}
	}
	return symbols.Items()
}

func findValidNumbers(grid []string, symbols []Coords) []int {
	bounds := StringGridBounds(grid)
	numberPattern := regexp.MustCompile(`[0-9]+`)
	numbers := make([]int, 0)
	for row, line := range grid {
		for _, m := range numberPattern.FindAllIndex([]byte(line), -1) {
			start, end := m[0], m[1]
			rowRange := Int3{row, start, end}
			if hasAdjacentSymbol(rowRange, symbols, bounds) {
				number := number.ParseInt(line[start:end])
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

func findGears(grid []string) []Coords {
	gears := make([]Coords, 0)
	for row, line := range grid {
		for col, char := range line {
			if char == '*' {
				gears = append(gears, Coords{row, col})
			}
		}
	}
	return gears
}

func findGearRatios(grid []string, gears []Coords) []int {
	bounds := StringGridBounds(grid)
	numberPattern := regexp.MustCompile(`[0-9]+`)
	adjacent := make(map[Coords][]int)
	for row, line := range grid {
		for _, m := range numberPattern.FindAllIndex([]byte(line), -1) {
			start, end := m[0], m[1]
			rowRange := Int3{row, start, end}
			number := number.ParseInt(line[start:end])
			for _, c := range getAdjacentSymbols(rowRange, gears, bounds) {
				adjacent[c] = append(adjacent[c], number)
			}
		}
	}

	numbers := make([]int, 0)
	for _, near := range adjacent {
		if len(near) == 2 {
			a, b := near[0], near[1]
			numbers = append(numbers, a*b)
		}
	}
	return numbers
}

func hasAdjacentSymbol(rowRange Int3, symbols []Coords, bounds Dims2) bool {
	return len(getAdjacentSymbols(rowRange, symbols, bounds)) > 0
}

func getAdjacentSymbols(rowRange Int3, symbols []Coords, bounds Dims2) []Coords {
	adjacent := getAdjacent(rowRange, bounds)
	set1 := ds.SetFrom(symbols)
	set2 := ds.SetFrom(adjacent)
	common := set1.Intersection(set2)
	return common.Items()
}

func getAdjacent(rowRange Int3, bounds Dims2) []Coords {
	y1, x1, x2 := rowRange.Tuple()
	rows, cols := bounds.Tuple()
	y0, y2 := y1-1, y1+1
	x0, x3 := x1-1, x2+1
	adjacent := make([]Coords, 0)

	start := lang.Ternary(x0 >= 0, x0, x1)
	end := lang.Ternary(x3 <= cols, x3, x2)
	addAbove := y0 >= 0
	addBelow := y2 < rows
	for _, x := range NumRange(start, end) {
		if addAbove {
			adjacent = append(adjacent, Coords{y0, x})
		}
		if addBelow {
			adjacent = append(adjacent, Coords{y2, x})
		}
	}
	if start != x1 {
		adjacent = append(adjacent, Coords{y1, x0})
	}
	if x2 < cols {
		adjacent = append(adjacent, Coords{y1, x2})
	}
	return adjacent
}
