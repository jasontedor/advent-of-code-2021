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

	locations, folds := ParseLocationsAndFolds(os.Args[1])

	for _, fold := range folds {
		FoldLocations(locations, fold)
	}

	maxX, maxY := 0, 0
	for location, _ := range locations {
		if location.x > maxX {
			maxX = location.x
		}
		if location.y > maxY {
			maxY = location.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if locations[Location{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
