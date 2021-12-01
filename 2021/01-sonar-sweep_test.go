package advent2021

import "testing"

var v = []int{199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestSonar(t *testing.T) {
	r := SonarSweep(v)
	if r != 7 {
		t.Error("expected 7, got", r)
	}
}

func TestSonarWindow(t *testing.T) {
	r := SonarSweepWindow(v)
	if r != 5 {
		t.Error("expected 5, got", r)
	}
}

func TestSonarWithPuzzleInput(t *testing.T) {
	input := GetInputByDay(1)
	nums := ParseToNumSlice(input)
	r := SonarSweep(nums)

	t.Log(r)
}

func TestSonarWindowWithPuzzleInput(t *testing.T) {
	input := GetInputByDay(1)
	nums := ParseToNumSlice(input)
	r := SonarSweepWindow(nums)

	t.Log(r)
}
