package aoc23

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/list"
)

func Day10() Solution {
	grid := data10(true)

	// Part 1
	maxDistance := bfsMaxDistance(grid)

	// Part 2
	// Idea: Expand the grid to twice its size, so we can perform floodfill
	// After getting the reachable points from the floodfill, filter out only those that have
	// even x and y coords (the original coords)
	points := dfsVisit(grid)
	grid2 := createExpandedGrid(points, grid.bounds)
	visited := floodFillMaze(grid2)
	for _, pt := range visited {
		y, x := pt.Tuple()
		if y%2 == 0 && x%2 == 0 {
			grid2[pt] = 'X'
		}
	}
	inside := list.Filter(dict.Values(grid2), func(tile rune) bool {
		return tile == 'X'
	})
	count := len(inside)

	return NewSolution(maxDistance, count)
}

func data10(full bool) *MazeGrid {
	lines := ReadLines(23, 10, full)
	grid := &MazeGrid{
		start:     Coords{0, 0},
		bounds:    Dims2{0, 0},
		edges:     make(map[Coords][]Coords),
		edgeCount: make(map[CoordPair]int),
	}
	grid.bounds = StringGridBounds(lines)
	for row, line := range lines {
		for col, char := range line {
			if char == '.' {
				continue
			}
			c1 := Coords{row, col}
			dirs := []Delta{}
			switch char {
			case '|':
				dirs = []Delta{U, D}
			case '-':
				dirs = []Delta{L, R}
			case 'L':
				dirs = []Delta{U, R}
			case 'J':
				dirs = []Delta{U, L}
			case '7':
				dirs = []Delta{L, D}
			case 'F':
				dirs = []Delta{R, D}
			case 'S':
				grid.start = c1
				dirs = []Delta{U, D, L, R}
			}
			for _, d := range dirs {
				c2 := Move(c1, d)
				if !InsideBounds(c2, grid.bounds) {
					continue
				}
				grid.addEdge(c1, c2)
			}
		}
	}
	grid.createEdges()
	return grid
}

type CoordPair = [2]Coords
type CoordDist struct {
	Coords
	Dist int
}

func (cd CoordDist) Tuple() (Coords, int) {
	return cd.Coords, cd.Dist
}

type MazeGrid struct {
	start     Coords
	bounds    Dims2
	edges     map[Coords][]Coords
	edgeCount map[CoordPair]int
}

func (g *MazeGrid) addEdge(node1, node2 Coords) {
	nodes := []Coords{node1, node2}
	slices.SortFunc(nodes, SortCoords)
	pair := CoordPair{nodes[0], nodes[1]}
	g.edgeCount[pair] += 1
}

func (g *MazeGrid) createEdges() {
	edges := make(map[Coords]*ds.Set[Coords])
	for pair, count := range g.edgeCount {
		if count != 2 {
			continue
		}
		node1, node2 := pair[0], pair[1]
		if _, ok := edges[node1]; !ok {
			edges[node1] = ds.NewSet[Coords]()
		}
		if _, ok := edges[node2]; !ok {
			edges[node2] = ds.NewSet[Coords]()
		}
		edges[node1].Add(node2)
		edges[node2].Add(node1)
	}
	for node, neighbors := range edges {
		nextCoords := neighbors.Items()
		slices.SortFunc(nextCoords, SortCoords)
		g.edges[node] = nextCoords
	}
}

func bfsMaxDistance(grid *MazeGrid) int {
	dist := make(map[Coords]int)
	visited := ds.NewSet[Coords]()
	q := ds.NewQueue[CoordDist]()
	q.Enqueue(CoordDist{grid.start, 0})
	for q.Len() > 0 {
		pair, _ := q.Dequeue()
		node, d := pair.Tuple()
		dist[node] = d
		visited.Add(node)
		for _, node2 := range grid.edges[node] {
			if visited.Has(node2) {
				continue
			}
			q.Enqueue(CoordDist{node2, d + 1})
		}
	}
	return slices.Max(dict.Values(dist))
}

func dfsVisit(grid *MazeGrid) []Coords {
	points := make([]Coords, 0)
	stack := ds.NewStack[Coords]()
	stack.Push(grid.start)
	for stack.Len() > 0 {
		node, _ := stack.Pop()
		if slices.Contains(points, node) {
			continue
		}
		points = append(points, node)
		for _, node2 := range grid.edges[node] {
			if slices.Contains(points, node2) {
				continue
			}
			stack.Push(node2)
		}
	}
	return points
}

func createExpandedGrid(points []Coords, bounds Dims2) map[Coords]rune {
	grid := make(map[Coords]rune)
	rows, cols := bounds.Tuple()
	for row := range rows * 2 {
		for col := range cols * 2 {
			grid[Coords{row, col}] = '.'
		}
	}
	points = append(points, points[0])
	yprev, xprev := points[0].Tuple()
	pt := Coords{yprev * 2, xprev * 2}
	grid[pt] = '#'
	for _, pt := range points[1:] {
		y, x := pt.Tuple()
		if yprev == y {
			xadd := lang.Ternary(x > xprev, 1, -1)
			pt = Coords{yprev * 2, (xprev * 2) + xadd}
			grid[pt] = '#'
		} else if xprev == x {
			yadd := lang.Ternary(y > yprev, 1, -1)
			pt = Coords{(yprev * 2) + yadd, xprev * 2}
			grid[pt] = '#'
		}
		pt = Coords{y * 2, x * 2}
		grid[pt] = '#'
		yprev, xprev = y, x
	}
	return grid
}

func floodFillMaze(grid map[Coords]rune) []Coords {
	visited := ds.NewSet[Coords]()
	pts := make(map[int][]int)
	for pt, tile := range grid {
		if tile == '#' {
			y, x := pt.Tuple()
			pts[y] = append(pts[y], x)
		}
	}
	ymin := slices.Min(dict.Keys(pts))
	xmin := slices.Min(pts[ymin])
	start := Coords{ymin + 1, xmin + 1}
	stack := ds.NewStack[Coords]()
	stack.Push(start)
	for stack.Len() > 0 {
		curr, _ := stack.Pop()
		if visited.Has(curr) {
			continue
		}
		visited.Add(curr)
		for _, nxt := range Surround4(curr) {
			if visited.Has(nxt) {
				continue
			}
			if grid[nxt] == '.' {
				stack.Push(nxt)
			}
		}
	}
	return visited.Items()
}
