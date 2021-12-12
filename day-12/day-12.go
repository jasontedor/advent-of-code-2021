package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseAdjacencyLists(path string) map[string][]string {
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

	adjacencyLists := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "-")
		if len(fields) != 2 {
			log.Fatalf("expected 2 fields but was %d: %s", len(fields), fields)
		}

		if adjacencyLists[fields[0]] == nil {
			adjacencyLists[fields[0]] = make([]string, 0)
		}
		adjacencyLists[fields[0]] = append(adjacencyLists[fields[0]], fields[1])

		if adjacencyLists[fields[1]] == nil {
			adjacencyLists[fields[1]] = make([]string, 0)
		}
		adjacencyLists[fields[1]] = append(adjacencyLists[fields[1]], fields[0])
	}

	return adjacencyLists
}

func FindPaths(
	adjacencyLists map[string][]string,
	c string,
	t string,
	path *[]string,
	paths *[][]string,
	accept func(path []string, c string) bool) {
	*path = append(*path, c)

	if c == t {
		tmp := make([]string, len(*path)-1)
		copy(tmp, *path)
		*paths = append(*paths, append(tmp, t))
		*path = tmp
		return
	} else {
		for _, n := range adjacencyLists[c] {
			if !accept(*path, n) {
				continue
			}
			FindPaths(adjacencyLists, n, t, path, paths, accept)
		}
	}

	tmp := make([]string, len(*path)-1)
	copy(tmp, *path)
	*path = tmp
}
