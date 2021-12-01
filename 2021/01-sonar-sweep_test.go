package advent2021

import "testing"

func TestSonar(t *testing.T) {
	v := []int{199,
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

	r := SonarSweep(v)
	if r != 7 {
		t.Error("expected 7, got", r)
	}
}

func TestSonarWithPuzzleInput(t *testing.T) {
	input := GetInputByDay(1)
	nums := ParseToNumSlice(input)
	r := SonarSweep(nums)

	t.Log(r)
}
