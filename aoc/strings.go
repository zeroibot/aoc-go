package aoc

import (
	"crypto/md5"
	"fmt"
	"slices"
	"strings"

	"github.com/zeroibot/fn/list"
)

func MD5Hash(word string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(word)))
}

func HasTwins(word string, gap int) bool {
	back := gap + 1
	limit := len(word)
	for i := back; i < limit; i++ {
		if word[i] == word[i-back] {
			return true
		}
	}
	return false
}

func CharFreq(word string, skip []rune) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range word {
		if skip != nil && slices.Contains(skip, char) {
			continue
		}
		freq[char] += 1
	}
	return freq
}

func ReverseString(word string) string {
	chars := make([]byte, 0)
	i := len(word) - 1
	for i >= 0 {
		chars = append(chars, word[i])
		i -= 1
	}
	return string(chars)
}

func SortedString(word string) string {
	letters := make([]rune, len(word))
	for i, letter := range word {
		letters[i] = letter
	}
	slices.Sort(letters)
	return string(letters)
}

func LowerChar(char byte) byte {
	return strings.ToLower(string(char))[0]
}

func GroupChunks(word string) []string {
	chunks := make([]string, 0)
	letters := []rune(word)
	curr, count := letters[0], 1
	for i := 1; i < len(letters); i++ {
		char := letters[i]
		if char == curr {
			count += 1
		} else {
			chunks = append(chunks, RepeatChar(curr, count))
			curr, count = char, 1
		}
	}
	chunks = append(chunks, RepeatChar(curr, count))
	return chunks
}

func RepeatChar(char rune, count int) string {
	return string(slices.Repeat([]rune{char}, count))
}

func RuneToString(char rune) string {
	return string([]rune{char})
}

func StringChars(text string) []string {
	return list.Map([]rune(text), RuneToString)
}
