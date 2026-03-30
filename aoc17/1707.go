package aoc17

import (
	"sort"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day07() Solution {
	t := data07(true)

	// Part 1
	var root string
	for _, name := range t.nodes {
		if dict.NoKey(t.parentOf, name) {
			root = name
			break
		}
	}

	// Part 2
	var targetWeight int
	weight := make(map[string]int)
	q := t.nodes
mainLoop:
	for len(q) > 0 {
		q2 := make([]string, 0)
		for _, node := range q {
			if dict.NoKey(t.children, node) {
				weight[node] = t.weight[node]
				continue
			}
			computable := list.All(t.children[node], func(child string) bool {
				return dict.HasKey(weight, child)
			})
			if computable {
				childWeights := ds.NewSet[int]()
				totalChild := 0
				for _, child := range t.children[node] {
					w := weight[child]
					childWeights.Add(w)
					totalChild += w
				}
				weight[node] = t.weight[node] + totalChild

				if childWeights.Len() == 2 {
					weights := childWeights.Items()
					sort.Ints(weights)
					target, heavy := weights[0], weights[1]
					heavyChild := ""
					for _, child := range t.children[node] {
						if weight[child] == heavy {
							heavyChild = child
							break
						}
					}
					targetWeight = t.weight[heavyChild] - (heavy - target)
					break mainLoop
				}
			} else {
				q2 = append(q2, node)
			}
		}
		q = q2
	}

	return NewSolution(root, targetWeight)
}

func data07(full bool) Tree {
	t := Tree{
		nodes:    make([]string, 0),
		weight:   make(map[string]int),
		parentOf: make(map[string]string),
		children: make(map[string][]string),
	}
	for _, line := range ReadLines(17, 7, full) {
		p := str.CleanSplit(line, "->")
		node := str.CleanSplit(p[0], "(")
		name := node[0]
		t.nodes = append(t.nodes, name)
		t.weight[name] = number.ParseInt(strings.TrimSuffix(node[1], ")"))
		if len(p) == 1 {
			continue
		}
		if dict.NoKey(t.children, name) {
			t.children[name] = make([]string, 0)
		}
		for _, child := range str.CleanSplit(p[1], ",") {
			t.parentOf[child] = name
			t.children[name] = append(t.children[name], child)
		}
	}
	return t
}

type Tree struct {
	nodes    []string
	weight   map[string]int
	parentOf map[string]string
	children map[string][]string
}
