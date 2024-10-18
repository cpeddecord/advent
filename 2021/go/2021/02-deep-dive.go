package advent2021

import (
	"regexp"
	"strconv"
	"strings"
)

type DivingInstruction struct {
	direction string
	size      int
}

type Instructions []DivingInstruction

func Dive(i Instructions) int {
	hor := 0
	depth := 0

	for _, v := range i {
		switch v.direction {
		case "forward":
			hor += v.size
		case "down":
			depth += v.size
		case "up":
			depth -= v.size
		}

	}

	return hor * depth
}

func DiveWithAim(i Instructions) int {
	hor := 0
	depth := 0
	aim := 0

	for _, v := range i {
		switch v.direction {
		case "forward":
			hor += v.size
			depth += v.size * aim
		case "down":
			aim += v.size
		case "up":
			aim -= v.size
		}
	}

	return hor * depth
}

/*
  ParseToDivingInstructions produces a slice of Instructions from the
  following data:

    forward 5
    down 5
*/
func ParseToDivingInstructions(i string) Instructions {
	splits := strings.Split(i, "\n")
	re := regexp.MustCompile(`(\w*) (\d*)`)

	var ret []DivingInstruction
	for _, v := range splits {
		t := re.FindStringSubmatch(v)

		if len(t) == 0 {
			continue
		}

		n, _ := strconv.Atoi(t[2])
		ret = append(ret, DivingInstruction{t[1], n})
	}

	return ret
}
