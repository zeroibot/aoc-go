package aoc16

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
)

func Day07() Solution {
	words := data07(true)
	count1, count2 := 0, 0
	for _, word := range words {
		// Part 1
		if isValidWord(word) {
			count1 += 1
		}

		// Part 2
		if isValidWord2(word) {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data07(full bool) []string {
	return ReadLines(16, 7, full)
}

func isValidWord(word string) bool {
	found := false
	flip := false
	for i := range len(word) - 3 {
		if word[i] == '[' {
			flip = true
		} else if word[i] == ']' {
			flip = false
		} else if isABBA(word[i : i+4]) {
			if flip {
				return false
			}
			found = true
		}
	}
	return found
}

func isValidWord2(word string) bool {
	look := ds.NewSet[string]()
	found := ds.NewSet[string]()
	flip := false
	for i := range len(word) - 2 {
		sub := word[i : i+3]
		if word[i] == '[' {
			flip = true
		} else if word[i] == ']' {
			flip = false
		} else if isABA(sub) {
			if flip {
				found.Add(toABA(sub))
			} else {
				look.Add(sub)
			}
		}
	}
	common := look.Intersection(found)
	return common.Len() > 0
}

func isABBA(word string) bool {
	if len(word) != 4 {
		return false
	}
	ok1 := word[0] != word[1]
	ok2 := word[1] == word[2]
	ok3 := word[0] == word[3]
	return ok1 && ok2 && ok3
}

func isABA(word string) bool {
	if len(word) != 3 {
		return false
	}
	ok1 := word[0] != word[1]
	ok2 := word[0] == word[2]
	return ok1 && ok2
}

func toABA(bab string) string {
	b, a := bab[0], bab[1]
	return string([]byte{a, b, a})
}
