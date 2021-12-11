package problems

import (
	"sort"
	"strings"
)

const DIGITS = "0123456789"
const LETTERS = "abcdefg"

var D0 = makeSet("abcefg")
var D1 = makeSet("cf")
var D2 = makeSet("acdeg")
var D3 = makeSet("acdfg")
var D4 = makeSet("bcdf")
var D5 = makeSet("abdfg")
var D6 = makeSet("abdefg")
var D7 = makeSet("acf")
var D8 = makeSet("abcdefg")
var D9 = makeSet("abcdfg")

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
	sets := convertStringsToSortedSets(signals)
	unscr := makePossibilitiesTracker()
	// 1 (2 digits)
	standardOne, scrambledOne := D1, sets[0]
	unscr.prunePossibilities(standardOne, scrambledOne)
	// 7 (3 digits)
	standardSeven, scrambledSeven := D7, sets[1]
	unscr.prunePossibilities(standardSeven, scrambledSeven)
	// 4 (4 digits)
	standardFour, scrambledFour := D4, sets[2]
	unscr.prunePossibilities(standardFour, scrambledFour)

	for _, scrambledTwoThreeOrFive := range sets[3:6] {
		if scrambledOne.issubset(scrambledTwoThreeOrFive) {
			// only three contains the same signals as one
			unscr.prunePossibilities(D3, scrambledTwoThreeOrFive)
		}
	}

	for _, scrambledZeroSixOrNine := range sets[6:9] {
		if !scrambledOne.issubset(scrambledZeroSixOrNine) {
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
	digits := []set{D0, D1, D2, D3, D4, D5, D6, D7, D8, D9}
	translator := make(translator)
	for n, digit := range digits {
		scrambled := make(set, len(digit))
		for key := range digit {
			scrambled.add(unscr[key].first())
		}
		translator[scrambled.toString()] = n
	}
	return translator
}

/// mapping of digits to possible values when unscrambled
///  * add, contains, first, difference, intersection
type unscrambler map[rune]set

/// given a standard signal like (c, f) [one] and a scrambled signal like (a, g), limits possibilities
/// accordingly. For instance, if we know that (c, f) scrambled is (a, g):
///   * both c and f must be either a or g, AND
///   * a and g cannot be any letter besides c or f
func (unscr unscrambler) prunePossibilities(standard, scrambled set) {
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
		new[digit] = makeSet(LETTERS)
	}
	return new
}

/// set structure, pretty standard, void should use no memory
/// methods: add, contains, first, difference, intersection
type void struct{}
type set map[rune]void

func (s set) add(item rune) {
	s[item] = void{}
}

func (s set) contains(item rune) bool {
	_, found := s[item]
	return found
}

func (s set) first() rune {
	for item := range s {
		return item
	}
	panic("empty set")
}

func (s set) difference(other set) set {
	new := make(set)
	for item := range s {
		if !other.contains(item) {
			new.add(item)
		}
	}
	return new
}

func (s set) intersection(other set) set {
	new := make(set)
	for item := range s {
		if other.contains(item) {
			new.add(item)
		}
	}
	return new
}

func (s set) issubset(other set) bool {
	for item := range s {
		if !other.contains(item) {
			return false
		}
	}
	return true
}

func (s set) toString() string {
	var runes []rune
	for item := range s {
		runes = append(runes, item)
	}
	alphabetizeRunes(runes)
	return string(runes)
}

/// for parsing inputs into a slice of sets
func convertStringsToSortedSets(slice []string) []set {
	sortByLength(slice)
	new := make([]set, len(slice))
	for i, str := range slice {
		new[i] = makeSet(str)
	}
	return new
}

/// set constructor from string
func makeSet(str string) set {
	new := make(set)
	for _, char := range str {
		new.add(char)
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
