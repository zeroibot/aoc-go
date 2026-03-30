package aoc17

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day10() Solution {
	line := data10(true)

	// Part 1
	lengths := ToIntList(line, ",")
	numbers := knotHash(lengths, 1)
	product := numbers[0] * numbers[1]

	// Part 2
	lengths = list.Map([]byte(line), func(b byte) int {
		return int(b)
	})
	numbers = knotHash(lengths, 64)
	result := make([]string, 0)
	for i := 0; i < knotHashLimit; i += 16 {
		r := numbers[i]
		for j := 1; j < 16; j++ {
			r ^= numbers[i+j]
		}
		result = append(result, hexCode(r))
	}
	code := strings.Join(result, "")

	return NewSolution(product, code)
}

func data10(full bool) string {
	return ReadFirstLine(17, 10, full)
}
