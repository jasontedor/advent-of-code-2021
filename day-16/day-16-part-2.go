package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("expected input line")
	}
	transmission := scanner.Text()
	if scanner.Scan() {
		log.Fatal("expected end of file")
	}

	packet, err := DecodePacket(transmission)
	if err != nil {
		log.Fatal(err)
	}
	value, err := packet.Value()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
