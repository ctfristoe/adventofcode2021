package problems

func CountDepthIncreases(input []int) int {
	n := 0
	for index, item := range input[1:] {
		if item > input[index] {
			n++
		}
	}
	return n
}

func CountSlidingWindowIncreases(input []int) int {
	n := 0
	// current := input[0] + input[1] + input[2]
	for index, item := range input[3:] {
		if item > input[index] {
			n++
		}
	}
	return n
}
