package advent2021

/*
	SonarSweep measures the number of times a depth measurement is
	greater than its previous measurement

	https://adventofcode.com/2021/day/1
*/
func SonarSweep(m []int) int {
	deviations := 0

	// index 0 skipped, nothing to compare to
	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			deviations += 1
		}
	}

	return deviations
}
