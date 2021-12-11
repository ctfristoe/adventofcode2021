package problems

import (
	"sort"
	"strings"
)

const DIGITS = "0123456789"
const LETTERS = "abcdefg"

var D0 = MakeRuneSet("abcefg")
var D1 = MakeRuneSet("cf")
var D2 = MakeRuneSet("acdeg")
var D3 = MakeRuneSet("acdfg")
var D4 = MakeRuneSet("bcdf")
var D5 = MakeRuneSet("abdfg")
var D6 = MakeRuneSet("abdefg")
var D7 = MakeRuneSet("acf")
var D8 = MakeRuneSet("abcdefg")
var D9 = MakeRuneSet("abcdfg")

type SevenSegmentSearch struct{}

func (SevenSegmentSearch) DoProblemOne(filepath string) (n int) {
	input := ReadLines(filepath)
	for _, str := range input {
		_, outputs := getSignalsAndOutputs(str)
		n += countUniqueLengthSignals(outputs)
	}
	return
}

func (SevenSegmentSearch) DoProblemTwo(filepath string) (n int) {
	for _, line := range ReadLines(filepath) {
		signals, outputs := getSignalsAndOutputs(line)
		unscrambler := getSignalUnscrambler(signals)
		translator := makeTranslator(unscrambler)
		n += getOutputValue(outputs, translator)
	}
	return
}

func getOutputValue(outputs []string, transl translator) (n int) {
	multiplier := 1000
	for _, output := range outputs {
		alphabetized := getAlphabetizedString(output)
		number := transl[alphabetized]
		n += (multiplier * number)
		multiplier /= 10
	}
	return
}

func getSignalUnscrambler(signals []string) unscrambler {
	RuneSets := convertStringsToSortedSets(signals)
	unscr := makePossibilitiesTracker()
	// 1 (2 digits)
	standardOne, scrambledOne := D1, RuneSets[0]
	unscr.prunePossibilities(standardOne, scrambledOne)
	// 7 (3 digits)
	standardSeven, scrambledSeven := D7, RuneSets[1]
	unscr.prunePossibilities(standardSeven, scrambledSeven)
	// 4 (4 digits)
	standardFour, scrambledFour := D4, RuneSets[2]
	unscr.prunePossibilities(standardFour, scrambledFour)

	for _, scrambledTwoThreeOrFive := range RuneSets[3:6] {
		if scrambledOne.issubRuneSet(scrambledTwoThreeOrFive) {
			// only three contains the same signals as one
			unscr.prunePossibilities(D3, scrambledTwoThreeOrFive)
		}
	}

	for _, scrambledZeroSixOrNine := range RuneSets[6:9] {
		if !scrambledOne.issubRuneSet(scrambledZeroSixOrNine) {
			// only six does not contain the same signals as one
			unscr.prunePossibilities(D6, scrambledZeroSixOrNine)
		}
	}
	return unscr
}

func countUniqueLengthSignals(signals []string) (n int) {
	for _, item := range signals {
		switch len(item) {
		case 2, 3, 4, 7:
			n++
		default:
			continue
		}
	}
	return
}

type translator map[string]int

func makeTranslator(unscr unscrambler) translator {
	digits := []RuneSet{D0, D1, D2, D3, D4, D5, D6, D7, D8, D9}
	translator := make(translator)
	for n, digit := range digits {
		scrambled := make(RuneSet, len(digit))
		for key := range digit {
			scrambled.add(unscr[key].first())
		}
		translator[scrambled.toString()] = n
	}
	return translator
}

/// mapping of digits to possible values when unscrambled
///  * add, contains, first, difference, intersection
type unscrambler map[rune]RuneSet

/// given a standard signal like (c, f) [one] and a scrambled signal like (a, g), limits possibilities
/// accordingly. For instance, if we know that (c, f) scrambled is (a, g):
///   * both c and f must be either a or g, AND
///   * a and g cannot be any letter besides c or f
func (unscr unscrambler) prunePossibilities(standard, scrambled RuneSet) {
	for char, possibilities := range unscr {
		if standard.contains(char) {
			unscr[char] = possibilities.intersection(scrambled)
		} else {
			unscr[char] = possibilities.difference(scrambled)
		}
	}
}

/// returns a new possibilies struct where every letter is possible for every letter
func makePossibilitiesTracker() unscrambler {
	new := make(unscrambler, 8)
	for _, digit := range LETTERS {
		new[digit] = MakeRuneSet(LETTERS)
	}
	return new
}

/// for parsing inputs into a slice of RuneSets
func convertStringsToSortedSets(slice []string) []RuneSet {
	sortByLength(slice)
	new := make([]RuneSet, len(slice))
	for i, str := range slice {
		new[i] = MakeRuneSet(str)
	}
	return new
}

/// sorting helpers
func sortByLength(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})
}

func getAlphabetizedString(str string) string {
	runes := []rune(str)
	alphabetizeRunes(runes)
	return string(runes)
}

func alphabetizeRunes(slice []rune) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

/// input parsing helper, splits the line along pipe and into fields
func getSignalsAndOutputs(str string) (signals, outputs []string) {
	divided := strings.Split(str, " | ")
	signals = strings.Fields(divided[0])
	outputs = strings.Fields(divided[1])
	return
}
