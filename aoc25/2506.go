package aoc25

import (
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day06() Solution {
	lines := data06(true)
	lastLine := len(lines) - 1

	// Part 1
	allOperands := make([][]int, 0)
	for i, line := range lines[:lastLine] {
		for col, x := range str.SpaceSplit(line) {
			op := number.ParseInt(x)
			if i == 0 {
				allOperands = append(allOperands, []int{op})
			} else {
				allOperands[col] = append(allOperands[col], op)
			}
		}
	}
	operators := str.SpaceSplit(lines[lastLine])
	total1 := list.Sum(list.IndexedMap(allOperands, func(i int, operands []int) int {
		return getResult(operators[i], operands)
	}))

	// Part 2
	grid := list.Map(lines[:lastLine], func(line string) []byte {
		return []byte(line)
	})
	gridRows, gridCols := GridBounds(grid).Tuple()
	indexes := make([]int, 0, len(operators))
	for i, char := range lines[lastLine] {
		if char == ' ' {
			continue
		}
		indexes = append(indexes, i)
	}
	indexes = append(indexes, gridCols+1)

	total2 := 0
	for i, operator := range operators {
		start, end := indexes[i], indexes[i+1]-1
		operands := make([]int, 0)
		for _, col := range list.NumRange(start, end) {
			num := make([]byte, 0)
			for row := range gridRows {
				num = append(num, grid[row][col])
			}
			operand := number.ParseInt(strings.TrimSpace(string(num)))
			operands = append(operands, operand)
		}
		total2 += getResult(operator, operands)
	}

	return NewSolution(total1, total2)
}

func data06(full bool) []string {
	return ReadRawLines(25, 6, full, false)
}

func getResult(operator string, operands []int) int {
	switch operator {
	case "+":
		return list.Sum(operands)
	case "*":
		return list.Product(operands)
	default:
		return 0
	}
}
