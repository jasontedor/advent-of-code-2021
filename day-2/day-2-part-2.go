package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	h, d, a := 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			log.Fatal("expected 2 fields, was %d, %s", len(fields), fields)
		}
		mag, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		if fields[0] == "forward" {
			h += mag
			d += a * mag
		} else if fields[0] == "up" {
			a -= mag
		} else if fields[0] == "down" {
			a += mag
		} else {
			log.Fatal("unexpected command, was %s", fields[0])
		}
	}

	fmt.Println(h * d)
}
