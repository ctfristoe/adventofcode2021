package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadIntegerLines(filepath string) []int {
	strLines := ReadLines(filepath)
	var intLines []int

	for _, line := range strLines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}
		intLines = append(intLines, number)
	}

	return intLines
}

func ReadIntegerMatrix(filepath string) [][]int {
	lines := ReadLines(filepath)
	numbers := make([][]int, len(lines))
	for y, str := range lines {
		numbers[y] = make([]int, len(str))
		for x, r := range []rune(str) {
			numbers[y][x] = int(r - '0')
		}
	}
	return numbers
}

func ReadLines(filepath string) []string {
	lines, err := tryReadLines(filepath)
	if err != nil {
		panic(err.Error())
	}
	return lines
}

func ReadCommaSeparatedIntegers(filepath string) (numbers []int) {
	lines, err := tryReadLines(filepath)
	if err != nil {
		panic(err.Error())
	}
	for _, str := range strings.Split(lines[0], ",") {
		number := parseInteger(str)
		numbers = append(numbers, number)
	}
	return
}

func tryReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
