package main

import (
	"fmt"
	"log"
	"os"
)

func adjust(value int, i int) int {
	if value+i <= 9 {
		return value + i
	} else {
		return (value+i)%10 + 1
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}

	riskLevels, rows, columns := ParseRiskLevels(os.Args[1])

	fmt.Println(ShortestPath(riskLevels, rows, columns))
}
