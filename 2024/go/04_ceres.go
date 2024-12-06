package advent2024

import (
	"strings"
)

func FindXMASCount(input string) int {
	lines := parseLines(input)
	count := 0
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	rows, cols := len(lines), len(lines[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if lines[i][j] != "X" {
				continue
			}

			for _, dir := range directions {
				if checkWord(lines, i, j, dir, "XMAS") {
					count++
				}
			}
		}
	}
	return count
}

func checkWord(grid [][]string, x, y int, dir []int, word string) bool {
	for k, char := range word {
		nx, ny := x+dir[0]*k, y+dir[1]*k
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) {
			return false
		}

		if grid[nx][ny] != string(char) {
			return false
		}
	}

	return true
}

func parseLines(input string) [][]string {
	lines := strings.Split(input, "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	data := make([][]string, len(lines))
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	return data
}
