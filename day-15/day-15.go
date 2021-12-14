package main

import (
	"bufio"
	"container/heap"
	"log"
	"math"
	"os"
	"regexp"
)

type Position struct {
	row    int
	column int
}

type Item struct {
	position  Position
	riskLevel int
	distance  int
	index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func ParseRiskLevels(path string) (map[Position]int, int, int) {
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

	columns := 0
	riskLevels := make(map[Position]int)
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		match, err := regexp.MatchString("\\d+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match \\d+ but was %s", line)
		}
		if columns == 0 {
			columns = len(line)
		}
		if len(line) != columns {
			log.Fatalf(
				"inconsistent line columns, expected %d but was %d: %s",
				columns,
				len(line),
				line)
		}
		for column, c := range line {
			riskLevels[Position{row, column}] = int(c - '0')
		}
		row++
	}

	return riskLevels, row, columns
}

func unvisitedNeighbors(
	position Position,
	visited map[Position]bool,
	rows int,
	columns int) []Position {
	var positions []Position
	if position.row-1 >= 0 && !visited[Position{position.row - 1, position.column}] {
		positions = append(positions, Position{position.row - 1, position.column})
	}
	if position.row+1 < rows && !visited[Position{position.row + 1, position.column}] {
		positions = append(positions, Position{position.row + 1, position.column})
	}
	if position.column-1 >= 0 && !visited[Position{position.row, position.column - 1}] {
		positions = append(positions, Position{position.row, position.column - 1})
	}
	if position.column+1 < columns && !visited[Position{position.row, position.column + 1}] {
		positions = append(positions, Position{position.row, position.column + 1})
	}
	return positions
}

func ShortestPath(riskLevels map[Position]int, rows int, columns int) int {
	items := make(map[Position]*Item)
	pq := make(PriorityQueue, rows*columns)
	for position, riskLevel := range riskLevels {
		index := rows*position.row + position.column
		distance := math.MaxInt
		if position.row == 0 && position.column == 0 {
			distance = 0
		}
		item := &Item{position, riskLevel, distance, index}
		items[position] = item
		pq[index] = item
	}

	heap.Init(&pq)

	visited := make(map[Position]bool)
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		visited[item.position] = true
		if item.position.row == rows-1 && item.position.column == columns-1 {
			break
		}
		for _, i := range unvisitedNeighbors(item.position, visited, rows, columns) {
			d := item.distance + riskLevels[i]
			if d < items[i].distance {
				items[i].distance = d
				heap.Fix(&pq, items[i].index)
			}
		}
	}

	return items[Position{rows - 1, columns - 1}].distance
}
