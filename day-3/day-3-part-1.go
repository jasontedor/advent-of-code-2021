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

	var diagnostics []string
	length := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diagnostic := scanner.Text()
		match, err := regexp.MatchString("^[01]+$", diagnostic)
		if err != nil {
			log.Fatal(err)
		}
		if !match {
			log.Fatalf("expected diagnostic to match ^[01]+$ but was %s", diagnostic)
		}
		if length == 0 {
			length = len(diagnostic)
		}
		if length != len(diagnostic) {
			log.Fatalf(
				"expected diagnostic %s to be of same length as previous length %d, but was %d",
				diagnostic,
				length,
				len(diagnostic),
			)
		}
		diagnostics = append(diagnostics, scanner.Text())
	}

	gamma, epsilon := 0, 0
	for i := 0; i < length; i++ {
		zeroes, ones := 0, 0
		for _, diagnostic := range diagnostics {
			if diagnostic[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes == ones {
			log.Fatalf("unable to determine most common bit for position %d", i)
		}
		if zeroes > ones {
			gamma = 2 * gamma
			epsilon = 2*epsilon + 1
		} else {
			gamma = 2*gamma + 1
			epsilon = 2 * epsilon
		}
	}

	fmt.Println(gamma * epsilon)
}
