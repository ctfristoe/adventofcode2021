package problems

import (
	"fmt"
	"strconv"
	"strings"
)

type Depth struct{}

func (Depth) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	finalposition := getFinalPositionFromInstructions(input)
	return finalposition.horizontal * finalposition.depth
}

func (Depth) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	finalposition := getFinalPositionFromInstructionsWithAim(input)
	return finalposition.horizontal * finalposition.depth
}

func getFinalPositionFromInstructions(input []string) position {
	position := position{horizontal: 0, depth: 0}
	for _, line := range input {
		instr := makeInstruction(line)
		position = position.move(instr)
	}
	return position
}

func getFinalPositionFromInstructionsWithAim(input []string) positionWithAim {
	position := positionWithAim{horizontal: 0, depth: 0, aim: 0}
	for _, line := range input {
		instr := makeInstruction(line)
		position = position.move(instr)
	}
	return position
}

type instruction struct {
	direction string
	magnitude int
}

func makeInstruction(s string) instruction {
	fields := strings.Fields(s)

	direction := fields[0]
	magnitude, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err.Error())
	}

	return instruction{direction, magnitude}
}

type position struct {
	horizontal int
	depth      int
}

func (pos position) move(instr instruction) position {
	switch instr.direction {
	case "forward":
		return position{
			horizontal: pos.horizontal + instr.magnitude,
			depth:      pos.depth,
		}
	case "up":
		return position{
			horizontal: pos.horizontal,
			depth:      pos.depth - instr.magnitude,
		}
	case "down":
		return position{
			horizontal: pos.horizontal,
			depth:      pos.depth + instr.magnitude,
		}
	default:
		panic(fmt.Sprint("cannot move in direction", instr.direction))
	}
}

type positionWithAim struct {
	horizontal int
	depth      int
	aim        int
}

func (pos positionWithAim) move(instr instruction) positionWithAim {
	switch instr.direction {
	case "forward":
		return positionWithAim{
			horizontal: pos.horizontal + instr.magnitude,
			depth:      pos.depth + (pos.aim * instr.magnitude),
			aim:        pos.aim,
		}
	case "up":
		return positionWithAim{
			horizontal: pos.horizontal,
			depth:      pos.depth,
			aim:        pos.aim - instr.magnitude,
		}
	case "down":
		return positionWithAim{
			horizontal: pos.horizontal,
			depth:      pos.depth,
			aim:        pos.aim + instr.magnitude,
		}
	default:
		panic(fmt.Sprint("cannot move in direction", instr.direction))
	}
}
