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
	days, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// assume the values all fit in memory
	var countOfFishWithTimer [9]int
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("expected line")
	}
	rawFish := strings.Split(scanner.Text(), ",")
	for _, raw := range rawFish {
		timer, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		if timer < 0 || timer > 8 {
			log.Fatalf("expected timer between 0 and 8 (inclusive) but was %d", timer)
		}
		countOfFishWithTimer[timer]++
	}
	if scanner.Scan() {
		log.Fatalf("expected end of file but was %s", scanner.Text())
	}

	fmt.Println(countOfFishWithTimer)
	for i := 0; i < days; i++ {
		count := countOfFishWithTimer[0]
		for timer := 1; timer <= 8; timer++ {
			// shift left
			countOfFishWithTimer[timer-1] = countOfFishWithTimer[timer]
		}
		// fish with timer zero reset to timer 6
		countOfFishWithTimer[6] += count
		// fish with timer zero create new fish with timer 8
		countOfFishWithTimer[8] = count
		fmt.Println(countOfFishWithTimer)
	}

	sum := 0
	for j := 0; j <= 8; j++ {
		sum += countOfFishWithTimer[j]
	}
	fmt.Println(sum)
}
