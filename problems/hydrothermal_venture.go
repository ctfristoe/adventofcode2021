package problems

import (
	"fmt"
	"strconv"
	"strings"
)

type HydrothermalVenture struct{}

func (HydrothermalVenture) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	lines := getLinesFromPuzzleInput(input)
	m := mapLines(lines, true)
	return countOverlappingLines(m)
}

func (HydrothermalVenture) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	lines := getLinesFromPuzzleInput(input)
	m := mapLines(lines, false)
	return countOverlappingLines(m)
}

type coordinate struct {
	x int
	y int
}

type line struct {
	start coordinate
	end   coordinate
}

func (line line) isDiagonal() bool {
	return !line.isHorizontal() && !line.isVertical()
}

func (line line) isHorizontal() bool {
	return line.start.y == line.end.y
}

func (line line) isVertical() bool {
	return line.start.x == line.end.x
}

func (line line) getPoints() []coordinate {
	if line.isHorizontal() {
		return generateHorizontalPoints(line.start.y, line.start.x, line.end.x)
	} else if line.isVertical() {
		return generateVerticalPoints(line.start.x, line.start.y, line.end.y)
	} else {
		return generateDiagonalPoints(line.start.x, line.start.y, line.end.x, line.end.y)
	}
}

func countOverlappingLines(m map[coordinate]int) (n int) {
	for _, overlapping := range m {
		if overlapping > 1 {
			n++
		}
	}
	return
}

func mapLines(lines []line, ignoreDiagonal bool) (m map[coordinate]int) {
	m = make(map[coordinate]int)
	for _, line := range lines {
		if ignoreDiagonal && line.isDiagonal() {
			continue
		}
		for _, coord := range line.getPoints() {
			m[coord]++
		}
	}
	return
}

func generateDiagonalPoints(x0 int, y0 int, x1 int, y1 int) (points []coordinate) {
	xmin, xmax := MinMax(x1, x0)
	ymin, ymax := MinMax(y1, y0)
	slope := (y1 - y0) / (x1 - x0)
	for i := 0; i <= (xmax - xmin); i++ {
		var x, y int
		x = xmin + i
		if slope < 0 {
			y = ymax - i
		} else {
			y = ymin + i
		}
		points = append(points, coordinate{x, y})
	}
	return
}

func generateHorizontalPoints(y int, x0 int, x1 int) (points []coordinate) {
	x0, x1 = MinMax(x0, x1)
	for i := x0; i <= x1; i++ {
		points = append(points, coordinate{x: i, y: y})
	}
	return
}

func generateVerticalPoints(x int, y0 int, y1 int) (points []coordinate) {
	y0, y1 = MinMax(y0, y1)
	for i := y0; i <= y1; i++ {
		points = append(points, coordinate{x: x, y: i})
	}
	return
}

func getLinesFromPuzzleInput(input []string) (lines []line) {
	for _, str := range input {
		lines = append(lines, getLineFromString(str))
	}
	return
}

func getLineFromString(input string) line {
	sl := strings.Split(input, " -> ")
	x0, y0 := getCoordinateFromString(sl[0])
	x1, y1 := getCoordinateFromString(sl[1])
	return line{
		start: coordinate{x: x0, y: y0},
		end:   coordinate{x: x1, y: y1},
	}
}

func getCoordinateFromString(str string) (x int, y int) {
	sl := strings.Split(str, ",")
	return getIntFromString(sl[0]), getIntFromString(sl[1])
}

func getIntFromString(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprint("cannot parse", str))
	}
	return n
}
