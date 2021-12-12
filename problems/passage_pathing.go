package problems

import "strings"

type PassagePathing struct{}

func (PassagePathing) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	caves := getCaves(input)
	return caves.getPaths("start", []string{}, false)
}

func (PassagePathing) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	caves := getCaves(input)
	return caves.getPaths("start", []string{}, true)
}

type stringset map[string]bool
type caves map[string]stringset

func (c caves) isSmall(label string) bool {
	return label == strings.ToLower(label)
}

func (c caves) createIfNotExists(label string) {
	_, exists := c[label]
	if exists {
		return
	}
	c[label] = make(stringset)
}

func (c caves) connect(label1, label2 string) {
	c[label1][label2] = true
	c[label2][label1] = true
}

func (c caves) getPaths(label string, avoid []string, revisit bool) (n int) {
	if label == "end" {
		return 1
	}
	if c.isSmall(label) {
		avoid = append(avoid, label)
	}
	for connection := range c[label] {
		if contains(avoid, connection) {
			if !revisit || connection == "start" {
				continue
			}
			n += c.getPaths(connection, avoid, false)
		} else {
			n += c.getPaths(connection, avoid, revisit)
		}
	}
	return
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if str == item {
			return true
		}
	}
	return false
}

func getCaves(lines []string) caves {
	caves := make(caves)
	for _, line := range lines {
		labels := strings.Split(line, "-")
		caves.createIfNotExists(labels[0])
		caves.createIfNotExists(labels[1])
		caves.connect(labels[0], labels[1])
	}
	return caves
}
