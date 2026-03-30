package aoc20

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day08() Solution {
	codes := data08(true)

	// Part 1
	value1, _ := runCodes(codes)

	// Part 2
	value2 := 0
	for i, code := range codes {
		cmd := code.Str
		if cmd == "acc" {
			continue // skip , not corrupted
		}
		codes2 := flipNopJmp(codes, i)
		value, stuck := runCodes(codes2)
		if !stuck {
			value2 = value
			break
		}
	}

	return NewSolution(value1, value2)
}

func data08(full bool) []StrInt {
	return list.Map(ReadLines(20, 8, full), func(line string) StrInt {
		p := str.SpaceSplit(line)
		return StrInt{Str: p[0], Int: number.ParseInt(p[1])}
	})
}

func runCodes(codes []StrInt) (int, bool) {
	i, x := 0, 0
	done := ds.NewSet[int]()
	numCodes := len(codes)
	for {
		if done.Has(i) {
			return x, true
		}
		if i >= numCodes {
			return x, false
		}

		done.Add(i)
		cmd, y := codes[i].Tuple()
		switch cmd {
		case "nop":
			i += 1
		case "acc":
			x += y
			i += 1
		case "jmp":
			i += y
		}
	}
}

func flipNopJmp(codes []StrInt, idx int) []StrInt {
	codes2 := make([]StrInt, len(codes))
	copy(codes2, codes)
	cmd, y := codes2[idx].Tuple()
	cmd2 := lang.Ternary(cmd == "jmp", "nop", "jmp")
	codes2[idx] = StrInt{Str: cmd2, Int: y}
	return codes2
}
