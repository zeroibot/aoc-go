package aoc20

import (
	"slices"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
)

func Day05() Solution {
	seats := data05(true)
	numSeats := len(seats)

	maxID := 0
	ids := make([]int, numSeats)
	for i, seat := range seats {
		seatID := computeID(seat)
		ids[i] = seatID
		// Part 1
		maxID = max(maxID, seatID)
	}

	// Part 2
	slices.Sort(ids)
	seatID := 0
	for i := 1; i < numSeats; i++ {
		if ids[i]-ids[i-1] > 1 {
			seatID = ids[i] - 1
			break
		}
	}

	return NewSolution(maxID, seatID)
}

func data05(full bool) []Seat {
	T := map[rune]int{
		'F': 0,
		'B': 1,
		'L': 0,
		'R': 1,
	}
	return list.Map(ReadLines(20, 5, full), func(line string) Seat {
		row := list.Translate([]rune(line[:7]), T)
		col := list.Translate([]rune(line[7:]), T)
		return Seat{row, col}
	})
}

type Seat [2][]int

func computeID(seat Seat) int {
	numRows, numCols := 128, 8
	rows, cols := seat[0], seat[1]
	row := binaryMoves(rows, numRows)
	col := binaryMoves(cols, numCols)
	return (row * 8) + col
}

func binaryMoves(sides []int, limit int) int {
	start, end := 0, limit
	last := len(sides) - 1
	for _, s := range sides[:last] {
		mid := start + ((end - start) / 2)
		switch s {
		case 0:
			end = mid
		case 1:
			start = mid
		}
	}
	if sides[last] == 0 {
		return start
	} else {
		return end - 1
	}
}
