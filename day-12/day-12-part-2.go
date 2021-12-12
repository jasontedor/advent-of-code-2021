package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}
	adjacencyLists := ParseAdjacencyLists(os.Args[1])

	path := make([]string, 0)
	paths := make([][]string, 0)
	FindPaths(
		adjacencyLists,
		"start",
		"end",
		&path,
		&paths,
		func(path []string, c string) bool {
			if c != strings.ToLower(c) {
				// not a small cave
				return true
			}
			counts := make(map[string]int)
			for _, n := range path {
				if n == strings.ToLower(n) {
					counts[n]++
				}
			}
			oneSmallCaveHasBeenVisitedTwice := false
			for _, v := range counts {
				if v > 1 {
					oneSmallCaveHasBeenVisitedTwice = true
				}
			}
			for _, n := range path {
				// the small cave is already in the path
				if n == c {
					if c == "start" {
						return false
					}
					if c == "end" {
						return false
					}
					if oneSmallCaveHasBeenVisitedTwice {
						return false
					} else {
						// we can accept the small cave a second time
						return true
					}
				}

			}
			// the small cave is not already in the path
			return true
		})

	fmt.Println(len(paths))
}
