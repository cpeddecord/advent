package advent2021

import (
	"regexp"
	"strconv"
	"strings"
)

type BingoBoard struct {
	coords   map[int][]int
	rowState map[int]int
	colState map[int]int
	numState map[int]bool
	hasBingo bool
}

func createBingoBoard(n map[int][]int) BingoBoard {
	return BingoBoard{
		coords:   n,
		rowState: map[int]int{},
		colState: map[int]int{},
		numState: map[int]bool{},
		hasBingo: false,
	}
}

func (b *BingoBoard) addNum(n int) {
	_, ok := b.coords[n]
	if !ok {
		return
	}

	coords := b.coords[n]
	b.numState[n] = true

	b.rowState[coords[0]] += 1
	b.colState[coords[1]] += 1

	if b.rowState[coords[0]] == 5 {
		b.hasBingo = true
	}

	if b.colState[coords[1]] == 5 {
		b.hasBingo = true
	}
}

func ParseToBingoGame(s string) ([]int, []BingoBoard) {
	instructions := []int{}
	boards := []BingoBoard{}

	splits := strings.Split(s, "\n")

	for _, s := range strings.Split(splits[0], ",") {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}

	re := regexp.MustCompile(`\s{1,2}`)
	m := make(map[int][]int)
	r := 0
	for i := 2; i < len(splits); i++ {
		if len(splits[i]) == 0 {
			boards = append(boards, createBingoBoard(m))
			r = 0
			m = make(map[int][]int)

			continue
		}

		for i, s := range re.Split(strings.Trim(splits[i], " "), -1) {
			n, _ := strconv.Atoi(s)
			m[n] = []int{r, i}
		}

		r++
	}

	return instructions, boards
}

func PlayBingo(i []int, boards []BingoBoard) int {
	winningNum := -1
	unmarkedSums := 0

	for _, v := range i {
		if winningNum != -1 {
			break
		}

		for _, b := range boards {
			b.addNum(v)
			if b.hasBingo {
				winningNum = v

				for k := range b.coords {
					if b.numState[k] {
						continue
					}

					unmarkedSums += k
				}

				break
			}
		}
	}

	return winningNum * unmarkedSums
}

func LetTheSquidWin(i []int, boards []BingoBoard) int {
	// lastBoard := boards[0]
	// unmarkedSums := 0
	// boardsWithoutBingo := len(boards)
	// idx := 0

	// for boardsWithoutBingo != 0 {
	// 	fmt.Println(i[idx])
	// 	for _, b := range boards {
	// 		b.addNum(i[idx])

	// 		if b.hasBingo {
	// 			lastBoard = b
	// 			boardsWithoutBingo--
	// 			idx--
	// 		}
	// 	}

	// 	idx++
	// }

	// for k := range lastBoard.coords {
	// 	if lastBoard.numState[k] {
	// 		continue
	// 	}

	// 	unmarkedSums += k
	// }
	// 	return unmarkedSums * i[idx]

	boardsWithoutBingo := 0
	winningNum := 0
	unmarkedSums := 0

	for _, v := range i {
		if boardsWithoutBingo == 0 {
			break
		}

		for _, b := range boards {
			b.addNum(v)

			if b.hasBingo {
				boardsWithoutBingo--

				if boardsWithoutBingo == 0 {
					winningNum = v
					for k := range b.coords {
						if b.numState[k] {
							continue
						}

						unmarkedSums += k
					}
				}
			}
		}
	}

	return winningNum * unmarkedSums
}
