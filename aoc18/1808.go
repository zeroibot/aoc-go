package aoc18

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day08() Solution {
	numbers := data08(true)
	limit := len(numbers)

	// Part 1
	child, meta := numbers[0], numbers[1]
	stack := ds.NewStack[Int2]()
	stack.Push(Int2{child, meta})
	total, i := 0, 2
	for i < limit {
		top, _ := stack.Top()
		if top[0] == 0 {
			top, _ := stack.Pop()
			meta = top[1]
			for m := range meta {
				total += numbers[i+m]
			}
			i += meta
			if stack.IsEmpty() {
				continue
			}
			top, _ = stack.Pop()
			child, meta = top.Tuple()
			stack.Push(Int2{child - 1, meta})
		} else {
			child, meta = numbers[i], numbers[i+1]
			stack.Push(Int2{child, meta})
			i += 2
		}
	}

	// Part 2
	child, meta = numbers[0], numbers[1]
	stack2 := ds.NewStack[Triple]()
	stack2.Push(Triple{child, meta, []int{}})
	value, i := 0, 2
	for i < limit {
		top, _ := stack2.Top()
		if top.child == len(top.values) {
			top, _ = stack2.Pop()
			child, meta, values := top.child, top.meta, top.values
			hasChild := child > 0
			value = 0
			for m := range meta {
				if hasChild {
					idx := numbers[i+m] - 1
					if idx < len(values) {
						value += values[idx]
					}
				} else {
					value += numbers[i+m]
				}
			}
			i += meta
			if stack2.IsEmpty() {
				break
			}
			top, _ = stack2.Pop()
			top.values = append(top.values, value)
			stack2.Push(top)
		} else {
			child, meta = numbers[i], numbers[i+1]
			stack2.Push(Triple{child, meta, []int{}})
			i += 2
		}
	}

	return NewSolution(total, value)
}

func data08(full bool) []int {
	line := ReadFirstLine(18, 8, full)
	return ToIntList(line, " ")
}

type Triple struct {
	child  int
	meta   int
	values []int
}
