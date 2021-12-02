package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
}

type AimedPosition struct {
	Horizontal int
	Depth      int
	Aim        int
}

type Instruction struct {
	direction string
	magnitude int
}

func (pos Position) Move(instr Instruction) Position {
	switch instr.direction {
	case "forward":
		return Position{
			Horizontal: pos.Horizontal + instr.magnitude,
			Depth:      pos.Depth,
		}
	case "up":
		return Position{
			Horizontal: pos.Horizontal,
			Depth:      pos.Depth - instr.magnitude,
		}
	case "down":
		return Position{
			Horizontal: pos.Horizontal,
			Depth:      pos.Depth + instr.magnitude,
		}
	default:
		panic(fmt.Sprint("cannot move in direction", instr.direction))
	}
}

func (pos AimedPosition) Move(instr Instruction) AimedPosition {
	switch instr.direction {
	case "forward":
		return AimedPosition{
			Horizontal: pos.Horizontal + instr.magnitude,
			Depth:      pos.Depth + (pos.Aim * instr.magnitude),
			Aim:        pos.Aim,
		}
	case "up":
		return AimedPosition{
			Horizontal: pos.Horizontal,
			Depth:      pos.Depth,
			Aim:        pos.Aim - instr.magnitude,
		}
	case "down":
		return AimedPosition{
			Horizontal: pos.Horizontal,
			Depth:      pos.Depth,
			Aim:        pos.Aim + instr.magnitude,
		}
	default:
		panic(fmt.Sprint("cannot move in direction", instr.direction))
	}
}

func NewInstruction(s string) Instruction {
	fields := strings.Fields(s)

	direction := fields[0]
	magnitude, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err.Error())
	}

	return Instruction{direction, magnitude}
}
