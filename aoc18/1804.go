package aoc18

import (
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day04() Solution {
	logs := data04(true)
	sleep := processLogs(logs)

	// Part 1
	entries := make([]Int2, 0, len(sleep))
	for guard, s := range sleep {
		total := list.Sum(dict.Values(s))
		entries = append(entries, Int2{total, guard})
	}
	guard := slices.MaxFunc(entries, SortInt2)[1]
	minute := maxSleepEntry(sleep[guard])[1]
	part1 := guard * minute

	// Part 2
	guards := make([]Int3, 0)
	for guard := range sleep {
		maxEntry := maxSleepEntry(sleep[guard])
		count, minute := maxEntry.Tuple()
		guards = append(guards, Int3{count, minute, guard})
	}
	maxEntry := slices.MaxFunc(guards, SortInt3)
	_, minute, guard = maxEntry.Tuple()
	part2 := guard * minute

	return NewSolution(part1, part2)
}

const (
	OFF   int = 0
	ON    int = 1
	GUARD int = 2
)

func data04(full bool) []Int2 {
	lines := ReadLines(18, 4, full)
	slices.Sort(lines)
	return list.Map(lines, func(line string) Int2 {
		p := str.CleanSplit(line, "]")
		head, tail := p[0], p[1]
		minute := number.ParseInt(str.CleanSplit(head, ":")[1])
		if strings.Contains(tail, "asleep") {
			return Int2{ON, minute}
		} else if strings.Contains(tail, "wakes") {
			return Int2{OFF, minute}
		} else {
			id := strings.TrimPrefix(str.SpaceSplit(tail)[1], "#")
			return Int2{GUARD, number.ParseInt(id)}
		}
	})
}

func processLogs(logs []Int2) map[int]map[int]int {
	limit := len(logs)
	i, guard := 0, 0
	sleep := make(map[int]map[int]int)
	for i < limit {
		cmd, x := logs[i][0], logs[i][1]
		switch cmd {
		case GUARD:
			guard = x
			i += 1
		case ON:
			end := logs[i+1][1]
			for m := x; m < end; m++ {
				if dict.NoKey(sleep, guard) {
					sleep[guard] = make(map[int]int)
				}
				sleep[guard][m] += 1
			}
			i += 2
		}
	}
	return sleep
}

func maxSleepEntry(sleep map[int]int) Int2 {
	entries := make([]Int2, 0, len(sleep))
	for minute, count := range sleep {
		entries = append(entries, Int2{count, minute})
	}
	return slices.MaxFunc(entries, SortInt2)
}
