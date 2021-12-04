package problems

import (
	"fmt"
	"math"
	"strconv"
)

type BinaryDiagnostic struct{}

func (BinaryDiagnostic) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	return getPowerConsumption(input)
}

func (BinaryDiagnostic) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	return getLifeSupportRating(input)
}

func getLifeSupportRating(input []string) int {
	common, uncommon := splitItemsRecursivelyIntoCommonAndUncommon(input, 0)
	O2GeneratorRating, _ := strconv.ParseInt(common[0], 2, 64)
	CO2ScrubberRating, _ := strconv.ParseInt(uncommon[0], 2, 64)
	return int(O2GeneratorRating * CO2ScrubberRating)
}

func splitItemsRecursivelyIntoCommonAndUncommon(items []string, index int) ([]string, []string) {
	common, uncommon := splitItemsIntoCommonAndUncommon(items, index)
	if len(common) > 1 {
		common, _ = splitItemsRecursivelyIntoCommonAndUncommon(common, index+1)
	}
	if len(uncommon) > 1 {
		_, uncommon = splitItemsRecursivelyIntoCommonAndUncommon(uncommon, index+1)
	}
	return common, uncommon
}

func splitItemsIntoCommonAndUncommon(items []string, index int) ([]string, []string) {
	hasZero, hasOne := splitItemsByBitAtIndex(items, index)
	if len(hasZero) > len(hasOne) {
		return hasZero, hasOne
	} else {
		return hasOne, hasZero
	}
}

func splitItemsByBitAtIndex(items []string, index int) ([]string, []string) {
	var hasZero []string
	var hasOne []string
	for _, item := range items {
		if item[index] == '0' {
			hasZero = append(hasZero, item)
		} else {
			hasOne = append(hasOne, item)
		}
	}
	return hasZero, hasOne
}

func getPowerConsumption(input []string) int {
	bitCount := countBitsColumnWise(input)
	gamma, epsilon := getGammaAndEpsilonRates(bitCount)
	return gamma * epsilon
}

func getGammaAndEpsilonRates(bitCount []int) (int, int) {
	gamma, epsilon := 0, 0
	significantDigit := len(bitCount) - 1
	for idx, count := range bitCount {
		exponent := float64(significantDigit - idx)
		bitValue := int(math.Pow(2, exponent))
		if count > 0 {
			gamma += bitValue
		} else if count < 0 {
			epsilon += bitValue
		}
	}
	return gamma, epsilon
}

func countBitsColumnWise(input []string) []int {
	figures := len(input[0])
	bitCount := make([]int, figures)
	for ln, str := range input {
		for idx, chr := range str {
			if chr == '1' {
				bitCount[idx]++
			} else if chr == '0' {
				bitCount[idx]--
			} else {
				panic(fmt.Sprintf("cannot read line %d index %d", ln, idx))
			}
		}
	}
	return bitCount
}
