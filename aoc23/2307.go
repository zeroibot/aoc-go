package aoc23

import (
	"slices"
	"strconv"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
)

func Day07() Solution {
	hands := data07(true)
	power := computeCardPower()

	// Part 1
	slices.SortFunc(hands, compareHandsFn(power))
	total1 := 0
	for i, hand := range hands {
		total1 += (i + 1) * hand.Int
	}

	// Part 2
	power['J'] = 1
	slices.SortFunc(hands, compareHands2Fn(power))
	total2 := 0
	for i, hand := range hands {
		total2 += (i + 1) * hand.Int
	}

	return NewSolution(total1, total2)
}

func data07(full bool) []StrInt {
	return list.Map(ReadLines(23, 7, full), func(line string) StrInt {
		return ToStrInt(line, " ")
	})
}

func computeCardPower() map[byte]int {
	power := make(map[byte]int)
	value := 14
	cards := "AKQJT98765432"
	for i := range len(cards) {
		power[cards[i]] = value
		value -= 1
	}
	return power
}

func compareCards(cards1 string, cards2 string, power map[byte]int) int {
	for i := range len(cards1) {
		card1, card2 := cards1[i], cards2[i]
		if card1 != card2 {
			return power[card1] - power[card2]
		}
	}
	return 0
}

func groupHand(cards string) map[rune]int {
	group := make(map[rune]int)
	for _, card := range cards {
		group[card] += 1
	}
	return group
}

func categoryScore(category string) int {
	switch category {
	case "1,1,1,1,1":
		return 1
	case "2,1,1,1":
		return 2
	case "2,2,1":
		return 3
	case "3,1,1":
		return 4
	case "3,2":
		return 5
	case "4,1":
		return 6
	case "5":
		return 7
	default:
		return 0
	}
}

func computeScore(cards string) int {
	group := groupHand(cards)
	counts := dict.Values(group)
	slices.SortFunc(counts, SortIntDesc)
	category := strings.Join(list.Map(counts, strconv.Itoa), ",")
	return categoryScore(category)
}

func computeScore2(cards string) int {
	group := groupHand(cards)
	if dict.HasKey(group, 'J') {
		values := make([]int, 0)
		for k, v := range group {
			if k != 'J' {
				values = append(values, v)
			}
		}
		if len(values) == 0 {
			values = append(values, 0)
		}
		slices.SortFunc(values, SortIntDesc)
		values[0] = values[0] + group['J']
		category := strings.Join(list.Map(values, strconv.Itoa), ",")
		return categoryScore(category)
	} else {
		return computeScore(cards)
	}
}

func compareHandsFn(power map[byte]int) func(StrInt, StrInt) int {
	return func(hand1, hand2 StrInt) int {
		cards1, cards2 := hand1.Str, hand2.Str
		score1 := computeScore(cards1)
		score2 := computeScore(cards2)
		if score1 == score2 {
			return compareCards(cards1, cards2, power)
		} else {
			return score1 - score2
		}
	}
}

func compareHands2Fn(power map[byte]int) func(StrInt, StrInt) int {
	return func(hand1, hand2 StrInt) int {
		cards1, cards2 := hand1.Str, hand2.Str
		score1 := computeScore2(cards1)
		score2 := computeScore2(cards2)
		if score1 == score2 {
			return compareCards(cards1, cards2, power)
		} else {
			return score1 - score2
		}
	}
}
