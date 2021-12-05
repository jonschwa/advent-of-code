package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strings"
)

type NumberMap map[int]bool

type BoardRun struct {
	board         BoardNumbers
	numbersCalled NumberMap
	score         int
}
type BoardNumbers [][]int

func (b *BoardRun) CallNumber(number int) int {
	b.numbersCalled[number] = true
	numHitInRow := make(map[int]int)
	numHitInCol := make(map[int]int)
	for i, row := range b.board {
		for j, num := range row {
			if b.numbersCalled[num] {
				numHitInRow[i]++
				numHitInCol[j]++
			}
			// did we win?
			if numHitInCol[j] == 5 || numHitInRow[i] == 5 {
				b.score = b.Score(number)
				return b.score
			}
		}
	}
	return 0
}

func (b *BoardRun) Score(winningNumber int) int {
	sumUnMarked := 0
	for _, row := range b.board {
		for _, num := range row {
			if !b.numbersCalled[num] {
				sumUnMarked += num
			}
		}
	}
	return sumUnMarked * winningNumber
}

func (b BoardRun) Print() {
	for _, row := range b.board {
		for _, num := range row {
			fmt.Print(num, " ")
		}
		fmt.Print("\n")
	}
}

func RunDay4() {
	data := helpers.GetDataForDay(2021, 4)
	strNums := strings.Split(data[0], ",")
	numbers := helpers.ConvertStringSliceToIntSlice(strNums)
	boardData := data[2:]

	fmt.Println("Part 1:", solveDay4Part1(numbers, boardData))
	fmt.Println("Part 2:", solveDay4Part2(numbers, boardData))

}

func initBoards(boardData []string) []BoardNumbers {
	allBoards := []BoardNumbers{}
	boardRows := BoardNumbers{}
	for _, row := range boardData {
		intRow := helpers.ConvertStringSliceToIntSlice(strings.Fields(row))
		// empty rows between boards
		if len(intRow) == 0 {
			continue
		}
		boardRows = append(boardRows, intRow)
		if len(boardRows) == 5 {
			// add to the lsit of boards
			allBoards = append(allBoards, boardRows)
			// reset the rows
			boardRows = BoardNumbers{}
		}
	}
	return allBoards
}

func solveDay4Part1(numbers []int, boardData []string) int {
	// make the boards
	boards := initBoards(boardData)
	orderedWinningBoards := make(map[int]BoardRun)
	// want to go through and score each board to see which turn it wins on
	for _, board := range boards {
		boardRun := BoardRun{board, NumberMap{}, 0}
		for i, number := range numbers {
			score := boardRun.CallNumber(number)
			if score > 0 {
				// we have a winning board!
				orderedWinningBoards[i] = boardRun
				break
			}

		}
	}

	var firstWinner BoardRun
	minNumRounds := 0
	for i, val := range orderedWinningBoards {
		if minNumRounds == 0 || i < minNumRounds {
			minNumRounds = i
			firstWinner = val
		}
	}

	return firstWinner.score
}

func solveDay4Part2(numbers []int, boardData []string) int {
	// make the boards
	boards := initBoards(boardData)
	orderedWinningBoards := make(map[int]BoardRun)
	// want to go through and score each board to see which turn it wins on
	for _, board := range boards {
		boardRun := BoardRun{board, NumberMap{}, 0}
		for i, number := range numbers {
			score := boardRun.CallNumber(number)
			if score > 0 {
				// we have a winning board!
				orderedWinningBoards[i] = boardRun
				break
			}

		}
	}

	var lastWinner BoardRun
	maxNumRounds := 0
	for i, val := range orderedWinningBoards {
		if i > maxNumRounds {
			maxNumRounds = i
			lastWinner = val
		}
	}

	return lastWinner.score
}
