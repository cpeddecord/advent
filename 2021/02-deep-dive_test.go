package advent2021

import "testing"

var v = ParseToDivingInstructions(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)

func TestDive(t *testing.T) {
	r := Dive(v)

	if r != 150 {
		t.Error("expected 150, got", r)
	}
}

func TestDiveWithAim(t *testing.T) {
	r := DiveWithAim(v)

	if r != 900 {
		t.Error("expected 900, got", r)
	}
}

func TestDiveWithPuzzleInput(t *testing.T) {
	i := GetInputByDay(2)
	v := ParseToDivingInstructions(i)

	partI := Dive(v)
	partII := DiveWithAim(v)

	t.Log("Part I", partI)
	t.Log("Part II", partII)
}
