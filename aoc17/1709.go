package aoc17

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day09() Solution {
	stream := data09(true)

	// Part 1 and 2
	i, limit := 0, len(stream)
	count, total := 0, 0
	garbage := false
	stack := ds.NewStack[int]()
	for i < limit {
		char := stream[i]
		if garbage {
			switch char {
			case '!':
				i += 1
			case '>':
				garbage = false
			default:
				count += 1
			}
		} else if char == '{' {
			score := 1
			if !stack.IsEmpty() {
				if top, err := stack.Top(); err == nil {
					score = top + 1
				}
			}
			stack.Push(score)
		} else if char == '}' {
			if top, err := stack.Pop(); err == nil {
				total += top
			}
		} else if char == '<' {
			garbage = true
		}
		i += 1
	}
	return NewSolution(total, count)
}

func data09(full bool) string {
	return ReadFirstLine(17, 9, full)
}
