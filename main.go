package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CountDepthIncreases(input []int) int {
	n := 0
	for index, item := range input[1:] {
		if item > input[index] {
			n++
		}
	}
	return n
}

func CountSlidingWindowIncreases(input []int) int {
	n := 0
	// current := input[0] + input[1] + input[2]
	for index, item := range input[3:] {
		if item > input[index] {
			n++
		}
	}
	return n
}

func ReadIntListFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, num)
	}
	return inputs, scanner.Err()
}

func main() {
	input, err := ReadIntListFromFile("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer1 := CountDepthIncreases(input)
	answer2 := CountSlidingWindowIncreases(input)
	fmt.Println("Day 1 Problem 1:", answer1)
	fmt.Println("Day 1 Problem 2:", answer2)
}
