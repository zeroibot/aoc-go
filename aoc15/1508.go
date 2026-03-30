package aoc15

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
)

func Day08() Solution {
	words := data08(true)

	// Part 1
	total1 := 0
	for _, word := range words {
		total1 += len(word) - len(parseString(word))
	}

	// Part 2
	total2 := 0
	for _, word := range words {
		total2 += len(expandString(word)) - len(word)
	}

	return NewSolution(total1, total2)
}

func data08(full bool) []string {
	return ReadLines(15, 8, full)
}

func parseString(text string) string {
	hex := regexp.MustCompile(`\\x[0-9a-f]{2}`)
	text = text[1 : len(text)-1]
	text = hex.ReplaceAllString(text, ".")
	text = strings.ReplaceAll(text, `\"`, `"`)
	text = strings.ReplaceAll(text, `\\`, ".")
	return text
}

func expandString(text string) string {
	chars := make([]string, 0)
	for _, x := range text {
		switch x {
		case '"':
			chars = append(chars, `\"`)
		case '\\':
			chars = append(chars, `\\`)
		default:
			chars = append(chars, string(x))
		}
	}
	return fmt.Sprintf(`"%s"`, strings.Join(chars, ""))
}
