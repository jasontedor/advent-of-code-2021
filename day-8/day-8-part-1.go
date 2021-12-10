package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " | ")
		if len(fields) != 2 {
			log.Fatalf("expected two pipe-separated fields, was %d: %s", len(fields), scanner.Text())
		}
		output := strings.Split(fields[1], " ")
		if len(output) != 4 {
			log.Fatalf("expected four-digit output, was %d: %s", len(output), fields[1])
		}
		for _, value := range output {
			if len(value) == 2 || len(value) == 4 || len(value) == 3 || len(value) == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}
