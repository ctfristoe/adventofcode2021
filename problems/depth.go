package problems

import "github.com/ctfristoe/adventofcode/utils"

func GetFinalPosition(input []string) utils.Position {
	position := utils.Position{Horizontal: 0, Depth: 0}
	for _, line := range input {
		instr := utils.NewInstruction(line)
		position = position.Move(instr)
	}
	return position
}

func GetFinalPositionUsingAim(input []string) utils.AimedPosition {
	position := utils.AimedPosition{Horizontal: 0, Depth: 0, Aim: 0}
	for _, line := range input {
		instr := utils.NewInstruction(line)
		position = position.Move(instr)
	}
	return position
}
