package aoc22

import (
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day05() Solution {
	// Part 1
	stacks, moves := data05(true)
	top1 := processMoves(stacks, moves, true)

	// Part 2
	stacks, moves = data05(true)
	top2 := processMoves(stacks, moves, false)

	return NewSolution(top1, top2)
}

func data05(full bool) ([][]rune, []Int3) {
	stacks := make([][]rune, 0)
	moves := make([]Int3, 0)
	stackMode := true
	for _, line := range ReadRawLines(22, 5, full, false) {
		clean := strings.TrimSpace(line)
		if clean == "" {
			stackMode = false
			continue
		}
		if stackMode {
			if !strings.HasPrefix(clean, "[") {
				continue
			}
			if len(stacks) == 0 {
				count := len(line) / 4
				for range count {
					stacks = append(stacks, []rune{})
				}
			}
			for i, char := range line {
				if i%4 != 1 || char == ' ' {
					continue
				}
				idx := i / 4
				stacks[idx] = append(stacks[idx], char)
			}
		} else {
			p := str.SpaceSplit(line)
			count := number.ParseInt(p[1])
			src := number.ParseInt(p[3]) - 1
			dst := number.ParseInt(p[5]) - 1
			moves = append(moves, Int3{count, src, dst})
		}
	}
	return stacks, moves
}

func processMoves(stacks [][]rune, moves []Int3, reverse bool) string {
	for _, move := range moves {
		count, idx1, idx2 := move.Tuple()
		transfer(stacks, count, idx1, idx2, reverse)
	}
	top := list.Map(stacks, func(stack []rune) rune {
		return stack[0]
	})
	return string(top)
}

func transfer(stacks [][]rune, count int, idx1 int, idx2 int, reverse bool) {
	s1, s2 := stacks[idx1], stacks[idx2]
	move := list.Copy(s1[:count])
	if reverse {
		slices.Reverse(move)
	}
	move = append(move, s2...)
	stacks[idx1] = s1[count:]
	stacks[idx2] = move
}
