package problems

import (
	"testing"

	"github.com/ctfristoe/adventofcode/utils"
)

func TestCountDepthIncreases(t *testing.T) {
	const expected_increases = 7
	test_input := utils.ReadIntegerLines("../inputs/sonar_sweep_test.txt")

	actual_increases := CountDepthIncreases(test_input)

	if expected_increases != actual_increases {
		t.Errorf("%q depth increases calculated, wanted %q", actual_increases, expected_increases)
	}
}

func TestCountSlidingWindowIncreases(t *testing.T) {
	const expected_increases = 5
	test_input := utils.ReadIntegerLines("../inputs/sonar_sweep_test.txt")

	actual_increases := CountSlidingWindowIncreases(test_input)

	if expected_increases != actual_increases {
		t.Errorf("%q depth increases calculated, wanted %q", actual_increases, expected_increases)
	}
}
