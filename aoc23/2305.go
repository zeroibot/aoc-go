package aoc23

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

func Day05() Solution {
	seeds, maps := data05(true)

	// Part 1
	locations1 := applyMapChain(seeds, maps)
	minLoc1 := slices.Min(locations1)

	// Part 2
	locations2 := applyMapRangeChain(seeds, maps)
	minLoc2 := slices.Min(locations2)

	return NewSolution(minLoc1, minLoc2)
}

func data05(full bool) ([]int, map[string][]Int3) {
	seeds := make([]int, 0)
	maps := make(map[string][]Int3)
	var key string
	for _, line := range ReadLines(23, 5, full) {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "seeds:") {
			tail := str.CleanSplit(line, ":")[1]
			seeds = ToIntList(tail, " ")
		} else if strings.HasSuffix(line, "map:") {
			key = str.SpaceSplit(line)[0]
			maps[key] = make([]Int3, 0)
		} else {
			p := ToIntList(line, " ")
			dst, src, count := p[0], p[1], p[2]
			maps[key] = append(maps[key], Int3{src, dst, count})
		}
	}
	return seeds, maps
}

var seedChain = []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func applyMapChain(seeds []int, maps map[string][]Int3) []int {
	current := seeds
	for i := range len(seedChain) - 1 {
		key := fmt.Sprintf("%s-to-%s", seedChain[i], seedChain[i+1])
		current = plantTranslate(current, maps[key])
	}
	return current
}

func plantTranslate(numbers []int, triples []Int3) []int {
	result := make([]int, 0)
	for _, x := range numbers {
		y := x
		for _, t := range triples {
			src, dst, count := t.Tuple()
			if src <= x && x < src+count {
				y = dst + (x - src)
				break
			}
		}
		result = append(result, y)
	}
	return result
}

func seedRanges(seeds []int) []Int2 {
	return list.Map(NumRangeInc(0, len(seeds), 2), func(i int) Int2 {
		return Int2{seeds[i], seeds[i] + seeds[i+1] - 1}
	})
}

func mapRanges(maps map[string][]Int3) map[string][]Int3 {
	m := make(map[string][]Int3)
	for key, triples := range maps {
		m[key] = make([]Int3, 0)
		for _, t := range triples {
			src, dst, count := t.Tuple()
			m[key] = append(m[key], Int3{src, src + count - 1, dst - src})
		}
	}
	return m
}

func applyMapRangeChain(seeds []int, maps map[string][]Int3) []int {
	currRanges := ds.QueueFrom(seedRanges(seeds))
	rangeMap := mapRanges(maps)
	for i := range len(seedChain) - 1 {
		key := fmt.Sprintf("%s-to-%s", seedChain[i], seedChain[i+1])
		nextRanges := ds.NewQueue[Int2]()
		for currRanges.Len() > 0 {
			currRange, _ := currRanges.Dequeue()
			found := false
			for _, t := range rangeMap[key] {
				start, end, diff := t.Tuple()
				ruleRange := Int2{start, end}
				if isInside(ruleRange, currRange) {
					first, last := currRange.Tuple()
					nextRanges.Enqueue(Int2{first + diff, last + diff})
					found = true
					break
				}
				match, extra := findIntersection(ruleRange, currRange)
				if match != nil && extra != nil {
					first, last := match.Tuple()
					nextRanges.Enqueue(Int2{first + diff, last + diff})
					currRanges.Enqueue(*extra)
					found = true
					break
				}
			}
			if !found {
				nextRanges.Enqueue(currRange)
			}
		}
		currRanges = nextRanges
	}
	result := make([]int, 0)
	for currRanges.Len() > 0 {
		front, _ := currRanges.Dequeue()
		result = append(result, front[0])
	}
	return result
}

func isInside(ruleRange, currRange Int2) bool {
	minValue, maxValue := ruleRange.Tuple()
	start, end := currRange.Tuple()
	return minValue <= start && start <= end && end <= maxValue
}

func findIntersection(ruleRange, currRange Int2) (*Int2, *Int2) {
	minValue, maxValue := ruleRange.Tuple()
	start, end := currRange.Tuple()
	if minValue <= start && start <= maxValue {
		return &Int2{start, maxValue}, &Int2{maxValue + 1, end}
	} else if minValue <= end && end <= maxValue {
		return &Int2{minValue, end}, &Int2{start, minValue - 1}
	} else {
		return nil, nil
	}
}
