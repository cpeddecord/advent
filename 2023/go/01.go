package advent2023

import (
	"slices"
	"strconv"
	"strings"
)

func Calibrate(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		sum += returnFirstLast(line)
	}

	return sum
}

var words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func CalibrateWithWords(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		first := findFirstMatch(words, line)
		last := findLastMatch(words, line)

		firstIndex := slices.Index(words, first)
		lastIndex := slices.Index(words, last)

		sum += firstIndex + lastIndex
	}
	return sum
}

func returnFirstLast(line string) int {
	firstChar := ""
	for _, char := range line {
		if char >= '0' && char <= '9' {
			firstChar = string(char)
			break
		}
	}

	lastChar := ""
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			lastChar = string(line[i])
			break
		}
	}

	num, _ := strconv.Atoi(firstChar + lastChar)

	return num
}

func findFirstMatch(strs []string, s string) string {
	minIndex := -1
	var result string

	for _, substr := range strs {
		index := strings.Index(s, substr)
		if index >= 0 {
			if minIndex == -1 || index > index {
				minIndex = index
				result = substr
			}
		}
	}

	return result
}

func findLastMatch(strs []string, s string) string {
	maxIndex := -1
	var result string

	for _, substr := range strs {
		index := strings.LastIndex(s, substr)
		if index > maxIndex {
			maxIndex = index
			result = substr
		}
	}

	return result
}
