package advent2021

import "testing"

func TestSonarWindow(t *testing.T) {
	v := []int{
		199,
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

	r := SonarSweepWindow(v, 1)
	if r != 7 {
		t.Error("expected 7, got", r)

	}

	r = SonarSweepWindow(v, 3)
	if r != 5 {
		t.Error("expected 5, got", r)
	}

}

func TestSonarWindowWithPuzzleInput(t *testing.T) {
	input := GetInputByDay(1)
	nums := ParseToNumSlice(input)

	single := SonarSweepWindow(nums, 1)
	triple := SonarSweepWindow(nums, 3)

	t.Log("Part I", single)
	t.Log("Part II", triple)
}
