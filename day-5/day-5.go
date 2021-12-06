package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	first  Point
	second Point
}

func ParsePoint(rawPoint string) Point {
	rawValues := strings.Split(rawPoint, ",")
	if len(rawValues) != 2 {
		log.Fatalf("expected 2 coordinates but was %d coordinates: %s", len(rawValues), rawValues)
	}
	x, err := strconv.Atoi(rawValues[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(rawValues[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}

func ParseLines(path string) []Line {
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

	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawPoints := strings.Split(scanner.Text(), " -> ")
		if len(rawPoints) != 2 {
			log.Fatalf("expected 2 points but was %d points: %s", len(rawPoints), rawPoints)
		}
		line := Line{ParsePoint(rawPoints[0]), ParsePoint(rawPoints[1])}
		lines = append(lines, line)
	}
	return lines
}

func MakeGrid(lines []Line) [][]int {
	maxX, maxY := 0, 0
	for _, line := range lines {
		if line.first.x > maxX {
			maxX = line.first.x
		}
		if line.second.x > maxX {
			maxX = line.second.x
		}
		if line.first.y > maxY {
			maxY = line.first.y
		}
		if line.second.y > maxY {
			maxY = line.second.y
		}
	}

	var grid = make([][]int, 1+maxX)
	for i := 0; i < 1+maxX; i++ {
		grid[i] = make([]int, 1+maxY)
	}
	return grid
}

func NextPointOnLine(current Point, target Point) Point {
	nextX, nextY := current.x, current.y
	if current.x < target.x {
		nextX++
	} else if current.x > target.x {
		nextX--
	}
	if current.y < target.y {
		nextY++
	} else if current.y > target.y {
		nextY--
	}
	return Point{nextX, nextY}
}

func CountGrid(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}
