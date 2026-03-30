package aoc21

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day03() Solution {
	binaryNumbers := data03(true)

	// Part 1
	count := make(map[int]int)
	for _, code := range binaryNumbers {
		for i, bit := range code {
			if bit == '1' {
				count[i] += 1
			}
		}
	}
	mid := len(binaryNumbers) / 2
	g := make([]rune, 0)
	e := make([]rune, 0)
	for i := range len(binaryNumbers[0]) {
		if count[i] > mid {
			g = append(g, '1')
			e = append(e, '0')
		} else {
			g = append(g, '0')
			e = append(e, '1')
		}
	}
	gb := number.ParseBinary(string(g))
	eb := number.ParseBinary(string(e))
	part1 := gb * eb

	// Part 2
	oxy := filterMax(binaryNumbers)
	co2 := filterMin(binaryNumbers)
	part2 := oxy * co2

	return NewSolution(part1, part2)
}

func data03(full bool) []string {
	return ReadLines(21, 3, full)
}

func countIndex(binaryNumbers []string, index int) int {
	valid := list.Filter(binaryNumbers, func(code string) bool {
		return code[index] == '1'
	})
	return len(valid)
}

func filterMax(binaryNumbers []string) int {
	bitLength := len(binaryNumbers[0])
	for i := range bitLength {
		c1 := countIndex(binaryNumbers, i)
		c0 := len(binaryNumbers) - c1
		maxBit := lang.Ternary(c1 >= c0, '1', '0')
		binaryNumbers = list.Filter(binaryNumbers, func(code string) bool {
			return code[i] == byte(maxBit)
		})
		if len(binaryNumbers) == 1 {
			break
		}
	}
	return number.ParseBinary(binaryNumbers[0])
}

func filterMin(binaryNumbers []string) int {
	bitLength := len(binaryNumbers[0])
	for i := range bitLength {
		c1 := countIndex(binaryNumbers, i)
		c0 := len(binaryNumbers) - c1
		minBit := lang.Ternary(c0 <= c1, '0', '1')
		binaryNumbers = list.Filter(binaryNumbers, func(code string) bool {
			return code[i] == byte(minBit)
		})
		if len(binaryNumbers) == 1 {
			break
		}
	}
	return number.ParseBinary(binaryNumbers[0])
}
