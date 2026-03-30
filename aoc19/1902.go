package aoc19

import . "github.com/zeroibot/aoc-go/aoc"

func Day02() Solution {
	// Part 1
	numbers := data02(true)
	numbers[1] = 12
	numbers[2] = 2
	output1 := runProgram02(numbers)

	// Part 2
	output2 := 0
mainLoop:
	for noun := range 100 {
		for verb := range 100 {
			numbers = data02(true)
			numbers[1] = noun
			numbers[2] = verb
			if runProgram02(numbers) == 19690720 {
				output2 = (100 * noun) + verb
				break mainLoop
			}
		}
	}

	return NewSolution(output1, output2)
}

func data02(full bool) []int {
	line := ReadFirstLine(19, 2, full)
	return ToIntList(line, ",")
}

func runProgram02(numbers []int) int {
	i := 0
	for {
		cmd := numbers[i]
		if cmd == 99 {
			break
		}
		in1, in2, out := numbers[i+1], numbers[i+2], numbers[i+3]
		a, b := numbers[in1], numbers[in2]
		switch cmd {
		case 1:
			numbers[out] = a + b
		case 2:
			numbers[out] = a * b
		}
		i += 4
	}
	return numbers[0]
}
