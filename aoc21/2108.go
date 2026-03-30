package aoc21

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

type Pair = [2][]string
type Candidates = map[int][]string
type Domain = map[string][]string
type Clues = [2]Candidates

func Day08() Solution {
	pairs := data08(true)

	// Part 1
	lengthOptions := []int{2, 3, 4, 7}
	total1 := 0
	for _, pair := range pairs {
		valid := list.Filter(pair[1], func(x string) bool {
			return slices.Contains(lengthOptions, len(x))
		})
		total1 += len(valid)
	}

	// Part 2
	m, clues := digitMapClues()
	total2 := 0
	for _, pair := range pairs {
		total2 += solveDigits(pair, m, clues)
	}

	return NewSolution(total1, total2)
}

func data08(full bool) []Pair {
	return list.Map(ReadLines(21, 8, full), func(line string) Pair {
		p := str.CleanSplit(line, "|")
		return Pair{str.SpaceSplit(p[0]), str.SpaceSplit(p[1])}
	})
}

func digitMapClues() (map[string]string, Clues) {
	digits := []string{
		"abcefg", "cf", "acdeg", "acdfg", "bcdf",
		"abdfg", "abdefg", "acf", "abcdefg", "abcdfg",
	}
	m := make(map[string]string)
	for i, code := range digits {
		m[code] = fmt.Sprintf("%d", i)
	}
	clues := getClues(digits)
	return m, clues
}

func getClues(digits []string) Clues {
	freq := make(map[string]int)
	size := make(map[string]int)
	for _, d := range digits {
		d = SortedString(d)
		size[d] = len(d)
		for _, x := range d {
			k := RuneToString(x)
			freq[k] += 1
		}
	}
	count := make(Candidates)
	length := make(Candidates)
	for k, v := range freq {
		count[v] = append(count[v], k)
	}
	for k, v := range size {
		length[v] = append(length[v], k)
	}
	return Clues{count, length}
}

func solveDigits(pair Pair, m map[string]string, clues Clues) int {
	digits, output := pair[0], pair[1]
	clues2 := getClues(digits)
	t := alignClues(clues, clues2)
	return translateOutput(output, m, t)
}

func alignClues(clues1 Clues, clues2 Clues) map[string]string {
	count1, length1 := clues1[0], clues1[1]
	count2, length2 := clues2[0], clues2[1]
	t := make(map[string]string)
	domain := make(Domain)
	for k, items2 := range count2 {
		items1 := count1[k]
		if len(items1) == 1 && len(items2) == 1 {
			s1, s2 := items1[0], items2[0]
			t[s2] = s1
		} else {
			for _, s := range items2 {
				domain[s] = items1[:]
			}
		}
	}

	for _, k := range []int{2, 3, 4} {
		code := StringChars(length2[k][0])
		choices := ds.SetFrom(StringChars(length1[k][0]))
		unmapped := ds.SetFrom(code)
		for choices.Len() > 1 {
			for _, c := range code {
				if dict.HasKey(t, c) {
					unmapped.Delete(c)
					choices.Delete(t[c])
				}
			}
		}
		if choices.Len() == 1 {
			b := unmapped.Items()[0]
			a := choices.Items()[0]
			t, domain = assign(b, a, t, domain)
		}
	}

	return t
}

func assign(b string, a string, t map[string]string, domain Domain) (map[string]string, Domain) {
	t[b] = a
	if dict.HasKey(domain, b) {
		delete(domain, b)
	}

	sure := make([]Str2, 0)
	domain2 := make(Domain)

	for k, items := range domain {
		if idx := slices.Index(items, a); idx != -1 {
			items = RemoveIndex(items, idx)
		}
		if len(items) == 1 {
			sure = append(sure, Str2{k, items[0]})
		} else {
			domain2[k] = items
		}
	}

	for _, pair := range sure {
		b, a := pair[0], pair[1]
		t, domain2 = assign(b, a, t, domain2)
	}

	return t, domain2
}

func translateOutput(output []string, m map[string]string, t map[string]string) int {
	orig := make([]string, 0)
	for _, code := range output {
		out := list.Translate(list.Map([]rune(code), RuneToString), t)
		slices.Sort(out)
		orig = append(orig, strings.Join(out, ""))
	}

	digit := list.Translate(orig, m)
	d := strings.Join(digit, "")
	return number.ParseInt(d)
}
