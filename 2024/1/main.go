package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	arrayOne, arrayTwo := readInput("real-data.txt")

	slices.Sort(arrayOne)
	slices.Sort(arrayTwo)

	totalDistance := 0
	// Assuming lists are same length
	for index, _ := range arrayOne {
		totalDistance += absDiff(arrayOne[index], arrayTwo[index])
	}

	return totalDistance
}

func partTwo() int {
	arrayOne, arrayTwo := readInput("real-data.txt")

	similarityScore := 0

	for _, item := range arrayOne {
		count := 0
		for _, itemTwo := range arrayTwo {
			if item == itemTwo {
				count += 1
			}
		}
		similarityScore += item * count
	}

	return similarityScore
}

func readInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)

	arrayOne := []int{}
	arrayTwo := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.FindAllString(line, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		arrayOne = append(arrayOne, num1)
		arrayTwo = append(arrayTwo, num2)
	}

	return arrayOne, arrayTwo
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
