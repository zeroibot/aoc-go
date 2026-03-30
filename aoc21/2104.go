package aoc21

import (
	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
)

func Day04() Solution {
	// Part 1
	numbers, cards := data04(true)
	score1 := playBingo(numbers, cards, 1)

	// Part 2
	numbers, cards = data04(true)
	score2 := playBingo(numbers, cards, len(cards))

	return NewSolution(score1, score2)
}

func data04(full bool) ([]int, []*Bingo) {
	lines := ReadRawLines(21, 4, full, true)
	numbers := ToIntList(lines[0], ",")
	cards := make([]*Bingo, 0)
	card := make(IntGrid, 0)
	for _, line := range lines[2:] {
		if line == "" {
			cards = append(cards, newBingo(card))
			card = make(IntGrid, 0)
		} else {
			card = append(card, ToIntList(line, " "))
		}
	}
	cards = append(cards, newBingo(card))
	return numbers, cards
}

type Bingo struct {
	card   IntGrid
	rows   int
	cols   int
	marked map[Coords]bool
	lookup map[int]Coords
}

func (b *Bingo) Mark(number int) {
	if dict.HasKey(b.lookup, number) {
		pt := b.lookup[number]
		b.marked[pt] = true
	}
}

func (b *Bingo) HasWon() bool {
	for row := range b.rows {
		bingo := list.All(NumRange(0, b.cols), func(col int) bool {
			return b.marked[Coords{row, col}]
		})
		if bingo {
			return true
		}
	}
	for col := range b.cols {
		bingo := list.All(NumRange(0, b.rows), func(row int) bool {
			return b.marked[Coords{row, col}]
		})
		if bingo {
			return true
		}
	}
	return false
}

func (b *Bingo) Score() int {
	unmarked := list.Filter(dict.Keys(b.marked), func(c Coords) bool {
		return !b.marked[c]
	})
	return list.Sum(list.Map(unmarked, func(c Coords) int {
		row, col := c.Tuple()
		return b.card[row][col]
	}))
}

func newBingo(card IntGrid) *Bingo {
	bounds := GridBounds(card)
	b := &Bingo{
		card:   card,
		rows:   bounds[0],
		cols:   bounds[1],
		marked: make(map[Coords]bool),
		lookup: make(map[int]Coords),
	}
	for row, line := range card {
		for col, number := range line {
			pt := Coords{row, col}
			b.lookup[number] = pt
			b.marked[pt] = false
		}
	}
	return b
}

func playBingo(numbers []int, cards []*Bingo, targetWinnerCount int) int {
	winners := ds.NewSet[int]()
	score := 0
mainLoop:
	for _, number := range numbers {
		for player, card := range cards {
			if winners.Has(player) {
				continue
			}

			card.Mark(number)
			if card.HasWon() {
				winners.Add(player)
			}
			if winners.Len() == targetWinnerCount {
				score = number * card.Score()
				break mainLoop
			}
		}
	}
	return score
}
