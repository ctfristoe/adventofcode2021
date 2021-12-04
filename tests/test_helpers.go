package tests

import "testing"

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %d, found %d", expected, actual)
	}
}
