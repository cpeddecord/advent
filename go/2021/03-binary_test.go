package advent2021

import "testing"

var h = ParseToStringSlice(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)

func TestGammaEpsilon(t *testing.T) {
	r := GammaEpsilon(h)

	if r != 198 {
		t.Error("expected 198, got", r)
	}
}

func TestLifeSupport(t *testing.T) {
	r := LifeSupport(h)

	if r != 230 {
		t.Error("expected 230, got", r)

	}
}

func TestBinariesWithPuzzleInput(t *testing.T) {
	var i = ParseToStringSlice(GetInputByDay(3))
	partI := GammaEpsilon(i)
	partII := LifeSupport(i)

	t.Log("Part I", partI)
	t.Log("Part II", partII)

}

func TestLifeSupportWithPuzzleInput(t *testing.T) {

}
