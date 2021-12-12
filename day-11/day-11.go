package main

import (
	"bufio"
	"log"
	"os"
)

type Location struct {
	i int
	j int
}

func ParseGrid(path string) *[10][10]int {
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

	var grid [10][10]int
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i >= 10 {
			log.Fatalf("expected ten lines, found additional line: %s", line)
		}
		if len(line) != 10 {
			log.Fatalf(
				"expected input to have length 10 but was %d: %s",
				len(line),
				line)
		}
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
		i++
	}

	return &grid
}

func Step(grid *[10][10]int) int {
	flashes := make(map[Location]bool) // locations that flashed this step
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j]++
		}
	}

	for {
		var queue []Location
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j] > 9 {
					flashes[Location{i, j}] = true
					grid[i][j] = 0
					queue = append(queue, Location{i, j})
				}
			}
		}

		if len(queue) == 0 {
			break
		}

		for _, l := range queue {
			if l.i-1 >= 0 {
				if l.j-1 >= 0 && !flashes[Location{l.i - 1, l.j - 1}] {
					grid[l.i-1][l.j-1]++
				}
				if !flashes[Location{l.i - 1, l.j}] {
					grid[l.i-1][l.j]++
				}
				if l.j+1 < 10 && !flashes[Location{l.i - 1, l.j + 1}] {
					grid[l.i-1][l.j+1]++
				}
			}
			if l.j-1 >= 0 && !flashes[Location{l.i, l.j - 1}] {
				grid[l.i][l.j-1]++
			}
			if l.j+1 < 10 && !flashes[Location{l.i, l.j + 1}] {
				grid[l.i][l.j+1]++
			}
			if l.i+1 < 10 {
				if l.j-1 >= 0 && !flashes[Location{l.i + 1, l.j - 1}] {
					grid[l.i+1][l.j-1]++
				}
				if !flashes[Location{l.i + 1, l.j}] {
					grid[l.i+1][l.j]++
				}
				if l.j+1 < 10 && !flashes[Location{l.i + 1, l.j + 1}] {
					grid[l.i+1][l.j+1]++
				}
			}
		}
		queue = make([]Location, 0)
	}

	return len(flashes)
}
