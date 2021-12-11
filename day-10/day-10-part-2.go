package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected 1 arg, was %d", len(os.Args)-1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var scores []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match, err := regexp.MatchString("[(\\[{<)\\]}>]+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match [(\\[{<)\\]}>]+ but was %s", line)
		}
		corrupt := false
		var stack []int32
		for _, c := range line {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack = append(stack, c)
				continue
			}
			if len(stack) == 0 {
				log.Fatalf("expected non-empty stack")
			}
			if c == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if c == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if c == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if c == '>' && stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1]
			} else {
				// line is corrupt, break
				corrupt = true
				break
			}
		}
		if !corrupt {
			score := 0
			for len(stack) > 0 {
				c := stack[len(stack)-1]
				switch c {
				case '(':
					score = 5*score + 1
				case '[':
					score = 5*score + 2
				case '{':
					score = 5*score + 3
				case '<':
					score = 5*score + 4
				}
				stack = stack[:len(stack)-1]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
