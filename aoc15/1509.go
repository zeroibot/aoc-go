package aoc15

import (
	"math"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day09() Solution {
	g := data09(true)
	minDistance := math.MaxInt
	maxDistance := 0
	vertices := g.Vertices
	for _, path := range Permutations(vertices, len(vertices)) {
		distance := computeDistance(path, g)
		// Part 1
		minDistance = min(minDistance, distance)
		// Part 2
		maxDistance = max(maxDistance, distance)
	}
	return NewSolution(minDistance, maxDistance)
}

func data09(full bool) *ds.Graph {
	g := ds.NewGraph()
	vertices := ds.NewSet[ds.Vertex]()
	for _, line := range ReadLines(15, 9, full) {
		p := str.CleanSplit(line, "=")
		v := str.CleanSplit(p[0], "to")
		w := number.ParseFloat(p[1])
		v1, v2 := v[0], v[1]
		vertices.Add(v1)
		vertices.Add(v2)
		g.EdgeWeightOf[ds.Edge{v1, v2}] = w
		g.EdgeWeightOf[ds.Edge{v2, v1}] = w
	}
	g.Vertices = vertices.Items()
	return g
}

func computeDistance(path []string, g *ds.Graph) int {
	total := 0
	for i := 1; i < len(path); i++ {
		pair := ds.Edge{path[i-1], path[i]}
		if weight, ok := g.EdgeWeight(pair); ok {
			total += int(weight)
		}
	}
	return total
}
