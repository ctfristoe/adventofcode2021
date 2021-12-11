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

func TestHydrothermalVentureProblemOne(t *testing.T) {
	const expected = 5

	problemSet := problems.HydrothermalVenture{}
	answer := problemSet.DoProblemOne("../inputs/hydrothermal_venture_test.txt")

	AssertEqual(t, expected, answer)
}

func TestHydrothermalVentureProblemTwo(t *testing.T) {
	const expected = 12

	problemSet := problems.HydrothermalVenture{}
	answer := problemSet.DoProblemTwo("../inputs/hydrothermal_venture_test.txt")

	AssertEqual(t, expected, answer)
}

func TestLanternfishProblemOne(t *testing.T) {
	const expected = 5934

	problemSet := problems.Lanternfish{}
	answer := problemSet.DoProblemOne("../inputs/lanternfish_test.txt")

	AssertEqual(t, expected, answer)
}

func TestLanternfishProblemTwo(t *testing.T) {
	const expected = 26984457539

	problemSet := problems.Lanternfish{}
	answer := problemSet.DoProblemTwo("../inputs/lanternfish_test.txt")

	AssertEqual(t, expected, answer)
}

func TestWhalesProblemOne(t *testing.T) {
	const expected = 37

	problemSet := problems.Whales{}
	answer := problemSet.DoProblemOne("../inputs/whales_test.txt")

	AssertEqual(t, expected, answer)
}

func TestWhalesProblemTwo(t *testing.T) {
	const expected = 168

	problemSet := problems.Whales{}
	answer := problemSet.DoProblemTwo("../inputs/whales_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSevenSegmentSearchProblemOne(t *testing.T) {
	const expected = 26

	problemSet := problems.SevenSegmentSearch{}
	answer := problemSet.DoProblemOne("../inputs/seven_segment_search_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSmokeBasinProblemOne(t *testing.T) {
	const expected = 15

	problemSet := problems.SmokeBasin{}
	answer := problemSet.DoProblemOne("../inputs/smoke_basin_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSmokeBasinProblemTwo(t *testing.T) {
	const expected = 1134

	problemSet := problems.SmokeBasin{}
	answer := problemSet.DoProblemTwo("../inputs/smoke_basin_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSyntaxScoringProblemOne(t *testing.T) {
	const expected = 26397

	problemSet := problems.SyntaxScoring{}
	answer := problemSet.DoProblemOne("../inputs/syntax_scoring_test.txt")

	AssertEqual(t, expected, answer)
}

func TestSyntaxScoringProblemTwo(t *testing.T) {
	const expected = 288957

	problemSet := problems.SyntaxScoring{}
	answer := problemSet.DoProblemTwo("../inputs/syntax_scoring_test.txt")

	AssertEqual(t, expected, answer)
}
