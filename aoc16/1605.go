package aoc16

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/ds"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
)

func Day05() Solution {
	door := data05(true)
	pwdLength := 8

	// Part 1
	hashGen := md5HashGenerator(door, "00000", 0)
	next, stop := iter.Pull2(hashGen)
	pwdChars := slices.Repeat([]byte{'.'}, pwdLength)
	for i := range pwdLength {
		_, hash, ok := next()
		if ok {
			pwdChars[i] = hash[5]
			fmt.Println(string(pwdChars))
		}
	}
	stop()
	pwd1 := string(pwdChars)

	// Part 2
	hashGen = md5HashGenerator(door, "00000", 0)
	next, stop = iter.Pull2(hashGen)
	indexes := ds.NewSet[byte]()
	for x := range pwdLength {
		b := fmt.Sprintf("%d", x)[0]
		indexes.Add(b)
	}
	pwdChars = slices.Repeat([]byte{'.'}, pwdLength)
	for {
		_, hash, ok := next()
		if !ok {
			break
		}
		if !indexes.Has(hash[5]) {
			continue
		}
		idx := number.ParseInt(string(hash[5]))
		if pwdChars[idx] == '.' {
			pwdChars[idx] = hash[6]
			fmt.Println(string(pwdChars))
		}
		allSet := list.All(pwdChars, func(char byte) bool {
			return char != '.'
		})
		if allSet {
			break
		}
	}
	stop()
	pwd2 := string(pwdChars)

	return NewSolution(pwd1, pwd2)
}

func data05(full bool) string {
	return ReadFirstLine(16, 5, full)
}

func md5HashGenerator(key string, goal string, start int) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := start
		for {
			word := fmt.Sprintf("%s%d", key, i)
			hash := MD5Hash(word)
			if strings.HasPrefix(hash, goal) {
				if !yield(i, hash) {
					return
				}
			}
			i += 1
		}
	}
}
