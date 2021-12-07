package problems

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinMax(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func NthTriangleNumber(n int) int {
	if n < 0 {
		panic("negative integer invalid")
	}
	return n * (n + 1) / 2
}
