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

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += Step(grid)
	}

	fmt.Println(flashes)
}
