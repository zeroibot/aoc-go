package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zeroibot/fn/io"
	"github.com/zeroibot/fn/lang"
	"github.com/zeroibot/fn/str"
)

func dataPath(year int, day int, full bool) string {
	folder := "test"
	if full {
		folder = fmt.Sprintf("20%d", year)
	}
	path := fmt.Sprintf("%s/%s/%d%.02d.txt", rootDir(), folder, year, day)
	return path

}

func ReadRawLines(year int, day int, full bool, strip bool) []string {
	path := dataPath(year, day, full)
	text, err := io.ReadFile(path)
	if err != nil {
		log.Fatal("Error:", err)
	}
	lines := lang.Ternary(strip, str.CleanSplit(text, "\n"), strings.Split(text, "\n"))
	return lines
}

func ReadLines(year int, day int, full bool) []string {
	path := dataPath(year, day, full)
	lines, err := io.ReadNonEmptyLines(path)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return lines
}

func ReadFirstLine(year int, day int, full bool) string {
	return ReadLines(year, day, full)[0]
}

func GetSolution(year int, day int) Solution {
	path := fmt.Sprintf("%s/solutions/all.csv", rootDir())
	solutions := make(map[string]Solution)
	lines, err := io.ReadLines(path)
	if err != nil {
		log.Fatal("Error:", err)
	}
	for _, line := range lines {
		p := str.CleanSplit(line, "|")
		k := fmt.Sprintf("%s%s", p[0], p[1])
		v := Solution{p[2], p[3]}
		solutions[k] = v
	}
	key := fmt.Sprintf("%d%.02d", year, day)
	return solutions[key]
}

func rootDir() string {
	root := os.Getenv("AOC_DATA_DIR")
	if root == "" {
		root = "../aoc-data"
	}
	return root
}
