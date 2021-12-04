package problems

import (
	"bufio"
	"os"
	"strconv"
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

func ReadLines(filepath string) []string {
	lines, err := tryReadLines(filepath)
	if err != nil {
		panic(err.Error())
	}
	return lines
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
