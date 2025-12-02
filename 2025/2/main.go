package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ranges := strings.Split(string(input), ",")
	var minMaxRanges [][2]int

	for _, numRange := range ranges {
		numRange := strings.Split(numRange, "-")
		min, _ := strconv.Atoi(numRange[0])
		max, _ := strconv.Atoi(numRange[1])

		minMaxRanges = append(minMaxRanges, [2]int{min, max})
	}

	start := time.Now()
	fmt.Printf("Part One: %d (%s)\n", partOne(minMaxRanges), time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d (%s)\n", partTwo(minMaxRanges), time.Since(start))
}

func partOne(input [][2]int) int {
	answer := 0
	for _, numRange := range input {
		for i := numRange[0]; i <= numRange[1]; i++ {
			// I'm not a fan of this conversion, but I can't think of a sensible
			// way of doing this without it, as i needs to be an int for the range
			// to work.
			numToCheck := strconv.Itoa(i)
			if len(numToCheck)%2 != 0 {
				// If it's not of even length, it can't be two identical number
				// sequences.
				continue
			}

			firstHalf := numToCheck[:len(numToCheck)/2]
			secondHalf := numToCheck[len(numToCheck)/2:]
			if firstHalf == secondHalf {
				answer += i
			}
		}
	}

	return answer
}

func partTwo(input [][2]int) int {
	answer := 0
	for _, numRange := range input {
		for i := numRange[0]; i <= numRange[1]; i++ {
			numToCheck := strconv.Itoa(i)
			length := len(numToCheck)
			aa := numToCheck + numToCheck

			// A string can be shown to be made up of repeated substrings if and only if
			// it's a rotation of itself. We can check this by seeing if we can find a
			// substring containing the original string in a a string that is the
			// original string repeated twice. If we can't, it's not made up of
			// repeated substrings.
			// We remove the first and last characters to avoid matching the original
			// string with itself.
			// e.g abab -> abababab
			// abababab[1:] -> bababab
			// bababab contains abab
			if strings.Index(aa[1:], numToCheck) != length-1 {
				answer += i
			}
		}
	}

	return answer
}
