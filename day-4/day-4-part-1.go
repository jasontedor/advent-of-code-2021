package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}
	file, err := os.Open(os.Args[1])
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
