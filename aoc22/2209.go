package aoc22

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day09() Solution {
	steps := data09(true)

	// Part 1
	head, tail := Coords{0, 0}, Coords{0, 0}
	visited1 := ds.SetFrom([]Coords{tail})

	// Part 2
	last := 9
	pos := make([]Coords, last+1)
	for i := range last + 1 {
		pos[i] = Coords{0, 0}
	}
	visited2 := ds.SetFrom([]Coords{pos[last]})

	for _, step := range steps {
		head, tail = moveRope(head, tail, step, visited1)
		pos = moveChain(pos, step, visited2)
	}

	count1 := visited1.Len()
	count2 := visited2.Len()

	return NewSolution(count1, count2)
}

func data09(full bool) []Step {
	return list.Map(ReadLines(22, 9, full), func(line string) Step {
		p := str.SpaceSplit(line)
		d := number.ParseInt(p[1])
		switch p[0] {
		case "U":
			return Step{U, d}
		case "D":
			return Step{D, d}
		case "L":
			return Step{L, d}
		case "R":
			return Step{R, d}
		}
		return Step{X, 0}
	})
}

type Step struct {
	Delta
	count int
}

func (s Step) Tuple() (Delta, int) {
	return s.Delta, s.count
}

func moveRope(head, tail Coords, step Step, visited *ds.Set[Coords]) (Coords, Coords) {
	d, count := step.Tuple()
	for range count {
		head = Move(head, d)
		if !isAdjacent(head, tail) {
			tail = follow(head, tail)
			visited.Add(tail)
		}
	}
	return head, tail
}

func moveChain(pos []Coords, step Step, visited *ds.Set[Coords]) []Coords {
	tail := len(pos) - 1
	d, count := step.Tuple()
	for range count {
		pos[0] = Move(pos[0], d)
		for _, i := range NumRange(1, tail+1) {
			if !isAdjacent(pos[i-1], pos[i]) {
				pos[i] = follow(pos[i-1], pos[i])
			}
		}
		visited.Add(pos[tail])
	}
	return pos
}

func isAdjacent(head, tail Coords) bool {
	y1, x1 := head.Tuple()
	y2, x2 := tail.Tuple()
	dy, dx := number.Abs(y2-y1), number.Abs(x2-x1)
	return dy <= 1 && dx <= 1
}

func follow(head, tail Coords) Coords {
	y1, x1 := head.Tuple()
	y2, x2 := tail.Tuple()
	dy := y1 - y2
	if dy > 0 {
		y2 += 1
	} else if dy < 0 {
		y2 -= 1
	}
	dx := x1 - x2
	if dx > 0 {
		x2 += 1
	} else if dx < 0 {
		x2 -= 1
	}
	return Coords{y2, x2}
}
