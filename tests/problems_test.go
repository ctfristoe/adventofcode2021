package tests

import (
	"testing"

	"github.com/ctfristoe/adventofcode/problems"
)

func TestSonarSweepProblemOne(t *testing.T) {
	const expected = 7

	problemSet := problems.SonarSweep{}
	answer := problemSet.DoProblemOne("../inputs/sonar_sweep_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSonarSweepProblemTwo(t *testing.T) {
	const expected = 5

	problemSet := problems.SonarSweep{}
	answer := problemSet.DoProblemTwo("../inputs/sonar_sweep_test.txt")

	AssertEqual(t, expected, answer)
}

func TestDepthProblemOne(t *testing.T) {
	const expected = 150

	problemSet := problems.Depth{}
	answer := problemSet.DoProblemOne("../inputs/depth_test.txt")

	AssertEqual(t, expected, answer)
}

func TestDepthProblemTwo(t *testing.T) {
	const expected = 900

	problemSet := problems.Depth{}
	answer := problemSet.DoProblemTwo("../inputs/depth_test.txt")

	AssertEqual(t, expected, answer)
}

func TestBinaryDiagnosisProblemOne(t *testing.T) {
	const expected = 198

	problemSet := problems.BinaryDiagnostic{}
	answer := problemSet.DoProblemOne("../inputs/binary_diagnostic_test.txt")

	AssertEqual(t, expected, answer)
}

func TestBinaryDiagnosisProblemTwo(t *testing.T) {
	const expected = 230

	problemSet := problems.BinaryDiagnostic{}
	answer := problemSet.DoProblemTwo("../inputs/binary_diagnostic_test.txt")

	AssertEqual(t, expected, answer)
}

func TestGiantSquidProblemOne(t *testing.T) {
	const expected = 4512

	problemSet := problems.GiantSquid{}
	answer := problemSet.DoProblemOne("../inputs/giant_squid_test.txt")

	AssertEqual(t, expected, answer)
}

func TestGiantSquidProblemTwo(t *testing.T) {
	const expected = 1924

	problemSet := problems.GiantSquid{}
	answer := problemSet.DoProblemTwo("../inputs/giant_squid_test.txt")

	AssertEqual(t, expected, answer)
}
