package problems

import (
	"testing"

	"github.com/ctfristoe/adventofcode/utils"
)

func TestGetFinalPosition(t *testing.T) {
	const expected_horizontal = 15
	const expected_depth = 10
	test_input := utils.ReadLines("../inputs/depth_test.txt")

	final_position := GetFinalPosition(test_input)

	if expected_horizontal != final_position.Horizontal {
		t.Errorf("expected horizontal %v, found %v", expected_horizontal, final_position.Horizontal)
	}
	if expected_depth != final_position.Depth {
		t.Errorf("expected depth %v, found %v", expected_depth, final_position.Depth)
	}
}

func TestGetFinalPositionUsingAim(t *testing.T) {
	const expected_horizontal = 15
	const expected_depth = 60
	test_input := utils.ReadLines("../inputs/depth_test.txt")

	final_position := GetFinalPositionUsingAim(test_input)

	if expected_horizontal != final_position.Horizontal {
		t.Errorf("expected horizontal %v, found %v", expected_horizontal, final_position.Horizontal)
	}
	if expected_depth != final_position.Depth {
		t.Errorf("expected depth %v, found %v", expected_depth, final_position.Depth)
	}
}
