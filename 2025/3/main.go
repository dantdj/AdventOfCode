package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var banks [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var bank []int
		for _, r := range line {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Printf("Error converting character to int: %v\n", err)
				panic(err)
			}
			bank = append(bank, num)
		}
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	start := time.Now()
	fmt.Printf("Part One: %d (%s)\n", partOne(banks), time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d (%s)\n", partTwo(banks), time.Since(start))

}

func partOne(banks [][]int) int {
	answer := 0

	// If we keep track of the largest first digit we've seen,
	// we can avoid a nested loop because we know we'll never go back
	// from that for the first digit, so we don't need to know about another
	// potential first digit.
	for _, bank := range banks {
		maxDigit := bank[0]
		maxValue := 0
		for i := 1; i < len(bank); i++ {
			// Add the two numbers together (e.g 9 and 7 become 97)
			value := maxDigit*10 + bank[i]
			if value > maxValue {
				maxValue = value
			}

			if bank[i] > maxDigit {
				maxDigit = bank[i]
			}
		}
		answer += maxValue
	}
	return answer
}

func partTwo(banks [][]int) int {
	return 0
}
