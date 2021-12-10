package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type Location struct {
	height int
	i      int
	j      int
}

func ParseGrid(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	length := 0
	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if length == 0 {
			length = len(line)
		}
		if length != len(line) {
			log.Fatalf(
				"expected length of line to be %d but was %d: %s",
				length,
				len(line),
				line)
		}
		match, err := regexp.MatchString("[0-9]+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match [0-9]+ but was %s", line)
		}
		var row = make([]int, length)
		for j, c := range line {
			row[j] = int(c - '0')
		}
		grid = append(grid, row)
	}

	return grid
}

func Above(i int, j int, grid [][]int) *Location {
	if i-1 >= 0 {
		return &Location{grid[i-1][j], i - 1, j}
	} else {
		return nil
	}
}

func Below(i int, j int, grid [][]int) *Location {
	if i+1 < len(grid) {
		return &Location{grid[i+1][j], i + 1, j}
	} else {
		return nil
	}
}

func Left(i int, j int, grid [][]int) *Location {
	if j-1 >= 0 {
		return &Location{grid[i][j-1], i, j - 1}
	} else {
		return nil
	}
}

func Right(i int, j int, grid [][]int) *Location {
	if j+1 < len(grid[i]) {
		return &Location{grid[i][j+1], i, j + 1}
	} else {
		return nil
	}
}

func LowPoints(grid [][]int) []Location {
	var locations []Location
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			a := Above(i, j, grid)
			b := Below(i, j, grid)
			l := Left(i, j, grid)
			r := Right(i, j, grid)
			if (a == nil || grid[i][j] < a.height) &&
				(b == nil || grid[i][j] < b.height) &&
				(l == nil || grid[i][j] < l.height) &&
				(r == nil || grid[i][j] < r.height) {
				locations = append(locations, Location{grid[i][j], i, j})
			}
		}
	}

	return locations
}
