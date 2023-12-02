package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	measureTime(partOne) // Correct answer: 54338
	measureTime(partTwo) // Correct answer: 53389
}

func partOne() {
	input, err := readInput("real-input")
	if err != nil {
		os.Exit(1)
	}

	var values []int

	for _, line := range input {
		numbers := extractNumbersFromStringPartOne(line)
		num1 := numbers[0]
		num2 := numbers[len(numbers)-1]
		value, _ := strconv.Atoi(fmt.Sprintf("%d%d", num1, num2))
		values = append(values, value)
	}

	result := 0

	for _, value := range values {
		result += value
	}

	fmt.Printf("Part One: %d", result)
}

func partTwo() {
	input, err := readInput("real-input")
	if err != nil {
		os.Exit(1)
	}

	var values []int

	for _, line := range input {
		numbers := extractNumbersFromStringPartTwo(line)
		num1 := numbers[0]
		num2 := numbers[len(numbers)-1]
		value, _ := strconv.Atoi(fmt.Sprintf("%d%d", num1.Value, num2.Value))
		values = append(values, value)
	}

	result := 0

	for _, value := range values {
		result += value
	}

	fmt.Printf("Part Two: %d", result)
}

// Extracts only digits from the given string
func extractNumbersFromStringPartOne(input string) []int {
	var numbers []int
	runes := []rune(input)

	for _, value := range runes {
		num := int(value - '0')
		if num > 0 && num < 10 {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

// Extracts digits and numbers in the form of words from the given string, and returns an array
// sorted by the index that they started at so we can use the ordering
func extractNumbersFromStringPartTwo(input string) DigitSlice {
	numberWords := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	var numbers DigitSlice

	// First, look through the string to find any numbers that are hanging around in word form
	// No particular reason, could do either this _or_ the pure digit extraction first
	// Need to make sure that it handles finding _all_ instances of a substring, as some strings have the same number word
	// more than once
	for key := range numberWords {
		indices := findAllMatchIndices(input, key)
		for _, index := range indices {
			numbers = append(numbers, Digit{numberWords[key], index})
		}
	}

	// Then, sweep through the string to find any digits, and store them and the indexes
	runes := []rune(input)

	for index, value := range runes {
		num := int(value - '0')
		if num > 0 && num < 10 {
			numbers = append(numbers, Digit{num, index})
		}
	}

	// Then sort by the index
	sort.Sort(numbers)

	return numbers
}

func findAllMatchIndices(input string, substr string) []int {
	indices := []int{}

	startIndex := 0
	for {
		index := strings.Index(input[startIndex:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, startIndex+index)
		startIndex += index + len(substr)
	}

	return indices
}

type Digit struct {
	Value int
	Index int
}

// Implement slice sorting methods by index
type DigitSlice []Digit

func (ds DigitSlice) Len() int           { return len(ds) }
func (ds DigitSlice) Swap(i, j int)      { ds[i], ds[j] = ds[j], ds[i] }
func (ds DigitSlice) Less(i, j int) bool { return ds[i].Index < ds[j].Index }

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
