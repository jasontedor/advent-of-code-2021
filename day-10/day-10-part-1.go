package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	scores := make(map[int32]int)
	scores[')'] = 3
	scores[']'] = 57
	scores['}'] = 1197
	scores['>'] = 25137
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		match, err := regexp.MatchString("[\\(\\[\\{<\\)\\]\\}>]+", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected line to match [\\(\\[\\{<\\)\\]\\}>]+ but was %s", line)
		}
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
				fmt.Println(line)
				score += scores[c]
				break
			}
		}
	}

	fmt.Println(score)
}
