package aoc18

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/str"
)

func Day10() Solution {
	particles := data10(true)
	numParticles := len(particles)
	limit := 15_000
	states := make(map[int][]Particle, limit)
	areas := make([]int, limit)
	for t := range limit {
		particles2 := make([]Particle, numParticles)
		for i, p := range particles {
			c, d := p.Coords, p.Delta
			c = Move(c, d)
			particles2[i] = Particle{Coords: c, Delta: d}
		}
		particles = particles2
		states[t] = particles2
		areas[t] = computeArea(particles2)
	}
	// Part 1 and 2
	t := ArgMin(areas)
	displayParticles(states[t])

	return NewSolution("", t+1)
}

func data10(full bool) []Particle {
	return list.Map(ReadLines(18, 10, full), func(line string) Particle {
		p := str.CleanSplit(line, ">")
		head := str.CleanSplit(p[0], "<")[1]
		tail := str.CleanSplit(p[1], "<")[1]
		c := ToInt2(head, ",")
		d := ToInt2(tail, ",")
		return Particle{
			Coords: Coords{c[1], c[0]},
			Delta:  Delta{d[1], d[0]},
		}
	})
}

type Particle struct {
	Coords
	Delta
}

func computeArea(particles []Particle) int {
	c1, c2 := computeBounds(particles)
	h := c2[0] - c1[0]
	w := c2[1] - c1[1]
	return h * w
}

func computeBounds(particles []Particle) (Coords, Coords) {
	ys := list.Map(particles, func(p Particle) int {
		return p.Coords[0]
	})
	xs := list.Map(particles, func(p Particle) int {
		return p.Coords[1]
	})
	x1, y1 := slices.Min(xs), slices.Min(ys)
	x2, y2 := slices.Max(xs), slices.Max(ys)
	return Coords{y1, x1}, Coords{y2, x2}
}

func displayParticles(particles []Particle) {
	c1, c2 := computeBounds(particles)
	y1, x1 := c1.Tuple()
	y2, x2 := c2.Tuple()
	g := make(map[Coords]string)
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			g[Coords{y, x}] = "."
		}
	}
	for _, p := range particles {
		g[p.Coords] = "#"
	}
	for y := y1; y <= y2; y++ {
		line := make([]string, 0)
		for x := x1; x <= x2; x++ {
			line = append(line, g[Coords{y, x}])
		}
		fmt.Println(strings.Join(line, ""))
	}
}
