package main

type Square struct {
	number int
	marked bool
}

type Board struct {
	squares [5][5]Square
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
