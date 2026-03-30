package aoc25

import (
	"cmp"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day03() Solution {
	banks := data03(true)
	total1 := list.Sum(list.Map(banks, computeTotalJolt(2)))
	total2 := list.Sum(list.Map(banks, computeTotalJolt(12)))
	return NewSolution(total1, total2)
}

func data03(full bool) [][]int {
	return list.Map(ReadLines(25, 3, full), ToIntLine)
}

func computeTotalJolt(numBat int) func([]int) int {
	return func(bank []int) int {
		n := len(bank)
		start := 0
		batteries := make([]int, 0)
		for d := range numBat {
			candidates := list.Map(list.NumRange(start, n-numBat+d+1), func(i int) [2]int {
				return [2]int{i, bank[i]}
			})
			best := slices.MaxFunc(candidates, func(a, b [2]int) int {
				score1 := cmp.Compare(a[1], b[1])
				if score1 != 0 {
					return score1
				}
				return cmp.Compare(b[0], a[0])
			})
			idx, battery := best[0], best[1]
			batteries = append(batteries, battery)
			start = idx + 1
		}
		digitString := strings.Join(list.Map(batteries, str.Int), "")
		return number.ParseInt(digitString)
	}
}
