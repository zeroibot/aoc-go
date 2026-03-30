package aoc23

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
)

func Day01() Solution {
	lines := data01(true)

	total1, total2 := 0, 0
	for _, line := range lines {
		// Part 1
		total1 += extractDigits(line)

		// Part 2
		total2 += extractNumber(line)
	}

	return NewSolution(total1, total2)
}

func data01(full bool) []string {
	return ReadLines(23, 1, full)
}

func extractDigits(line string) int {
	first, last := 0, 0
	for _, x := range line {
		digit := parseDigit(x)
		if digit >= 0 {
			first, last = updateDigits(first, last, digit)
		}
	}
	return (first * 10) + last
}

func extractNumber(line string) int {
	first, last := 0, 0
	for i, x := range line {
		digit := parseDigit(x)
		if digit >= 0 {
			first, last = updateDigits(first, last, digit)
			continue
		}
		digit = parseNumber(line[i:])
		if digit >= 0 {
			first, last = updateDigits(first, last, digit)
		}
	}
	return (first * 10) + last
}

func parseDigit(x rune) int {
	v := int(x)
	if 48 <= v && v <= 57 {
		return v - 48
	}
	return -1
}

var numberWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseNumber(text string) int {
	for word, number := range numberWords {
		if strings.HasPrefix(text, word) {
			return number
		}
	}
	return -1
}

func updateDigits(first, last, digit int) (int, int) {
	if first == 0 {
		first = digit
	}
	last = digit
	return first, last
}
