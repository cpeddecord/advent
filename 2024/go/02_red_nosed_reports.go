package advent2024

import (
	"fmt"
	"strconv"
	"strings"
)

func SafetyReport(input string) int {
	safe := 0

	nums_list := lineParser(input)
	for _, nums := range nums_list {
		if determineSafetyWithoutDampener(nums) {
			safe++
		}
	}

	return safe
}

func SafetyReportWithDampener(input string) int {
	safe := 0

	nums_list := lineParser(input)
	for _, nums := range nums_list {
		if determineSafetyWithDampener(nums) {
			safe++
		}
	}

	return safe
}

func determineSafetyWithoutDampener(nums []int) bool {
	shouldAscend := false

	if nums[0] < nums[1] {
		shouldAscend = true
	}

	if nums[0] == nums[1] {
		return false
	}

	for i := 0; i < len(nums)-1; i++ {
		if shouldAscend {
			if nums[i] >= nums[i+1] || nums[i+1]-nums[i] > 3 {
				return false
			}
		}

		if !shouldAscend {
			if nums[i] <= nums[i+1] || nums[i]-nums[i+1] > 3 {
				return false
			}
		}
	}

	return true
}

func determineSafetyWithDampener(nums []int) bool {
	faults := 0
	shouldAscend := false
	shouldDescend := false
	startIdx := 0

	if nums[0] == nums[1] {
		faults++
		startIdx = 1
	}

	if nums[startIdx] < nums[startIdx+1] {
		shouldAscend = true
	}

	if nums[startIdx] > nums[startIdx+1] {
		shouldDescend = true
	}

	for i := 0; i < len(nums)-1; i++ {
		if shouldAscend {
			if nums[i] >= nums[i+1] || nums[i+1]-nums[i] > 3 {
				faults++
				if faults > 1 {
					return false
				}
			}
		}

		if shouldDescend {
			if nums[i] <= nums[i+1] || nums[i]-nums[i+1] > 3 {
				faults++
				if faults > 1 {
					return false
				}
			}
		}
	}

	if faults > 1 {
		return false
	}

	fmt.Println(nums)

	return true
}

func lineParser(input string) [][]int {
	nums_list := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], _ = strconv.Atoi(p)
		}

		nums_list = append(nums_list, nums)
	}

	return nums_list
}
