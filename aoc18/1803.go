package aoc18

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day03() Solution {
	claims := data03(true)

	// Part 1
	g := NewIntGrid(1000, 1000, 0)
	for _, claim := range claims {
		row, col := claim.coords.Tuple()
		h, w := claim.dims.Tuple()
		for dy := range h {
			r := row + dy
			for dx := range w {
				c := col + dx
				g[r][c] += 1
			}
		}
	}
	total := list.Sum(list.Map(g, func(row []int) int {
		count := 0
		for _, x := range row {
			if x > 1 {
				count += 1
			}
		}
		return count
	}))

	// Part 2
	g = NewIntGrid(1000, 1000, 0)
	clean := ds.NewSet[int]()
	for _, claim := range claims {
		row, col := claim.coords.Tuple()
		h, w := claim.dims.Tuple()
		ok := true
		for dy := range h {
			r := row + dy
			for dx := range w {
				c := col + dx
				if g[r][c] == 0 {
					g[r][c] = claim.id
				} else {
					ok = false
					owner := g[r][c]
					if clean.Has(owner) {
						clean.Delete(owner)
					}
				}
			}
		}
		if ok {
			clean.Add(claim.id)
		}
	}
	id := clean.Items()[0]

	return NewSolution(total, id)
}

func data03(full bool) []Claim {
	return list.Map(ReadLines(18, 3, full), func(line string) Claim {
		claim := Claim{}
		p := str.CleanSplit(line, "@")
		claim.id = number.ParseInt(strings.TrimPrefix(p[0], "#"))
		t := str.CleanSplit(p[1], ":")
		c := ToIntList(t[0], ",")
		claim.coords = Coords{c[1], c[0]}
		d := ToIntList(t[1], "x")
		claim.dims = Dims2{d[1], d[0]}
		return claim
	})
}

type Claim struct {
	id     int
	coords Coords
	dims   Dims2
}
