package advent2024

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func TotalDistance(input string) int {
	total := 0
	left_list := []int{}
	right_list := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var num1, num2 int
		fmt.Sscanf(line, "%d   %d", &num1, &num2)
		left_list = append(left_list, num1)
		right_list = append(right_list, num2)
	}

	sort.Ints(left_list)
	sort.Ints(right_list)

	for i := 0; i < len(left_list); i++ {
		total += int(math.Abs(float64(left_list[i] - right_list[i])))
	}

	return total
}

func ScoreSimilarity(input string) int {
	total := 0

	left_list := []int{}
	right_list := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var num1, num2 int
		fmt.Sscanf(line, "%d   %d", &num1, &num2)
		left_list = append(left_list, num1)
		right_list = append(right_list, num2)
	}

	for _, num := range left_list {
		found := 0
		for _, right_num := range right_list {
			if num == right_num {
				found++
			}
		}

		total += num * found

	}

	return total
}
