package advent2021

import (
	"strconv"
)

func GammaEpsilon(v []string) int64 {
	g := ""
	e := ""

	width := len(v[0])
	for i := 0; i < width; i++ {
		runningZeroes := 0

		for _, r := range v {
			if r[i] == 48 {
				runningZeroes += 1
			}
		}

		if runningZeroes > (len(v))/2 {
			g += "0"
			e += "1"
		} else {
			g += "1"
			e += "0"
		}
	}

	gamma, _ := strconv.ParseInt(g, 2, 64)
	epsilon, _ := strconv.ParseInt(e, 2, 64)

	return gamma * epsilon
}

func LifeSupport(v []string) int64 {
	o := ReadRating(v, 0, withOxygenCriteria)
	c := ReadRating(v, 0, withCarbonCriteria)

	return o * c
}

type criteria func(z int, v []string) bool

func withCarbonCriteria(z int, v []string) bool {
	return z <= len(v)/2
}

func withOxygenCriteria(z int, v []string) bool {
	return z > (len(v))/2
}

func ReadRating(v []string, i int, fn criteria) int64 {
	if len(v) == 1 {
		rating, _ := strconv.ParseInt(v[0], 2, 64)
		return rating
	}

	runningZeroes := 0
	leadingZeros := []string{}
	leadingOnes := []string{}
	for _, r := range v {
		if r[i] == 48 {
			runningZeroes += 1
			leadingZeros = append(leadingZeros, r)
		} else {
			leadingOnes = append(leadingOnes, r)
		}
	}

	if fn(runningZeroes, v) {
		return ReadRating(leadingZeros, i+1, fn)
	}

	return ReadRating(leadingOnes, i+1, fn)
}
