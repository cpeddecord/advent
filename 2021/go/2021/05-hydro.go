package advent2021

import (
	"regexp"
	"strconv"
)

type VentNode struct {
	a         [2]int
	b         [2]int
	direction string
}

func ParseToVents(i string) []VentNode {
	re := regexp.MustCompile(`(?m)((\d*),(\d*))`)
	nodes := re.FindAllStringSubmatch(i, -1)

	ret := []VentNode{}

	nullVent := [2]int{-1, -1}
	a := nullVent
	b := nullVent
	for i, n := range nodes {
		if len(n) == 0 {
			continue
		}

		x, _ := strconv.Atoi(n[2])
		y, _ := strconv.Atoi(n[3])

		if i%2 == 0 {
			a = [2]int{x, y}
		} else {
			b = [2]int{x, y}

			direction := "diagonal"
			if a[0] == b[0] {
				direction = "horizontal"
			}
			if a[1] == b[1] {
				direction = "vertical"
			}

			ret = append(ret, VentNode{
				a:         a,
				b:         b,
				direction: direction,
			})

			a = nullVent
			b = nullVent
		}
	}

	return ret
}

func minOf(a, b int) int {
	if b < a {
		return b
	}

	return a
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func addTo(a, b int) int {
	return a + b
}

func subFrom(a, b int) int {
	return a - b
}

func FindOverlappingVents(vents []VentNode, withDiags bool) int {
	o := 0

	m := make(map[[2]int]int)

	for _, v := range vents {
		if !withDiags && v.direction == "diagonal" {
			continue
		}

		if v.direction == "diagonal" {
			start := v.a
			end := v.b

			// get start coord to iterate l-r
			minX := minOf(v.a[0], v.b[0])
			if end[0] == minX {
				start = v.b
				end = v.a
			}

			// get fn for going up/down
			fn := addTo
			if start[1] < end[1] {
				fn = subFrom
			}

			for i := 0; i <= end[0]-start[0]; i++ {
				x := fn(start[0], i)
				y := fn(start[1], i)

				m[[2]int{x, y}] += 1
			}

		}

		if v.direction == "horizontal" {
			x := v.a[0]

			min := minOf(v.a[1], v.b[1])
			max := maxOf(v.a[1], v.b[1])

			for i := min; i <= max; i++ {
				m[[2]int{x, i}] += 1
			}
		}

		if v.direction == "vertical" {
			y := v.a[1]

			min := minOf(v.a[0], v.b[0])
			max := maxOf(v.a[0], v.b[0])

			for i := min; i <= max; i++ {
				m[[2]int{i, y}] += 1
			}
		}
	}

	for _, v := range m {
		if v > 1 {
			o++
		}
	}

	return o
}
