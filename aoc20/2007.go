package aoc20

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

type Hierarchy = map[string][]StrInt

func Day07() Solution {
	h := data07(true)

	// Part 1
	parents := make(map[string][]string)
	for parentColor, bagCounts := range h {
		for _, bag := range bagCounts {
			color := bag.Str
			parents[color] = append(parents[color], parentColor)
		}
	}
	valid := ds.NewSet[string]()
	q := ds.QueueFrom(parents["shiny gold"])
	for q.Len() > 0 {
		color, _ := q.Dequeue()
		valid.Add(color)
		for _, nxtColor := range parents[color] {
			q.Enqueue(nxtColor)
		}
	}
	count1 := valid.Len()

	// Part 2
	count2 := countInside("shiny gold", h)

	return NewSolution(count1, count2)
}

func data07(full bool) Hierarchy {
	h := make(Hierarchy)
	for _, line := range ReadLines(20, 7, full) {
		p := str.CleanSplit(line, "contain")
		head, tail := p[0], p[1]
		if tail == "no other bags." {
			continue
		}
		p = str.SpaceSplit(head)
		p = p[:len(p)-1] // remove 'bags'
		color := strings.Join(p, " ")
		bags := str.CleanSplit(tail, ",")
		h[color] = list.Map(bags, bagCount)
	}
	return h
}

func bagCount(text string) StrInt {
	p := str.SpaceSplit(text)
	color := strings.Join(p[1:len(p)-1], " ")
	count := number.ParseInt(p[0])
	return StrInt{Str: color, Int: count}
}

func countInside(color string, h Hierarchy) int {
	total := 0
	if dict.HasKey(h, color) {
		for _, bag := range h[color] {
			color2, count := bag.Tuple()
			total += count + (count * countInside(color2, h))
		}
	}
	return total
}
