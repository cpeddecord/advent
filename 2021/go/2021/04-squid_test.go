package advent2021

import (
	"testing"
)

var i = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

func slicesEqual(a, b []int) bool {
	for i, v := range a {
		if len(a) != len(b) {
			return false
		}
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParseToBingoGame(t *testing.T) {
	insts, boards := ParseToBingoGame(i)

	if !slicesEqual(insts[0:5], []int{7, 4, 9, 5, 11}) {
		t.Error("instructions not correct, got", insts[0:4])
	}

	if len(boards) != 3 {
		t.Error("received wrong number of boards, wanted 3 got", len(boards), boards)
	}

	coords := boards[0].coords

	if len(coords) == 0 {
		t.Error("no coords created")
	}

	if coords[23][0] != 1 && coords[23][1] != 2 {
		t.Error("coords incorrect, expected 2[0, 1], got", coords[2])
	}
}

func TestBingoMethods(t *testing.T) {
	_, boards := ParseToBingoGame(i)
	b := boards[0]
	c := boards[2]

	b.addNum(1000)
	if len(b.numState) != 0 {
		t.Error("update failed")
	}

	b.addNum(20)

	if b.rowState[4] != 1 && b.colState[2] != 1 {
		t.Error("update failed, got", b.rowState, b.colState)
	}

	if !b.numState[20] {
		t.Error("update failed, got", b.numState)
	}

	b.addNum(1)

	if b.rowState[4] != 2 {
		t.Error("row update failed, expected 2 got", b.rowState)
	}

	if b.colState[0] != 1 {
		t.Error("col update failed, expected 1 got", b.colState)
	}

	if b.hasBingo {
		t.Error("bingo is not bingo", b.rowState, b.colState)
	}

	b.addNum(12)
	b.addNum(15)

	if b.hasBingo {
		t.Error("bingo is not bingo", b.rowState, b.colState)
	}

	b.addNum(19)

	if !b.hasBingo {
		t.Error("bingo is not bingo", b.rowState)
	}

	for i, n := range []int{14, 10, 18, 22, 2} {
		c.addNum(n)

		if i != 4 && c.hasBingo {
			t.Error("bingo is not bingo", c.colState)
		}

		if i == 4 && !c.hasBingo {
			t.Error("bingo is not bingo", c.colState)
		}
	}

}

func TestPlayBingo(t *testing.T) {
	insts, boards := ParseToBingoGame(i)
	r := PlayBingo(insts, boards)

	if r != 4512 {
		t.Error("expected 4512, got", r)
	}
}

func TestLetTheSquidWin(t *testing.T) {
	insts, boards := ParseToBingoGame(i)
	r := LetTheSquidWin(insts, boards)

	if r != 1924 {
		t.Error("expected 1924, got", r)
	}

}

func TestPlayBingoWithPuzzleInput(t *testing.T) {
	insts, boards := ParseToBingoGame(GetInputByDay(4))
	partI := PlayBingo(insts, boards)
	partII := LetTheSquidWin(insts, boards)

	t.Log("Part I", partI)
	t.Log("Part II", partII)
}
