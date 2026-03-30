package aoc15

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

func Day03() Solution {
	moves := data03(true)

	// Part 1
	visited := walk(moves)
	part1 := visited.Len()

	// Part 2
	limit := len(moves)
	santa := make([]Delta, 0)
	robo := make([]Delta, 0)
	for i := 0; i < limit; i += 2 {
		santa = append(santa, moves[i])
		robo = append(robo, moves[i+1])
	}
	visited1 := walk(santa)
	visited2 := walk(robo)
	visited = visited1.Union(visited2)
	part2 := visited.Len()

	return NewSolution(part1, part2)
}

func data03(full bool) []Delta {
	T := map[rune]Delta{
		'>': R,
		'<': L,
		'^': U,
		'v': D,
	}
	line := ReadFirstLine(15, 3, full)
	return list.Translate([]rune(line), T)
}

func walk(moves []Delta) *ds.Set[Coords] {
	start := Coords{0, 0}
	visited := ds.NewSet[Coords]()
	visited.Add(start)
	curr := start
	for _, delta := range moves {
		curr = Move(curr, delta)
		visited.Add(curr)
	}
	return visited
}
