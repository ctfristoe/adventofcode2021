package main

import (
	"testing"
)

func getTestInput() []int {
	return []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
}

func TestCountDepthIncreases(t *testing.T) {
	test_input := getTestInput()
	expected_increases := 7
	actual_increases := CountDepthIncreases(test_input)
	if expected_increases != actual_increases {
		t.Errorf("%q depth increases calculated, wanted %q", actual_increases, expected_increases)
	}
}

func TestCountSlidingWindowIncreases(t *testing.T) {
	test_input := getTestInput()
	expected_increases := 5
	actual_increases := CountSlidingWindowIncreases(test_input)
	if expected_increases != actual_increases {
		t.Errorf("%q depth increases calculated, wanted %q", actual_increases, expected_increases)
	}
}

func TestReadProblemInput(t *testing.T) {
	expected := getTestInput()
	list, err := ReadIntListFromFile("test1.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	for index, item := range list {
		expected_item := expected[index]
		if item != expected_item {
			t.Errorf("Expected %q on line %q; found %q", expected_item, index, item)
		}
	}
}
