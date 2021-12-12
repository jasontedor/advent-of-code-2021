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

	grid := ParseGrid(os.Args[1])

	steps := 0
	for {
		steps++
		if Step(grid) == 100 {
			break
		}
	}

	fmt.Println(steps)
}
