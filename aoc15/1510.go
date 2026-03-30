package aoc15

import (
	"fmt"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
)

func Day10() Solution {
	text := data10(true)

	// Part 1
	length1 := repeatExpand(text, 40)

	// Part 2
	length2 := repeatExpand(text, 50)

	return NewSolution(length1, length2)
}

func data10(full bool) string {
	return ReadFirstLine(15, 10, full)
}

func repeatExpand(text string, count int) int {
	curr := text
	for range count {
		next := make([]string, 0)
		d, r := curr[0], 1
		for i := range len(curr) {
			if i == 0 {
				continue
			}
			x := curr[i]
			if x == d {
				r += 1
			} else {
				next = append(next, fmt.Sprintf("%d", r))
				next = append(next, string(d))
				d, r = x, 1
			}
		}
		next = append(next, fmt.Sprintf("%d", r))
		next = append(next, string(d))
		curr = strings.Join(next, "")
	}
	return len(curr)
}
