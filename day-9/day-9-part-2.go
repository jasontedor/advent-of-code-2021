package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Basin struct {
	locations []Location
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}

	grid := ParseGrid(os.Args[1])
	lowPoints := LowPoints(grid)

	lowPointToBasin := make(map[Location][]Location)

	for _, lowPoint := range lowPoints {
		lowPointToBasin[lowPoint] = make([]Location, 0)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 9 is not in any basin
			if grid[i][j] == 9 {
				continue
			}
			// find the low point of the basin containing the given point
			lowPoint := flow(Location{grid[i][j], i, j}, grid)
			lowPointToBasin[lowPoint] =
				append(lowPointToBasin[lowPoint], Location{grid[i][j], i, j})
		}
	}

	var basins [][]Location
	for _, v := range lowPointToBasin {
		basins = append(basins, v)
	}

	// sort the basins by length
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	// we need the lengths of the largest three basins
	fmt.Println(len(basins[0]) * len(basins[1]) * len(basins[2]))
}

func flow(start Location, grid [][]int) Location {
	current := start
	// flow to a lower point until we flow no more
	for {
		a := Above(current.i, current.j, grid)
		b := Below(current.i, current.j, grid)
		l := Left(current.i, current.j, grid)
		r := Right(current.i, current.j, grid)
		// it does not matter if there are multiple points we could flow to
		if a != nil && a.height < current.height {
			current = *a
		} else if b != nil && b.height < current.height {
			current = *b
		} else if l != nil && l.height < current.height {
			current = *l
		} else if r != nil && r.height < current.height {
			current = *r
		} else {
			break
		}
	}

	// this is the low point of the basin containing the start point
	return current
}
