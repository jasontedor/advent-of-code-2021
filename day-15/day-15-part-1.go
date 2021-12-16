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

	fmt.Println(ShortestPath(riskLevels, rows, columns))
}
