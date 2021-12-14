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

	riskLevels, rows, columns := ParseRiskLevels(os.Args[1])

	for i := 0; i < 5*rows; i++ {
		for j := 0; j < 5*columns; j++ {
			row := i / rows
			column := j / columns
			value := riskLevels[Position{i % rows, j % columns}]
			adjustedValue := value + (row + column)
			if adjustedValue > 9 {
				adjustedValue = adjustedValue%10 + 1
			}
			riskLevels[Position{i, j}] = adjustedValue
		}
	}

	rows = 5 * rows
	columns = 5 * columns

	fmt.Println(ShortestPath(riskLevels, rows, columns))
}
