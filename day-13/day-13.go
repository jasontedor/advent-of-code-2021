package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Location struct {
	x int
	y int
}

type Orientation string

const (
	Vertical   Orientation = "y"
	Horizontal             = "x"
)

type Fold struct {
	orientation Orientation
	position    int
}

func ParseLocationsAndFolds(path string) (map[Location]bool, []Fold) {
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

	locations := make(map[Location]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		match, err := regexp.MatchString("\\d+,\\d+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match \\d+,\\d+ but was %s", line)
		}
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			log.Fatalf("expected 2 fields but was %d: %s", len(fields), line)
		}
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		locations[Location{x, y}] = true
	}

	var folds []Fold
	for scanner.Scan() {
		line := scanner.Text()
		match, err := regexp.MatchString("fold along [xy]=\\d+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match fold along [xy]=\\d+ but was %s", line)
		}
		fields := strings.Split(line[len("fold along "):len(line)], "=")
		value, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		folds = append(folds, Fold{Orientation(fields[0]), value})
	}

	return locations, folds
}

func FoldLocations(locations map[Location]bool, fold Fold) {
	for location, _ := range locations {
		switch fold.orientation {
		case "x":
			if location.x > fold.position {
				difference := location.x - fold.position
				locations[Location{fold.position - difference, location.y}] = true
				delete(locations, location)
			}
		case "y":
			if location.y > fold.position {
				difference := location.y - fold.position
				locations[Location{location.x, fold.position - difference}] = true
				delete(locations, location)
			}
		}
	}
}
