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

	FoldLocations(locations, folds[0])

	fmt.Println(len(locations))
}
