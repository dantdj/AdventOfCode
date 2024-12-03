package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	input := readInput("real-input.txt")
	mulRe := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	numRe := regexp.MustCompile("[0-9]+")

	total := 0
	mulGroups := mulRe.FindAllString(input, -1)
	for _, group := range mulGroups {
		numbers := numRe.FindAllString(group, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		total += (num1 * num2)
	}

	return total
}

func partTwo() int {
	input := readInput("real-input.txt")
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	numRe := regexp.MustCompile("[0-9]+")

	total := 0
	mulGroups := mulRe.FindAllString(input, -1)
	shouldCount := true
	for _, group := range mulGroups {
		if group == "do()" {
			shouldCount = true
		} else if group == "don't()" {
			shouldCount = false
		} else {
			if shouldCount {
				numbers := numRe.FindAllString(group, -1)
				num1, _ := strconv.Atoi(numbers[0])
				num2, _ := strconv.Atoi(numbers[1])
				total += (num1 * num2)
			}
		}
	}

	return total
}

func readInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)

	return string(bytes)
}
