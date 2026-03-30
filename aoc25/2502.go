package aoc25

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day02() Solution {
	ranges := data02(true)
	total1, total2 := 0, 0
	for _, r := range ranges {
		first, last := r[0], r[1]
		for _, x := range list.NumRange(first, last+1) {
			if isInvalid(x) {
				total1 += x
			}
			if isInvalid2(x) {
				total2 += x
			}
		}
	}
	return NewSolution(total1, total2)
}

func data02(full bool) [][2]int {
	ranges := make([][2]int, 0)
	for _, pair := range str.CleanSplit(ReadFirstLine(25, 2, full), ",") {
		p := str.CleanSplit(pair, "-")
		first, last := number.ParseInt(p[0]), number.ParseInt(p[1])
		ranges = append(ranges, [2]int{first, last})
	}
	return ranges
}

func isInvalid(x int) bool {
	s := str.Int(x)
	length := len(s)
	if length%2 != 0 {
		return false
	}
	half := length / 2
	return s[:half] == s[half:]
}

func isInvalid2(x int) bool {
	s := str.Int(x)
	length := len(s)
	half := length / 2
	for _, w := range list.NumRange(1, half+1) {
		if length%w != 0 {
			continue
		}
		repeat := length / w
		s2 := strings.Repeat(s[:w], repeat)
		if s == s2 {
			return true
		}
	}
	return false
}
