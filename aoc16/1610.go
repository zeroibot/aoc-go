package aoc16

import (
	"sort"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day10() Solution {
	chips, bots := data10(true)
	goal := Int2{17, 61}
	output := make(map[int][]int)
	botValues := createBotValues(bots)

	for _, p := range chips {
		value, who := p.Tuple()
		botValues[who] = append(botValues[who], value)
	}

	var goalBot int = 0
	for {
		botValues2 := createBotValues(bots)
		hasMovement := false
		for b, v := range botValues {
			if len(v) == 2 {
				sort.Ints(v)
				// Part 1
				if goalBot == 0 && v[0] == goal[0] && v[1] == goal[1] {
					goalBot = b
				}
				for i, p := range bots[b] {
					dest, who := p.Tuple()
					if dest == "bot" {
						botValues2[who] = append(botValues2[who], v[i])
					} else {
						output[who] = append(output[who], v[i])
					}
				}
				hasMovement = true
			} else {
				botValues2[b] = append(botValues2[b], v...)
			}
		}
		botValues = botValues2
		if !hasMovement {
			break
		}
	}
	// Part 2
	a, b, c := output[0][0], output[1][0], output[2][0]
	product := a * b * c

	return NewSolution(goalBot, product)
}

func data10(full bool) ([]Int2, map[int]Bot) {
	chips := make([]Int2, 0)
	bots := make(map[int]Bot)
	for _, line := range ReadLines(16, 10, full) {
		p := str.SpaceSplit(line)
		N := len(p)
		switch p[0] {
		case "value":
			value, who := number.ParseInt(p[1]), number.ParseInt(p[N-1])
			chips = append(chips, Int2{value, who})
		case "bot":
			who := number.ParseInt(p[1])
			low := StrInt{
				Str: p[5],
				Int: number.ParseInt(p[6]),
			}
			high := StrInt{
				Str: p[N-2],
				Int: number.ParseInt(p[N-1]),
			}
			bots[who] = [2]StrInt{low, high}
		}
	}
	return chips, bots
}

type Bot = [2]StrInt

func createBotValues(bots map[int]Bot) map[int][]int {
	botValues := make(map[int][]int)
	for b := range bots {
		botValues[b] = make([]int, 0)
	}
	return botValues
}
