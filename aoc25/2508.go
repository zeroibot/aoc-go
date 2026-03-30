package aoc25

import (
	"cmp"
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
)

type pointPair struct {
	pt1, pt2 Int3
	distance float64
}

func Day08() Solution {
	points := data08(true)
	pairLimit := 1000
	numPoints := len(points)

	pairs := make([]pointPair, 0)
	for i := range numPoints {
		p1 := points[i]
		for j := i + 1; j < numPoints; j++ {
			p2 := points[j]
			pairs = append(pairs, pointPair{p1, p2, Euclidean3(p1, p2)})
		}
	}
	slices.SortFunc(pairs, func(p1, p2 pointPair) int {
		return cmp.Compare(p1.distance, p2.distance)
	})

	groupOf := make(map[Int3]int)
	inGroup := make(map[int][]Int3)
	group, size, length := 0, 0, 0
	allGrouped := false
	for i, pair := range pairs {
		p1, p2 := pair.pt1, pair.pt2
		hasGroup1, hasGroup2 := dict.HasKey(groupOf, p1), dict.HasKey(groupOf, p2)
		if hasGroup1 && hasGroup2 {
			group1, group2 := groupOf[p1], groupOf[p2]
			if group1 != group2 {
				for _, pt := range inGroup[group2] {
					groupOf[pt] = group1
				}
				inGroup[group1] = append(inGroup[group1], inGroup[group2]...)
				delete(inGroup, group2)
			}
		} else if hasGroup1 {
			group1 := groupOf[p1]
			groupOf[p2] = group1
			inGroup[group1] = append(inGroup[group1], p2)
		} else if hasGroup2 {
			group2 := groupOf[p2]
			groupOf[p1] = group2
			inGroup[group2] = append(inGroup[group2], p1)
		} else {
			groupOf[p1] = group
			groupOf[p2] = group
			inGroup[group] = []Int3{p1, p2}
			group += 1
		}

		if i == pairLimit-1 {
			sizes := list.Map(dict.Values(inGroup), list.Length)
			slices.SortFunc(sizes, func(a, b int) int {
				return cmp.Compare(b, a)
			})
			size = sizes[0] * sizes[1] * sizes[2]
		}

		if !allGrouped {
			allGrouped = len(groupOf) == numPoints
		}
		if allGrouped && list.AllSame(dict.Values(groupOf)) {
			length = p1[0] * p2[0]
			break
		}
	}

	return NewSolution(size, length)
}

func data08(full bool) []Int3 {
	return list.Map(ReadLines(25, 8, full), func(line string) Int3 {
		return ToInt3(line, ",")
	})
}
