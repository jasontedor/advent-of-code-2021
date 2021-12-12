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
			for _, n := range path {
				// the small cave is already in the path
				if n == c {
					return false
				}
			}
			// the small cave is not already in the path
			return true
		})
	fmt.Println(len(paths))
}
