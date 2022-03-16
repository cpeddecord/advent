package advent2021

/*
	SonarSweepWindow measures a window of w values and records
	deviations between previous windows
*/
func SonarSweepWindow(m []int, w int) int {
	prevSum := 0
	for i := 0; i < w; i++ {
		prevSum += m[i]
	}

	d := 0
	for i := w; i < len(m); i++ {
		currSum := prevSum - m[i-w] + m[i]
		if currSum > prevSum {
			d += 1
		}

		prevSum = currSum
	}

	return d
}
