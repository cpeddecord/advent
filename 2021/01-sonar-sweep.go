package advent2021

/*
	SonarSweep measures the number of times a depth measurement is
	greater than its previous measurement

	https://adventofcode.com/2021/day/1
*/
func SonarSweep(m []int) int {
	d := 0

	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			d += 1
		}
	}

	return d
}

/*
	SonarSweepWindow measures a window of three values and records
	deviations between previous windows
*/
func SonarSweepWindow(m []int) int {
	d := 0
	prevSum := m[0] + m[1] + m[2]

	for i := 3; i < len(m); i++ {
		currSum := m[i] + m[i-1] + m[i-2]
		if currSum > prevSum {
			d += 1
		}

		prevSum = currSum
	}

	return d
}
