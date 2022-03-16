package advent2021

import "testing"

var hydroTest = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestParseToVents(t *testing.T) {
	r := ParseToVents(hydroTest)

	first := r[0]
	if first.a != [2]int{0, 9} && first.b != [2]int{5, 9} && first.direction != "horizontal" {
		t.Error("expected 0,9;5,9, got", first)
	}

	sec := r[1]
	if sec.a != [2]int{8, 0} && sec.b != [2]int{0, 8} && sec.direction != "diagonal" {
		t.Error("expected 8,0;0,8, got", sec)
	}
}

func TestFindOverlappingVents(t *testing.T) {
	v := ParseToVents(hydroTest)
	r := FindOverlappingVents(v, false)

	if r != 5 {
		t.Error("expected 5, got", r)
	}
}

func TestFindOverlappingVentsWithDiags(t *testing.T) {
	v := ParseToVents(hydroTest)
	r := FindOverlappingVents(v, true)

	if r != 12 {
		t.Error("expected 12, got", r)
	}
}

func TestFindOverlappingVentsWithPuzzleInput(t *testing.T) {
	v := ParseToVents(GetInputByDay(5))
	partI := FindOverlappingVents(v, false)
	partII := FindOverlappingVents(v, true)

	t.Log("part I", partI)
	t.Log("part II", partII)
}
