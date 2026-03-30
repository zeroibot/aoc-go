package aoc17

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day08() Solution {
	instructions := data08(true)
	reg := make(map[string]int)
	maxVal := 0
	for _, cmd := range instructions {
		if cmd.isSatisfied(reg[cmd.condReg]) {
			reg[cmd.targetReg] += cmd.inc
			// Part 2
			maxVal = max(maxVal, reg[cmd.targetReg])
		}
	}
	// Part 1
	maxReg := slices.Max(dict.Values(reg))
	return NewSolution(maxReg, maxVal)
}

func data08(full bool) []Instruction {
	return list.Map(ReadLines(17, 8, full), func(line string) Instruction {
		cmd := Instruction{}
		p := str.CleanSplit(line, " if ")
		head, tail := p[0], p[1]
		h := str.SpaceSplit(head)
		cmd.targetReg = h[0]
		cmd.inc = number.ParseInt(h[2])
		if h[1] == "dec" {
			cmd.inc *= -1
		}
		t := str.SpaceSplit(tail)
		cmd.condReg = t[0]
		cmd.op = t[1]
		cmd.value = number.ParseInt(t[2])
		return cmd
	})
}

type Instruction struct {
	targetReg string
	condReg   string
	op        string
	inc       int
	value     int
}

func (i Instruction) isSatisfied(value int) bool {
	switch i.op {
	case "==":
		return value == i.value
	case "!=":
		return value != i.value
	case ">":
		return value > i.value
	case ">=":
		return value >= i.value
	case "<":
		return value < i.value
	case "<=":
		return value <= i.value
	default:
		return false
	}
}
