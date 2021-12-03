package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	oxygen := rating(diagnostics, length, func(zeroes int, ones int, i int, o string) bool {
		return (zeroes > ones && o[i] == '0') || (zeroes <= ones && o[i] == '1')
	})

	co2 := rating(diagnostics, length, func(zeroes int, ones int, i int, o string) bool {
		return (zeroes <= ones && o[i] == '0') || (zeroes > ones && o[i] == '1')
	})

	fmt.Println(oxygen * co2)
}

func rating(diagnostics []string, length int, p func(int, int, int, string) bool) int64 {
	filtered := make([]string, len(diagnostics))
	copy(filtered, diagnostics)
	for i := 0; i < length; i++ {
		if len(filtered) == 1 {
			break
		}
		zeroes, ones := 0, 0
		for _, o := range filtered {
			if o[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		var tmp []string
		for _, o := range filtered {
			if p(zeroes, ones, i, o) {
				tmp = append(tmp, o)
			}
		}
		filtered = tmp
	}

	value, err := strconv.ParseInt(filtered[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
