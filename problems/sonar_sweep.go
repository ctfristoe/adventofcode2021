package problems

type SonarSweep struct{}

func (SonarSweep) DoProblemOne(filepath string) int {
	input := ReadIntegerLines(filepath)
	return countDepthIncreases(input)
}

func (SonarSweep) DoProblemTwo(filepath string) int {
	input := ReadIntegerLines(filepath)
	return countSlidingWindowIncreases(input)
}

func countDepthIncreases(input []int) int {
	n := 0
	for index, item := range input[1:] {
		if item > input[index] {
			n++
		}
	}
	return n
}

func countSlidingWindowIncreases(input []int) int {
	n := 0
	for index, item := range input[3:] {
		if item > input[index] {
			n++
		}
	}
	return n
}
