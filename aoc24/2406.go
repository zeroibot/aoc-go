package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
)

func Day06() Solution {
	cfg := data06(true)

	// Part 1
	visited, _ := findExit(cfg, cfg.start, cfg.dir, nil, nil)
	numVisited := len(visited)

	// Part 2
	loopPoints := countLoopPoints(cfg)

	return NewSolution(numVisited, loopPoints)
}

func data06(full bool) *Config {
	lines := ReadLines(24, 6, full)
	cfg := &Config{
		bounds:  StringGridBounds(lines),
		blocked: ds.NewSet[Coords](),
	}
	for row, line := range lines {
		for col, char := range line {
			c := Coords{row, col}
			switch char {
			case '#':
				cfg.blocked.Add(c)
			case '^':
				cfg.start = c
				cfg.dir = U
			}
		}
	}
	return cfg
}

type Config struct {
	bounds  Dims2
	blocked *ds.Set[Coords]
	start   Coords
	dir     Delta
}

type History = map[Coords]*ds.Set[Delta]

func findExit(cfg *Config, start Coords, d Delta, previsit History, obstacle *Coords) (History, bool) {
	c := start
	visited := make(History)
	for k, v := range previsit {
		visited[k] = ds.SetFrom(v.Items())
	}
	dict.SetDefault(visited, c, ds.NewSet[Delta]())
	visited[c].Add(d)
	stuckInLoop := false
	for {
		nxt := Move(c, d)

		if !InsideBounds(nxt, cfg.bounds) {
			break
		} else if cfg.blocked.Has(nxt) || (obstacle != nil && nxt == *obstacle) {
			d = RightOf[d]
		} else {
			c = nxt
			if obstacle != nil && dict.HasKey(visited, c) && visited[c].Has(d) {
				stuckInLoop = true
				break
			}
			dict.SetDefault(visited, c, ds.NewSet[Delta]())
			visited[c].Add(d)
		}
	}
	return visited, stuckInLoop
}

func countLoopPoints(cfg *Config) int {
	c, d := cfg.start, cfg.dir
	obstacles := ds.NewSet[Coords]()
	visited := make(History)
	dict.SetDefault(visited, c, ds.NewSet[Delta]())
	visited[c].Add(d)
	for {
		nxt := Move(c, d)

		if !InsideBounds(nxt, cfg.bounds) {
			break
		} else if cfg.blocked.Has(nxt) {
			d = RightOf[d]
		} else {
			if dict.NoKey(visited, nxt) {
				previsit := make(History)
				for k, v := range visited {
					previsit[k] = ds.SetFrom(v.Items())
				}
				_, stuckInLoop := findExit(cfg, c, RightOf[d], previsit, &nxt)
				if stuckInLoop {
					obstacles.Add(nxt)
				}
			}

			c = nxt
			dict.SetDefault(visited, c, ds.NewSet[Delta]())
			visited[c].Add(d)
		}
	}
	return obstacles.Len()
}
