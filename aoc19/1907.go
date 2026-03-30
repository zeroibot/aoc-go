package aoc19

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day07() Solution {
	numbers := data07(true)
	N := 5

	// Part 1
	options1 := []int{0, 1, 2, 3, 4}
	maxOutput1 := 0
	for _, settings := range Permutations(options1, N) {
		programs := make([]*Program, N)
		for i := range N {
			programs[i] = newProgram(numbers[:], settings[i])
		}
		programs[0].inputs = append(programs[0].inputs, 0)
		for i := range N {
			output, _ := runProgram07(programs[i])
			programs[i].output = output
			if i < N-1 && output != nil {
				programs[i+1].inputs = append(programs[i+1].inputs, *output)
			}
		}
		lastOutput := programs[N-1].output
		if lastOutput != nil {
			maxOutput1 = max(maxOutput1, *lastOutput)
		}
	}

	// Part 2
	options2 := []int{5, 6, 7, 8, 9}
	maxOutput2 := 0
	for _, settings := range Permutations(options2, N) {
		programs := make([]*Program, N)
		for i := range N {
			programs[i] = newProgram(numbers[:], settings[i])
		}
		programs[0].inputs = append(programs[0].inputs, 0)

		p := 0
		for {
			output, idx := runProgram07(programs[p])
			programs[p].index = idx
			programs[p].stop = output == nil
			if output != nil {
				programs[p].output = output
			}
			p = (p + 1) % N
			if output != nil {
				programs[p].inputs = append(programs[p].inputs, *output)
			}
			stop := list.All(programs, func(p *Program) bool {
				return p.stop
			})
			if stop {
				break
			}
		}
		lastOutput := programs[N-1].output
		if lastOutput != nil {
			maxOutput2 = max(maxOutput2, *lastOutput)
		}
	}

	return NewSolution(maxOutput1, maxOutput2)
}

func data07(full bool) []int {
	return ToIntList(ReadFirstLine(19, 7, full), ",")
}

func runProgram07(p *Program) (*int, int) {
	i := p.index
	for {
		head, tail := commandParts(p.numbers[i])
		cmd := number.ParseInt(tail)
		if cmd == 99 {
			break
		}

		switch cmd {
		case 1, 2, 7, 8:
			in1, in2, out := p.numbers[i+1], p.numbers[i+2], p.numbers[i+3]
			m := intcodeModes(head, 3)
			m1, m2 := m[0], m[1]
			a := intcodeParam(in1, m1, p.numbers)
			b := intcodeParam(in2, m2, p.numbers)
			switch cmd {
			case 1:
				p.numbers[out] = a + b
			case 2:
				p.numbers[out] = a * b
			case 7:
				p.numbers[out] = lang.Ternary(a < b, 1, 0)
			case 8:
				p.numbers[out] = lang.Ternary(a == b, 1, 0)
			}
			i += 4
		case 3:
			idx := p.numbers[i+1]
			p.numbers[idx], p.inputs = p.inputs[0], p.inputs[1:]
			i += 2
		case 4:
			m := intcodeModes(head, 1)[0]
			out := intcodeParam(p.numbers[i+1], m, p.numbers)
			i += 2
			return &out, i
		case 5, 6:
			p1, p2 := p.numbers[i+1], p.numbers[i+2]
			m := intcodeModes(head, 2)
			m1, m2 := m[0], m[1]
			isZero := intcodeParam(p1, m1, p.numbers) == 0
			doJump := lang.Ternary(cmd == 6, isZero, !isZero)
			if doJump {
				i = intcodeParam(p2, m2, p.numbers)
			} else {
				i += 3
			}
		}
	}
	return nil, i
}

type Program struct {
	numbers []int
	index   int
	stop    bool
	output  *int
	inputs  []int
}

func newProgram(numbers []int, input int) *Program {
	return &Program{
		numbers: numbers,
		index:   0,
		stop:    false,
		output:  nil,
		inputs:  []int{input},
	}
}
