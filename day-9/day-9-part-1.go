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

	lowPoints := LowPoints(ParseGrid(os.Args[1]))

	risk := 0
	for i := 0; i < len(lowPoints); i++ {
		risk += 1 + lowPoints[i].height
	}

	fmt.Println(risk)
}
