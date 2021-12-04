package problems

import (
	"log"
	"math"
	"strconv"
)

type BinaryDiagnostic struct{}

func (BinaryDiagnostic) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	mostCommonBits := getMostCommonBits(input)
	leastCommonBits := invertBitRepresentation(mostCommonBits)
	gamma := convertBitsToInt(mostCommonBits)
	epsilon := convertBitsToInt(leastCommonBits)
	return gamma * epsilon
}

func (BinaryDiagnostic) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	itemWithMostCommonBits := filterItemsByBitPrevalence(input, true, 0)
	itemWithLeastCommonBits := filterItemsByBitPrevalence(input, false, 0)
	O2GeneratorRating := convertBitLikeStringToInt(itemWithMostCommonBits)
	CO2ScrubberRating := convertBitLikeStringToInt(itemWithLeastCommonBits)
	return O2GeneratorRating * CO2ScrubberRating
}

func filterItemsByBitPrevalence(items []string, useMostCommon bool, index int) string {
	if len(items) == 1 {
		return items[0]
	}
	withZero, withOne := partitionItemsByBitAtIndex(items, index)
	zeroIsMoreCommon := len(withZero) > len(withOne)
	if useMostCommon == zeroIsMoreCommon {
		return filterItemsByBitPrevalence(withZero, useMostCommon, index+1)
	} else {
		return filterItemsByBitPrevalence(withOne, useMostCommon, index+1)
	}
}

func partitionItemsByBitAtIndex(items []string, index int) (withZero []string, withOne []string) {
	for line, item := range items {
		if item[index] == '0' {
			withZero = append(withZero, item)
		} else if item[index] == '1' {
			withOne = append(withOne, item)
		} else {
			log.Fatalf("cannot parse line %d at index %d", line, index)
		}
	}
	return
}

func getMostCommonBits(input []string) (bits []bool) {
	zeroes, ones := countZeroesAndOnes(input)
	for idx, countZero := range zeroes {
		countOne := ones[idx]
		if countZero < countOne {
			bits = append(bits, true)
		} else if countZero > countOne {
			bits = append(bits, false)
		} else {
			log.Fatalf("equal number of '0' and '1' in column %d", idx)
		}
	}
	return
}

func countZeroesAndOnes(input []string) ([]int, []int) {
	digits := len(input[0])
	zeroes := make([]int, digits)
	ones := make([]int, digits)
	for _, str := range input {
		bits := convertBitLikeStringToBits(str)
		for idx, isOne := range bits {
			if isOne {
				zeroes[idx]++
			} else {
				ones[idx]++
			}
		}
	}
	return zeroes, ones
}

func invertBitRepresentation(bits []bool) (inverted []bool) {
	for _, isOne := range bits {
		inverted = append(inverted, !isOne)
	}
	return
}

func convertBitsToInt(bits []bool) (number int) {
	maxExponent := len(bits) - 1
	for idx, isOne := range bits {
		if !isOne {
			continue
		}
		exponent := float64(maxExponent - idx)
		value := math.Pow(2, exponent)
		number += int(value)
	}
	return
}

func convertBitLikeStringToBits(str string) (bits []bool) {
	for idx, chr := range str {
		if chr == '1' {
			bits = append(bits, true)
		} else if chr == '0' {
			bits = append(bits, false)
		} else {
			log.Fatalf("cannot parse '%s' at index %d", str, idx)
		}
	}
	return
}

func convertBitLikeStringToInt(str string) int {
	number, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatalf("cannot parse %s as a bit-like string", str)
	}
	return int(number)
}
