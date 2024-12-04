package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	input := readInput("real-input.txt")
	count := 0
	for yIdx, _ := range input {
		for xIdx, _ := range input[yIdx] {
			if input[yIdx][xIdx] == "X" {

				// get 3 letters to the left
				if xIdx-3 >= 0 {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx][xIdx-i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going up-left
				if xIdx-3 >= 0 && yIdx-3 >= 0 {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx-i][xIdx-i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going down-left
				if xIdx-3 >= 0 && yIdx+3 < len(input) {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx+i][xIdx-i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going up
				if yIdx-3 >= 0 {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx-i][xIdx]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going down
				if yIdx+3 < len(input) {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx+i][xIdx]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going right
				if xIdx+3 < len(input[yIdx]) {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx][xIdx+i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going up-right
				if xIdx+3 < len(input[yIdx]) && yIdx-3 >= 0 {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx-i][xIdx+i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}

				// get 3 letters going down-right
				if xIdx+3 < len(input[yIdx]) && yIdx+3 < len(input) {
					nextThree := ""
					for i := 1; i <= 3; i++ {
						nextThree += input[yIdx+i][xIdx+i]
					}
					if nextThree == "MAS" {
						count += 1
					}
				}
			}
		}
	}

	return count
}

func partTwo() int {
	input := readInput("real-input.txt")
	count := 0
	for yIdx, _ := range input {
		for xIdx, _ := range input[yIdx] {
			if input[yIdx][xIdx] == "A" {
				if yIdx-1 >= 0 && xIdx-1 >= 0 && yIdx+1 < len(input) && xIdx+1 < len(input[yIdx]) {
					// get the two arms of the x

					arm1 := ""
					arm1 += input[yIdx-1][xIdx-1]
					arm1 += input[yIdx][xIdx]
					arm1 += input[yIdx+1][xIdx+1]

					arm2 := ""
					arm2 += input[yIdx-1][xIdx+1]
					arm2 += input[yIdx][xIdx]
					arm2 += input[yIdx+1][xIdx-1]

					if (arm1 == "SAM" || arm1 == "MAS") && (arm2 == "SAM" || arm2 == "MAS") {
						count += 1
					}
				}
			}
		}
	}

	return count
}

func readInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	listOfSlices := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}

		listOfSlices = append(listOfSlices, chars)
	}

	return listOfSlices
}
