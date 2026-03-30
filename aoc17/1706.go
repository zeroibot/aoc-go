package aoc17

import (
	"fmt"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day06() Solution {
	banks := data06(true)
	numBanks := len(banks)
	done := make(map[string]int)
	state := bankState(banks)
	count := 0
	done[state] = count
	for {
		index := ArgMax(banks)
		take := banks[index]
		banks[index] = 0
		for i := range take {
			idx := (index + i + 1) % numBanks
			banks[idx] += 1
		}

		count += 1
		state = bankState(banks)
		if stateCount, ok := done[state]; ok {
			// Part 1 and 2
			return NewSolution(count, count-stateCount)
		}
		done[state] = count
	}
}

func data06(full bool) []int {
	return ToIntList(ReadFirstLine(17, 6, full), " ")
}

func bankState(banks []int) string {
	state := list.Map(banks, func(bank int) string {
		return fmt.Sprintf("%d", bank)
	})
	return strings.Join(state, ",")
}
