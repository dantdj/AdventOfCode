package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	input := readInput("real-data.txt")
	safeReports := 0
	for levelIdx, _ := range input {
		safeReport := true
		increasing := false

		if input[levelIdx][0] < input[levelIdx][1] {
			increasing = true
		}
		for reportIdx := range len(input[levelIdx]) - 1 {
			if increasing && input[levelIdx][reportIdx] > input[levelIdx][reportIdx+1] {
				safeReport = false
				break
			} else if !increasing && input[levelIdx][reportIdx] < input[levelIdx][reportIdx+1] {
				safeReport = false
				break
			}
			difference := absDiff(input[levelIdx][reportIdx], input[levelIdx][reportIdx+1])
			if difference > 3 || difference < 1 {
				// unsafe, don't continue with this level
				safeReport = false
				break
			}
		}
		if safeReport {
			safeReports += 1
		}
	}
	return safeReports
}

func partTwo() int {
	input := readInput("real-data.txt")
	safeReports := 0
	for levelIdx, _ := range input {
		if checkSafety(input[levelIdx]) {
			safeReports += 1
		}
	}
	return safeReports
}

func checkSafety(input []int) bool {
	safeReports := 0
	for idx1 := range len(input) {
		removedLevelInput := make([]int, len(input))
		copy(removedLevelInput, input)
		removedLevelInput = remove(removedLevelInput, idx1)

		increasing := false
		if removedLevelInput[0] < removedLevelInput[1] {
			increasing = true
		}
		unsafeLevels := 0

		for idx := range len(removedLevelInput) - 1 {
			current := removedLevelInput[idx]
			next := removedLevelInput[idx+1]

			if increasing && current > next {
				unsafeLevels += 1
				continue
			} else if !increasing && current < next {
				unsafeLevels += 1
				continue
			}

			difference := absDiff(current, next)
			if difference > 3 || difference < 1 {
				unsafeLevels += 1
				continue
			}
		}
		if unsafeLevels == 0 {
			safeReports += 1
		}
	}

	return safeReports > 0
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	listOfSlices := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := []int{}
		numStrs := strings.Fields(line)
		for _, value := range numStrs {
			num, _ := strconv.Atoi(value)
			nums = append(nums, num)
		}
		listOfSlices = append(listOfSlices, nums)
	}

	return listOfSlices
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
