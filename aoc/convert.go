package aoc

import (
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func ToIntList(line string, sep string) []int {
	parts := splitParts(line, sep)
	return list.Map(parts, number.ParseInt)
}

func ToIntLine(line string) []int {
	digits := make([]string, 0)
	for i := range len(line) {
		digits = append(digits, string(line[i:i+1]))
	}
	return list.Map(digits, number.ParseInt)
}

func ToDims3(line string, sep string) Dims3 {
	p := ToIntList(line, sep)
	return Dims3{p[0], p[1], p[2]}
}

func ToInt2(line string, sep string) Int2 {
	p := ToIntList(line, sep)
	return Int2{p[0], p[1]}
}

func ToInt3(line string, sep string) Int3 {
	p := ToIntList(line, sep)
	return Int3{p[0], p[1], p[2]}
}

func ToCharInt(line string) CharInt {
	chars := []rune(line)
	return CharInt{
		Char: chars[0],
		Int:  number.ParseInt(string(chars[1:])),
	}
}

func ToStrInt(line string, sep string) StrInt {
	parts := splitParts(line, sep)
	return StrInt{Str: parts[0], Int: number.ParseInt(parts[1])}
}

func ToStr2(line string, sep string) Str2 {
	p := splitParts(line, sep)
	return Str2{p[0], p[1]}
}

func splitParts(line string, sep string) []string {
	var parts []string
	if sep == " " {
		parts = str.SpaceSplit(line)
	} else {
		parts = str.CleanSplit(line, sep)
	}
	return parts
}
