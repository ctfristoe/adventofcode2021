package main

import (
	"fmt"

	"github.com/ctfristoe/adventofcode/problems"
	"github.com/ctfristoe/adventofcode/utils"
)

func main() {
	printDayOneProblems()
	printDayTwoProblems()
}

func printDayOneProblems() {
	input := utils.ReadIntegerLines("inputs/sonar_sweep.txt")
	answer1 := problems.CountDepthIncreases(input)
	answer2 := problems.CountSlidingWindowIncreases(input)
	fmt.Println("Day 1 Problem 1:", answer1)
	fmt.Println("Day 1 Problem 2:", answer2)
	fmt.Println()
}

func printDayTwoProblems() {
	input := utils.ReadLines("inputs/depth.txt")
	answer1 := problems.GetFinalPosition(input)
	answer2 := problems.GetFinalPositionUsingAim(input)
	fmt.Println("Day 2 Problem 1:", answer1.Horizontal*answer1.Depth)
	fmt.Println("Day 2 Problem 2:", answer2.Horizontal*answer2.Depth)
	fmt.Println()
}
