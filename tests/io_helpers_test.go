package tests

import (
	"testing"

	"github.com/ctfristoe/adventofcode/problems"
)

func TestReadProblemInput(t *testing.T) {
	expected := []int{0, 100, 9980, -1, 1000000}

	input := problems.ReadIntegerLines("../inputs/integers_test.txt")

	AssertEqual(t, expected, input)
}
