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

	lines := ParseLines(os.Args[1])
	grid := MakeGrid(lines)

	for _, line := range lines {
		if isHorizontal(line) || isVertical(line) {
			grid[line.first.x][line.first.y]++
			current := NextPointOnLine(line.first, line.second)
			for current != line.second {
				grid[current.x][current.y]++
				current = NextPointOnLine(current, line.second)
			}
			grid[line.second.x][line.second.y]++
		}
	}

	fmt.Println(CountGrid(grid))
}

func isHorizontal(line Line) bool {
	return line.first.y == line.second.y
}

func isVertical(line Line) bool {
	return line.first.x == line.second.x
}
