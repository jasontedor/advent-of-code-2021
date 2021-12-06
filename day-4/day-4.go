package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	number int
	marked bool
}

type Board struct {
	squares [5][5]Square
}

func ParseBoardsAndDraws(path string) ([]*Board, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var boards []*Board
	var draws []int
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("expected draws")
	}
	rawDraws := strings.Split(scanner.Text(), ",")
	draws = make([]int, len(rawDraws))
	for i, rawDraw := range rawDraws {
		draws[i], err = strconv.Atoi(rawDraw)
		if err != nil {
			log.Fatal(err)
		}
	}
	for scanner.Scan() {
		var board Board
		for i := 0; i < 5; i++ {
			if !scanner.Scan() {
				log.Fatal("expected next board line")
			}
			rawNumbers := strings.Fields(scanner.Text())
			if len(rawNumbers) != 5 {
				log.Fatalf(
					"expected board line %d: %s to contain 5 elements but was %d",
					i,
					rawNumbers,
					len(rawNumbers))
			}
			for j, rawNumber := range rawNumbers {
				n, err := strconv.Atoi(rawNumber)
				if err != nil {
					log.Fatal(err)
				}
				board.squares[i][j] = Square{n, false}
			}
		}
		boards = append(boards, &board)
	}
	return boards, draws
}

func MarkDrawOnBoard(board *Board, draw int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.squares[i][j].number == draw {
				board.squares[i][j].marked = true
			}
		}
	}
}

func CheckBoard(board Board) bool {
	for row := 0; row < 5; row++ {
		if CheckRow(board, row) {
			return true
		}
	}
	for column := 0; column < 5; column++ {
		if CheckColumn(board, column) {
			return true
		}
	}
	return false
}

func CheckRow(board Board, row int) bool {
	complete := true
	for column := 0; column < 5; column++ {
		complete = complete && board.squares[row][column].marked
		if !complete {
			break
		}
	}
	return complete
}

func CheckColumn(board Board, column int) bool {
	complete := true
	for row := 0; row < 5; row++ {
		complete = complete && board.squares[row][column].marked
		if !complete {
			break
		}
	}
	return complete
}

func Score(board Board) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.squares[i][j].marked {
				sum += board.squares[i][j].number
			}
		}
	}
	return sum
}
