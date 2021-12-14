package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("expected 2 args, was %d", len(os.Args)-1)
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

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("expected first line to be polymer template")
	}
	polymerTemplate := scanner.Text()
	match, err := regexp.MatchString("[A-Z]+", polymerTemplate)
	if err != nil {
		log.Fatal(err)
	}
	if !match {
		log.Fatalf("expected polymer template to match [A-Z]+ but was %s", polymerTemplate)
	}
	if !scanner.Scan() {
		log.Fatal("expected second line")
	}
	if len(scanner.Text()) > 0 {
		log.Fatal("expected second line to be blank")
	}

	pairs := make(map[string]int)
	for j := 0; j+1 < len(polymerTemplate); j++ {
		pairs[polymerTemplate[j:j+2]]++
	}

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		match, err := regexp.MatchString("^[A-Z]{2} -> [A-Z]$", line)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf(
				"expected pair insertion rule to match ^[A-Z]{2} -> [A-Z]{2}$ but was %s",
				line)
		}
		fields := strings.Split(line, " -> ")
		rules[fields[0]] = fields[1]
	}

	for i := 0; i < steps; i++ {
		current := make(map[string]int)
		for pair, count := range pairs {
			rule := rules[pair]
			current[pair[0:1]+rule] += count
			current[rule+pair[1:2]] += count
		}
		pairs = current
	}

	counts := make(map[string]int)
	for pair, count := range pairs {
		/*
		 * Only count the first character, the second character will be counted
		 * when it begins a pair.
		 */
		counts[pair[0:1]] += count
	}
	/*
	 * We need to count the last character of the polymer template, it never
	 * started a pair.
	 */
	counts[polymerTemplate[len(polymerTemplate)-1:len(polymerTemplate)]]++

	min, max := math.MaxInt, 0
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max - min)
}
