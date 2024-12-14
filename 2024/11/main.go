package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	defer timeTrack(time.Now(), "partOne")
	rocks := readInput("real-input.txt")

	blinkCounter := 0

	for {
		//fmt.Printf("blinked %d times\n", blinkCounter)
		if blinkCounter == 25 {
			break
		}

		newRocks := []int{}
		for i := range rocks {
			if rocks[i] == 0 {
				newRocks = append(newRocks, 1)
			} else if numberOfDigits(rocks[i])%2 == 0 {
				oldRock, newRock := splitRock(rocks[i])
				newRocks = append(newRocks, oldRock)
				newRocks = append(newRocks, newRock)
			} else {
				newRocks = append(newRocks, rocks[i]*2024)
			}
		}

		rocks = newRocks

		blinkCounter++
	}

	return len(rocks)
}

func partTwo() int {
	defer timeTrack(time.Now(), "partTwo")
	rocks := readInput("real-input.txt")

	totalRocks := 0

	// Map of rock value to number of times it appears.
	// This works as you aren't storing duplicates of the same value,
	// just the number of times it appears. That's fine here for our calculations
	// as they're not dependent on the rock order, just that they're present.
	rockMap := map[int]int{}

	// Convert input to the above map
	for _, rock := range rocks {
		rockMap[rock] = rockMap[rock] + 1
	}

	blinkCounter := 0

	for {
		//fmt.Printf("blinked %d times\n", blinkCounter)
		if blinkCounter == 75 {
			break
		}

		newRocks := map[int]int{}
		for key, value := range rockMap {
			if key == 0 {
				newRocks[1] = newRocks[1] + value
			} else if numberOfDigits(key)%2 == 0 {
				oldRock, newRock := splitRock(key)
				newRocks[oldRock] = newRocks[oldRock] + value
				newRocks[newRock] = newRocks[newRock] + value
			} else {
				newKey := key * 2024
				newRocks[newKey] = newRocks[newKey] + value
			}
		}

		rockMap = newRocks

		blinkCounter++
	}

	// Sum up the number of rocks in the map
	for _, value := range rockMap {
		totalRocks += value
	}

	return totalRocks
}

// Assuming an even-length value, splits a given number in two
func splitRock(value int) (int, int) {
	numStr := strconv.Itoa(value)
	middle := len(numStr) / 2

	first := numStr[:middle]
	second := numStr[middle:]

	firstInt, _ := strconv.Atoi(first)
	secondInt, _ := strconv.Atoi(second)

	return firstInt, secondInt
}

// Return the number of digits the given number has
func numberOfDigits(n int) int {
	if n == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func readInput(filename string) []int {
	content, _ := os.ReadFile(filename)

	strContent := string(content)
	fields := strings.Fields(strContent)

	rocks := []int{}

	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		rocks = append(rocks, num)
	}

	return rocks
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
