package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("expected 2 args, was %d", len(os.Args)-1)
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
	var costFunction func(moves int) int
	switch os.Args[2] {
	case "linear":
		costFunction = func(moves int) int {
			return moves
		}
	case "quadratic":
		costFunction = func(moves int) int {
			return moves * (moves + 1) / 2
		}
	default:
		log.Fatalf("unrecognized cost function %s", os.Args[2])
	}

	minPosition, maxPosition := math.MaxInt, math.MinInt
	countOfCrabWithPosition := make(map[int]int)
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("expected line")
	}
	rawPositions := strings.Split(scanner.Text(), ",")
	for _, raw := range rawPositions {
		position, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		countOfCrabWithPosition[position]++
		if position < minPosition {
			minPosition = position
		}
		if position > maxPosition {
			maxPosition = position
		}
	}
	if scanner.Scan() {
		log.Fatalf("expected end of file but was %s", scanner.Text())
	}

	minTotalCost := math.MaxInt
	for alignment := minPosition; alignment <= maxPosition; alignment++ {
		totalCost := 0
		for position, count := range countOfCrabWithPosition {
			if position < alignment {
				totalCost += count * costFunction(alignment-position)
			} else if position > alignment {
				totalCost += count * costFunction(position-alignment)
			}
		}
		if totalCost < minTotalCost {
			minTotalCost = totalCost
		}
	}

	fmt.Println(minTotalCost)
}
