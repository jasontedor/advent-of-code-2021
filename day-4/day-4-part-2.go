package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}

	draws, boards := ParseDrawsAndBoards(os.Args[1])

	for _, draw := range draws {
		for _, board := range boards {
			MarkDrawOnBoard(board, draw)
		}

		if len(boards) > 1 {
			// until there is one board left, filter out the winning boards
			boards = pruneWinningBoards(boards)
		} else {
			if CheckBoard(*boards[0]) {
				// we found the last winning board
				fmt.Println(Score(*boards[0]) * draw)
				break
			}
		}
	}
}

func pruneWinningBoards(boards []*Board) []*Board {
	var b []*Board
	for _, board := range boards {
		if !CheckBoard(*board) {
			b = append(b, board)
		}
	}
	return b
}
