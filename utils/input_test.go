package utils

import (
	"testing"
)

func TestReadProblemInput(t *testing.T) {
	const expected_len = 5

	list := ReadIntegerLines("../inputs/integers_test.txt")

	if len(list) != expected_len {
		t.Errorf("Expected %d items in %v", expected_len, list)
	}
}
