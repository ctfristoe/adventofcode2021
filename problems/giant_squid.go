package problems

import (
	"log"
	"strconv"
	"strings"
)

type GiantSquid struct{}

func (GiantSquid) DoProblemOne(filepath string) int {
	input := ReadLines(filepath)
	game := parseGame(input)
	return game.getWinningScore()
}

func (GiantSquid) DoProblemTwo(filepath string) int {
	input := ReadLines(filepath)
	game := parseGame(input)
	return game.getLosingScore()
}

type game struct {
	draws  []int
	boards []board
}

type board struct {
	marks [][]bool
	tiles [][]int
	rHits []int
	cHits []int
}

func (g game) getWinningScore() int {
	for _, number := range g.draws {
		for _, board := range g.boards {
			board.tryMark(number)
			if !board.hasBingo() {
				continue
			}
			sum := board.sumUnmarkedTiles()
			return number * sum
		}
	}
	panic("all numbers drawn, no one won")
}

func (g game) getLosingScore() int {
	for _, number := range g.draws {
		liveBoards := g.getLiveBoards()
		for _, board := range liveBoards {
			board.tryMark(number)
		}
		if len(liveBoards) == 1 {
			board := liveBoards[0]
			if board.hasBingo() {
				return number * board.sumUnmarkedTiles()
			}
		}
	}
	panic("all numbers drawn, no one won")
}

func (g game) getLiveBoards() (liveBoards []board) {
	for _, board := range g.boards {
		if !board.hasBingo() {
			liveBoards = append(liveBoards, board)
		}
	}
	return
}

func (b board) sumUnmarkedTiles() int {
	sum := 0
	for row, rowOfMarks := range b.marks {
		for col, isMarked := range rowOfMarks {
			if !isMarked {
				sum += b.tiles[row][col]
			}
		}
	}
	return sum
}

func (b board) hasBingo() bool {
	for i := 0; i < 5; i++ {
		if b.rHits[i] == 5 || b.cHits[i] == 5 {
			return true
		}
	}
	return false
}

func (b board) tryMark(value int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if value == b.tiles[row][col] {
				b.markHit(row, col)
			}
		}
	}
}

func (b board) markHit(row int, col int) {
	b.marks[row][col] = true
	b.rHits[row]++
	b.cHits[col]++
}

func parseGame(lines []string) game {
	draws := parseDraws(lines[0])
	var boards []board
	for i := 1; i < len(lines); i += 6 {
		chunk := lines[i+1 : i+6]
		board := parseBingoBoard(chunk)
		boards = append(boards, board)
	}
	return game{draws, boards}
}

func parseBingoBoard(lines []string) board {
	return board{
		marks: makeBlankMarks(),
		tiles: parseTiles(lines),
		rHits: make([]int, 5),
		cHits: make([]int, 5),
	}
}

func makeBlankMarks() (marks [][]bool) {
	for i := 0; i < 5; i++ {
		marks = append(marks, make([]bool, 5))
	}
	return
}

func parseTiles(lines []string) (tiles [][]int) {
	for _, line := range lines {
		var row []int
		for _, field := range strings.Fields(line) {
			row = append(row, parseInteger(field))
		}
		tiles = append(tiles, row)
	}
	return
}

func parseDraws(str string) (draws []int) {
	for _, strInt := range strings.Split(str, ",") {
		number := parseInteger(strInt)
		draws = append(draws, number)
	}
	return
}

func parseInteger(str string) int {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalf("cannot parse %s as integer", str)
	}
	return int(value)
}
