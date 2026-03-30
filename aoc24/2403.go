package aoc24

import (
	"regexp"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day03() Solution {
	text := data03(true)
	mulPattern := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	commands := make([]StrInt, 0)

	// Part 1
	total1 := 0
	for _, m := range mulPattern.FindAllIndex([]byte(text), -1) {
		start, end := m[0], m[1]
		cmd := text[start:end]
		total1 += execCommand(cmd)
		commands = append(commands, StrInt{Str: cmd, Int: start})
	}

	// Part 2
	offPattern := regexp.MustCompile(`don't\(\)`)
	onPattern := regexp.MustCompile(`do\(\)`)
	regions := []Int2{{0, 1}} // idx, 0(F)/1(T)
	for flag, pattern := range []*regexp.Regexp{offPattern, onPattern} {
		for _, m := range pattern.FindAllIndex([]byte(text), -1) {
			regions = append(regions, Int2{m[0], flag})
		}
	}
	slices.SortFunc(regions, SortInt2)

	ignore := make([]Int2, 0)
	var offStart *int = nil
	for _, region := range regions {
		start, flag := region.Tuple()
		if flag == 0 && offStart == nil {
			offStart = &start
		} else if flag == 1 && offStart != nil {
			ignore = append(ignore, Int2{*offStart, start - 1})
			offStart = nil
		}
	}
	if offStart != nil {
		ignore = append(ignore, Int2{*offStart, len(text) - 1})
	}

	total2 := 0
	for _, cmd := range commands {
		invalid := list.Any(ignore, func(x Int2) bool {
			return x[0] <= cmd.Int && cmd.Int <= x[1]
		})
		if !invalid {
			total2 += execCommand(cmd.Str)
		}
	}

	return NewSolution(total1, total2)
}

func data03(full bool) string {
	return strings.Join(ReadLines(24, 3, full), "")
}

func execCommand(cmd string) int {
	cmd = strings.Trim(cmd, "mul()")
	p := ToIntList(cmd, ",")
	return p[0] * p[1]
}
