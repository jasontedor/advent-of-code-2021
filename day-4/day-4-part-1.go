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
			MarkDrawOnBoard(board, draw)
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
