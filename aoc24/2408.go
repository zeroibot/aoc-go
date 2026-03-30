package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day08() Solution {
	antenna, bounds := day08(true)

	// Part 1
	count1 := countAntiNodes(antenna, bounds, false)

	// Part 2
	count2 := countAntiNodes(antenna, bounds, true)

	return NewSolution(count1, count2)
}

func day08(full bool) (map[rune][]Coords, Dims2) {
	lines := ReadLines(24, 8, full)
	bounds := StringGridBounds(lines)
	antenna := make(map[rune][]Coords)
	for row, line := range lines {
		for col, char := range line {
			if char == '.' {
				continue
			}
			antenna[char] = append(antenna[char], Coords{row, col})
		}
	}
	return antenna, bounds
}

func countAntiNodes(antenna map[rune][]Coords, bounds Dims2, extend bool) int {
	anti := ds.NewSet[Coords]()
	for _, positions := range antenna {
		for _, pair := range Combinations(positions, 2) {
			y1, x1 := pair[0].Tuple()
			y2, x2 := pair[1].Tuple()
			d1 := Delta{y1 - y2, x1 - x2}
			d2 := Delta{y2 - y1, x2 - x1}
			pairs := map[Coords]Delta{
				pair[0]: d1,
				pair[1]: d2,
			}
			for a, d := range pairs {
				for {
					a = Move(a, d)
					if !InsideBounds(a, bounds) {
						break
					}
					anti.Add(a)
					if !extend {
						break
					}
				}
			}
		}
	}
	if extend {
		for _, positions := range antenna {
			anti = anti.Union(ds.SetFrom(positions))
		}
	}
	return anti.Len()
}
