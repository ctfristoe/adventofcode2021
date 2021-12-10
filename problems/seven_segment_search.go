package problems

import (
	"strings"
)

type SevenSegmentSearch struct{}

func (SevenSegmentSearch) DoProblemOne(filepath string) (n int) {
	input := ReadLines(filepath)
	for _, str := range input {
		n += countOccurrancesOfSpecialSevenSegmentNumbers(str)
	}
	return
}

func (SevenSegmentSearch) DoProblemTwo(filepath string) int {
	return 0
}

func countOccurrancesOfSpecialSevenSegmentNumbers(str string) (n int) {
	divided := strings.Split(str, " | ")
	outputs := strings.Fields(divided[1])
	for _, item := range outputs {
		switch len(item) {
		case 2, 3, 4, 7:
			n++
		default:
			continue
		}
	}
	return
}
