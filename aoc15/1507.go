package aoc15

import (
	"maps"
	"strconv"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

func Day07() Solution {
	operations := data07(true)

	// Part 1
	a1 := solveA(operations, nil)

	// Part 2
	override := map[string]int{"b": a1}
	a2 := solveA(operations, override)

	return NewSolution(a1, a2)
}

const (
	LET    string = "LET"
	AND    string = "AND"
	OR     string = "OR"
	NOT    string = "NOT"
	LSHIFT string = "LSHIFT"
	RSHIFT string = "RSHIFT"
)

func data07(full bool) []Operation {
	binaryCommands := []string{AND, OR, LSHIFT, RSHIFT}
	return list.Map(ReadLines(15, 7, full), func(line string) Operation {
		var opType, p1, p2 string
		p := str.CleanSplit(line, "->")
		expr, result := p[0], p[1]
		found := false
		for _, cmd := range binaryCommands {
			if strings.Contains(expr, cmd) {
				v := str.CleanSplit(expr, cmd)
				p1, p2 = v[0], v[1]
				opType = cmd
				found = true
				break
			}
		}
		if !found {
			if strings.Contains(expr, NOT) {
				p1 = Last(str.CleanSplit(expr, NOT), 1)
				opType = NOT
			} else {
				p1 = expr
				opType = LET
			}
		}

		return Operation{
			Type:      opType,
			Result:    result,
			Param1:    p1,
			Param2:    p2,
			Variables: []string{},
			Values:    []int{},
		}
	})
}

type Operation struct {
	Type      string
	Result    string
	Param1    string
	Param2    string
	Variables []string
	Values    []int
}

func (o *Operation) SetVariables() {
	p1 := tryParseInt(o.Param1)
	if p1 == nil {
		o.Variables = append(o.Variables, o.Param1)
	} else {
		o.Values = append(o.Values, *p1)
	}
	p2 := tryParseInt(o.Param2)
	if p2 == nil {
		o.Variables = append(o.Variables, o.Param2)
	} else {
		o.Values = append(o.Values, *p2)
	}
}

func solveA(operations []Operation, override map[string]int) int {
	value := make(map[string]int)
	q := make([]Operation, 0)

	for _, op := range operations {
		switch op.Type {
		case LET:
			x := tryParseInt(op.Param1)
			if x == nil {
				op.Variables = append(op.Variables, op.Param1)
				q = append(q, op)
			} else {
				value[op.Result] = *x
			}
		case NOT, LSHIFT, RSHIFT:
			op.Variables = append(op.Variables, op.Param1)
			q = append(q, op)
		case AND, OR:
			op.SetVariables()
			q = append(q, op)
		}
	}

	if override != nil {
		maps.Copy(value, override)
	}

	for len(q) > 0 {
		q2 := make([]Operation, 0)
		for _, op := range q {
			hasUnknown := list.Any(op.Variables, func(v string) bool {
				_, found := value[v]
				return !found
			})
			if hasUnknown {
				q2 = append(q2, op)
				continue
			}
			var0 := op.Variables[0]
			switch op.Type {
			case AND, OR:
				var param int
				if len(op.Variables) == 2 {
					param = value[op.Variables[1]]
				} else {
					param = op.Values[0]
				}
				var result int
				switch op.Type {
				case AND:
					result = value[var0] & param
				case OR:
					result = value[var0] | param
				}
				value[op.Result] = result
			case LSHIFT:
				value[op.Result] = value[var0] << number.ParseInt(op.Param2)
			case RSHIFT:
				value[op.Result] = value[var0] >> number.ParseInt(op.Param2)
			case NOT:
				value[op.Result] = ^value[var0]
			case LET:
				value[op.Result] = value[var0]
			}
			if value[op.Result] < 0 {
				value[op.Result] += 65536
			}
		}
		q = q2
		if dict.HasKey(value, "a") {
			break
		}
	}
	return value["a"]
}

func tryParseInt(text string) *int {
	x, err := strconv.Atoi(text)
	if err == nil {
		return &x
	} else {
		return nil
	}
}
