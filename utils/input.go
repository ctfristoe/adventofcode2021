package utils

import (
	"bufio"
	"os"
	"strconv"
)

func tryReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
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

func ReadLines(filename string) []string {
	lines, err := tryReadLines(filename)
	if err != nil {
		panic(err.Error())
	}
	return lines
}

func ReadIntegerLines(filename string) []int {
	strLines := ReadLines(filename)
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
