package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}

func partOne() {
	input, _ := readInput("real-input")
	
	nextValues := 0
	for _, history := range input {
		nextValues += extrapolate(history, false)
	}
	
	fmt.Printf("Part One: %d", nextValues)
}

func partTwo() {
	input, _ := readInput("real-input")

	nextValues := 0
	for _, history := range input {
		nextValues += extrapolate(history, true)
	}

	fmt.Printf("Part Two: %d", nextValues)
}

func extrapolate(history []int, backwards bool) int {
	// Check if our input is all zeroes - if it is, we've reached the base case, so return 0
	allZeroes := true
	for _, entry := range history {
		if entry != 0 {
			allZeroes = false
		}
	}
	
	if allZeroes {
		return 0
	}
	
	// Otherwise, create an array that contains the differences between each number in the input,
	// and then pass that to this function again. Return the number at the end of the array added
	// to the number returned by this function
	differences := generateDifferences(history)
	value := extrapolate(differences, backwards)
	
	// If we're extrapolating backwards, subtract the value from the first history value
	if backwards {
		return history[0] - value
	}
	
	// If extrapolating forwards, add the value to the last history value
	return history[len(history)-1] + value
}

func generateDifferences(input []int) []int {
	differences := []int{}
	
	// len(input)-1 so we don't go out of bounds when comparing i and i+1 
	for i := 0; i < len(input) - 1; i++ {
		differences = append(differences, diff(input[i], input[i+1]))
	}
	
	return differences
}

func diff(a, b int) int {
	return b - a
}

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return [][]int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	readings := [][]int{}
	for scanner.Scan() {
		numStrs := strings.Split(scanner.Text(), " ")
		readings = append(readings, convertStrSliceToIntSlice(numStrs))
	}
	
	return readings, nil
}

func convertStrSliceToIntSlice(input []string) []int {
	nums := make([]int, len(input))
	for i, x := range input {
		nums[i], _ = strconv.Atoi(x)
	}
	return nums
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}