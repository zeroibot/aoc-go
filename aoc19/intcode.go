package aoc19

import (
	"slices"
	"strconv"

	"github.com/zeroibot/fn/number"
)

func commandParts(number int) (string, string) {
	word := strconv.Itoa(number)
	n := len(word)
	if n <= 2 {
		return "", word
	} else {
		return word[:n-2], word[n-2:]
	}
}

func intcodeModes(cmd string, count int) []int {
	m := make([]int, count)
	i := 0
	digits := []rune(cmd)
	slices.Reverse(digits)
	for _, x := range digits {
		m[i] = number.ParseInt(string(x))
		i += 1
	}
	return m
}

func intcodeParam(x int, mode int, numbers []int) int {
	if mode == 0 {
		return numbers[x]
	}
	return x
}

func intcodeParam2(x int, mode int, rbase int, numbers map[int]int) int {
	switch mode {
	case 0, 2:
		idx := intcodeIndex(x, mode, rbase)
		return numbers[idx]
	case 1:
		return x
	}
	return 0
}

func intcodeIndex(x int, mode int, rbase int) int {
	idx := x
	if mode == 2 {
		idx += rbase
	}
	return idx
}
