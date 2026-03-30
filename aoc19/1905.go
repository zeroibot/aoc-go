package aoc19

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/number"
)

func Day05() Solution {
	// Part 1
	numbers := data05(true)
	out1 := runProgram05(numbers, 1)

	// Part 2
	numbers = data05(true)
	out2 := runProgram05(numbers, 5)

	return NewSolution(out1, out2)
}

func data05(full bool) []int {
	return ToIntList(ReadFirstLine(19, 5, full), ",")
}

func runProgram05(numbers []int, start int) int {
	i := 0
	output := 0
	for {
		head, tail := commandParts(numbers[i])
		cmd := number.ParseInt(tail)
		if cmd == 99 {
			break
		}

		switch cmd {
		case 1, 2, 7, 8:
			in1, in2, out := numbers[i+1], numbers[i+2], numbers[i+3]
			m := intcodeModes(head, 3)
			m1, m2 := m[0], m[1]
			a := intcodeParam(in1, m1, numbers)
			b := intcodeParam(in2, m2, numbers)
			switch cmd {
			case 1:
				numbers[out] = a + b
			case 2:
				numbers[out] = a * b
			case 7:
				numbers[out] = lang.Ternary(a < b, 1, 0)
			case 8:
				numbers[out] = lang.Ternary(a == b, 1, 0)
			}
			i += 4
		case 3:
			idx := numbers[i+1]
			numbers[idx] = start
			i += 2
		case 4:
			m := intcodeModes(head, 1)[0]
			out := intcodeParam(numbers[i+1], m, numbers)
			if out != 0 {
				output = out
			}
			i += 2
		case 5, 6:
			p1, p2 := numbers[i+1], numbers[i+2]
			m := intcodeModes(head, 2)
			m1, m2 := m[0], m[1]
			isZero := intcodeParam(p1, m1, numbers) == 0
			doJump := lang.Ternary(cmd == 6, isZero, !isZero)
			if doJump {
				i = intcodeParam(p2, m2, numbers)
			} else {
				i += 3
			}
		}
	}
	return output
}
