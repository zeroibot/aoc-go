package aoc21

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
)

func Day10() Solution {
	lines := data10(true)

	// Part 1
	score1 := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	total := 0
	for _, line := range lines {
		illegal := findIllegal(line)
		if dict.HasKey(score1, illegal) {
			total += score1[illegal]
		}
	}

	// Part 2
	score2 := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	scores := make([]int, 0)
	for _, line := range lines {
		incomplete := findIncomplete(line)
		if incomplete != nil {
			scores = append(scores, computeScore(*incomplete, score2))
		}
	}
	slices.Sort(scores)
	mid := len(scores) / 2
	midScore := scores[mid]

	return NewSolution(total, midScore)
}

func data10(full bool) []string {
	return ReadLines(21, 10, full)
}

var closer = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var opener = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func findIllegal(line string) rune {
	stack := ds.NewStack[rune]()
	for _, x := range line {
		if dict.HasKey(closer, x) {
			y, _ := stack.Pop()
			if y != closer[x] {
				return x
			}
		} else {
			stack.Push(x)
		}
	}
	return ' '
}

func findIncomplete(line string) *string {
	stack := ds.NewStack[rune]()
	for _, x := range line {
		if dict.HasKey(closer, x) {
			y, _ := stack.Pop()
			if y != closer[x] {
				return nil
			}
		} else {
			stack.Push(x)
		}
	}
	mirror := make([]rune, 0)
	for stack.Len() > 0 {
		top, _ := stack.Pop()
		mirror = append(mirror, opener[top])
	}
	incomplete := string(mirror)
	return &incomplete
}

func computeScore(text string, score map[rune]int) int {
	total := 0
	for _, x := range text {
		total = (total * 5) + score[x]
	}
	return total
}
