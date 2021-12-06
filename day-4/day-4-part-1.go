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
	boards, draws := ParseBoardsAndDraws(os.Args[1])

	for _, draw := range draws {
		for _, board := range boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.squares[i][j].number == draw {
						board.squares[i][j].marked = true
					}
				}
			}
		}

		if b := maybeFirstWinningBoard(boards); b != nil {
			fmt.Printf("%d\n", Score(*b)*draw)
			break
		}
	}
}

func maybeFirstWinningBoard(boards []*Board) *Board {
	for _, board := range boards {
		if CheckBoard(*board) {
			return board
		}
	}
	return nil
}
