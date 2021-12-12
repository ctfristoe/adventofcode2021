package main

import (
	"fmt"

	"github.com/ctfristoe/adventofcode/problems"
)

type ProblemSet interface {
	DoProblemOne(filepath string) int
	DoProblemTwo(filepath string) int
}

type ProblemSetRunner struct {
	problemSet ProblemSet
	filepath   string
}

func (runner ProblemSetRunner) PrintResults(number int) {
	answer1 := runner.problemSet.DoProblemOne(runner.filepath)
	answer2 := runner.problemSet.DoProblemTwo(runner.filepath)
	fmt.Printf("Problem %d.1: %d \n", number, answer1)
	fmt.Printf("Problem %d.2: %d \n\n", number, answer2)
}

func main() {
	runners := []ProblemSetRunner{
		{problems.SonarSweep{}, "inputs/sonar_sweep.txt"},
		{problems.Depth{}, "inputs/depth.txt"},
		{problems.BinaryDiagnostic{}, "inputs/binary_diagnostic.txt"},
		{problems.GiantSquid{}, "inputs/giant_squid.txt"},
		{problems.HydrothermalVenture{}, "inputs/hydrothermal_venture.txt"},
		{problems.Lanternfish{}, "inputs/lanternfish.txt"},
		{problems.Whales{}, "inputs/whales.txt"},
		{problems.SevenSegmentSearch{}, "inputs/seven_segment_search.txt"},
		{problems.SmokeBasin{}, "inputs/smoke_basin.txt"},
		{problems.SyntaxScoring{}, "inputs/syntax_scoring.txt"},
		{problems.DumboOctopus{}, "inputs/dumbo_octopus.txt"},
		{problems.PassagePathing{}, "inputs/passage_pathing.txt"},
	}
	for index, runner := range runners {
		runner.PrintResults(index + 1)
	}
}
