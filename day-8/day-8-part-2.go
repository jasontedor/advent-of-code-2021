package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"regexp"
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

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " | ")
		if len(fields) != 2 {
			log.Fatalf("expected two pipe-separated fields, was %d: %s", len(fields), scanner.Text())
		}
		signals := parseSignals(strings.Split(fields[0], " "))
		output := strings.Split(fields[1], " ")
		if len(output) != 4 {
			log.Fatalf("expected four-digit output, was %d: %s", len(output), fields[1])
		}

		signalToDigit := identifySignals(signals)
		value := 0
		for _, o := range output {
			value = 10*value + signalToDigit[parseSignal(o)]
		}
		sum += value
	}

	fmt.Println(sum)
}

func parseSignals(rawSignals []string) []uint {
	var signals []uint
	for _, signal := range rawSignals {
		signals = append(signals, parseSignal(signal))
	}
	return signals
}

func parseSignal(signal string) uint {
	match, err := regexp.MatchString("[abcdefg]+", signal)
	if err != nil {
		log.Fatal(err)
	}
	if !match {
		log.Fatalf("expected signal to match [abcdefg]+, was %s", signal)
	}
	// treat a -> 1, b -> 10, c -> 100, ..., g -> 1000000
	a := uint(0)
	for _, c := range signal {
		a |= 1 << (c - 'a')
	}
	return a
}

func identifySignals(signals []uint) map[uint]int {
	/*
	 * Identify the segments with the bits in a 7-bit number.
	 *
	 *      a
	 *    _____
	 *   |     |
	 * f |     | b
	 *   |  g  |
	 *    -----
	 *   |     |
	 * e |     | c
	 *   |  d  |
	 *    -----
	 *
	 * This means that we have the following mapping:
	 *      abcdefg
	 * 0 -> 1111110
	 * 1 -> 0110000
	 * 2 -> 1101101
	 * 3 -> 1111001
	 * 4 -> 0110011
	 * 5 -> 1011011
	 * 6 -> 1011111
	 * 7 -> 1110000
	 * 8 -> 1111111
	 * 9 -> 1110011
	 *
	 * From here, we can derive the representations.
	 */

	var one uint
	var four uint
	signalToDigit := make(map[uint]int)

	/*
	 * 1, 4, 7, and 8 can be identified by counting bits as they uniquely have
	 * one-bits of count 2, 4, 3, and 7 respectively.
	 */
	for _, signal := range signals {
		pc := bits.OnesCount(signal)
		switch pc {
		case 2:
			one = signal
			signalToDigit[signal] = 1
		case 4:
			four = signal
			signalToDigit[signal] = 4
		case 3:
			signalToDigit[signal] = 7
		case 7:
			signalToDigit[signal] = 8
		}
	}

	/*
	 * We know 0110000 is 1, and 0110011 is 4. This means that if we subtract
	 * the representation of 1 from 4, we find the representation of 0000011.
	 * This will important for masking to identify the remaining signals.
	 */
	bc := one
	fg := four - one

	/*
	 * Now we inspect the signals with 5 bits set. They can be uniquely
	 * identified by masking with bc and fg. Among those signals with five bits
	 * set, only 3 masks with bc, only 5 masks with fg, and 2 masks with
	 * neither.
	 */
	for _, signal := range signals {
		pc := bits.OnesCount(signal)
		if pc == 5 {
			if (signal & bc) == bc {
				signalToDigit[signal] = 3
			} else if (signal & fg) == fg {
				signalToDigit[signal] = 5
			} else {
				signalToDigit[signal] = 2
			}
		}
	}

	/*
	 * Now we inspect the signals with 6 bits set. They can be uniquely
	 * identified by masking with bc and fg. Among those signals with six bits
	 * set, only 9 masks with fg and bc, only 6 masks with fg but not bc, and
	 * only 0 masks with bc but not fg.
	 */
	for _, signal := range signals {
		pc := bits.OnesCount(signal)
		if pc == 6 {
			if (signal&fg) == fg && (signal&bc) == bc {
				signalToDigit[signal] = 9
			} else if (signal & fg) == fg {
				signalToDigit[signal] = 6
			} else {
				signalToDigit[signal] = 0
			}
		}
	}

	return signalToDigit
}
