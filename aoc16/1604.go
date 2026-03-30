package aoc16

import (
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day04() Solution {
	rooms := data04(true)
	total := 0
	roomID := 0
	for _, room := range rooms {
		// Part 1
		if room.isReal() {
			total += room.id
		}

		// Part 2
		if roomID == 0 && room.decrypt() == "northpole-object-storage" {
			roomID = room.id
		}
	}
	return NewSolution(total, roomID)
}

func data04(full bool) []Room {
	return list.Map(ReadLines(16, 4, full), func(line string) Room {
		p := str.CleanSplit(line, "[")
		head, tail := p[0], p[1]
		h := str.CleanSplit(head, "-")
		lastIdx := len(h) - 1
		return Room{
			checksum: strings.TrimSuffix(tail, "]"),
			name:     strings.Join(h[:lastIdx], "-"),
			id:       number.ParseInt(h[lastIdx]),
		}
	})
}

type Room struct {
	checksum string
	name     string
	id       int
}

func (r Room) isReal() bool {
	freq := CharFreq(r.name, []rune{'-'})
	if len(freq) < 5 {
		return false
	}
	pairs := make([]CharInt, 0)
	for k, v := range freq {
		pairs = append(pairs, CharInt{Char: k, Int: v})
	}
	slices.SortFunc(pairs, SortCharIntDesc)
	top5 := list.Map(pairs[:5], func(pair CharInt) rune {
		return pair.Char
	})
	return string(top5) == r.checksum
}

func (r Room) decrypt() string {
	msg := []rune(r.name)
	for range r.id {
		for i, char := range msg {
			msg[i] = rotateLetter(char)
		}
	}
	return string(msg)
}

func rotateLetter(letter rune) rune {
	switch letter {
	case '-':
		return '-'
	case 'z':
		return 'a'
	default:
		return letter + 1
	}
}
