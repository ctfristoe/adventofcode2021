package problems

import (
	"sort"
)

type SyntaxScoring struct{}

var theProblem1Scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var theProblem2Scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func (SyntaxScoring) DoProblemOne(filepath string) (score int) {
	input := ReadLines(filepath)
	for _, line := range input {
		char := getFirstInvalidCharacter(line)
		if char != rune(0) {
			score += theProblem1Scores[char]
		}
	}
	return
}

func (SyntaxScoring) DoProblemTwo(filepath string) (score int) {
	input := ReadLines(filepath)
	return getMedianAutocompleteScore(input)
}

type lookup struct {
	open  map[rune]rune
	close map[rune]rune
}

func (l lookup) get(char rune) (pair rune, open bool) {
	pair, open = l.open[char]
	if !open {
		pair = l.close[char]
	}
	return
}

var theLookup = lookup{
	open: map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	},
	close: map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	},
}

type stack struct {
	inner []rune
}

func (s *stack) add(char rune) {
	s.inner = append(s.inner, char)
}

func (s *stack) pop() rune {
	size := len(s.inner)
	top := s.inner[size-1]
	s.inner = s.inner[:size-1]
	return top
}

func (s *stack) isEmpty() bool {
	return len(s.inner) == 0
}

func getMedianAutocompleteScore(lines []string) int {
	var scores []int
	for _, line := range lines {
		autocompleted := getAutoCompletedPortion(line)
		score := scoreAutoCompletedPortion(autocompleted)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func scoreAutoCompletedPortion(str string) int {
	score := 0
	for _, char := range str {
		score *= 5
		score += theProblem2Scores[char]
	}
	return score
}

func getAutoCompletedPortion(line string) string {
	stack := stack{inner: make([]rune, 0)}
	var runes []rune
	for _, char := range line {
		pair, open := theLookup.get(char)
		if open {
			stack.add(char)
		} else {
			top := stack.pop()
			if top != pair {
				return "" // invalid
			}
		}
	}
	for !stack.isEmpty() {
		top := stack.pop()
		autoclose, _ := theLookup.get(top)
		runes = append(runes, autoclose)
	}
	return string(runes)
}

func getFirstInvalidCharacter(line string) rune {
	stack := stack{inner: make([]rune, 0)}
	for _, char := range line {
		pair, open := theLookup.get(char)
		if open {
			stack.add(char)
		} else {
			top := stack.pop()
			if top != pair {
				return char
			}
		}
	}
	return rune(0)
}
