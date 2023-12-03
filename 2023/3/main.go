package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

func main() {
	measureTime(partOne) // Correct answer: 526404
	measureTime(partTwo)
}

func partOne() {
	input, _ := readInput("real-input")

	partNumberSum := 0

	for y, row := range input {
		// Extract numbers and their indexes from the row
		numbers := extractNumbersFromRow(row, y)

		for _, number := range numbers {
			foundSymbol := false
			for _, point := range number.Points {
				if checkForSymbolPartOne(input, point) {
					foundSymbol = true
					//time.Sleep(1*time.Second)
					break
				}
			}

			if foundSymbol {
				partNumberSum += number.Value
			}
		}
	}

	fmt.Printf("Part One: %d", partNumberSum)
}

func partTwo() {
	// Do the same as part one, and find all the numbers and begin iterating through them. This time, store the location where the symbol was found
	// Store the location of any '*' characters, and then see if _only_ two numbers have a symbol location that matches that location.
	// This means I'll also need to persist the nubers for longer than just a row - I can maintain just a flat list of them outside though

	input, _ := readInput("real-input")
	numbersList := []Number{}
	asterisks := []Symbol{}

	for y, row := range input {
		// Extract numbers and their indexes from the row
		numbers := extractNumbersFromRow(row, y)

		for _, number := range numbers {
			for _, point := range number.Points {
				symbolFound, symbol := checkForSymbolPartTwo(input, point)
				if symbolFound {
					if symbol.IsAsterisk {
						number.AdjacentAsteriskCoord = symbol.Coord

						numbersList = append(numbersList, number)
						asterisks = append(asterisks, symbol)
					}

					//time.Sleep(1*time.Second)
					break
				}
			}
		}
	}

	// For each number in the numbers list, check the other numbers and see if they have the same adjacent asterisk coord
	// Add those numbers to another list - if the length of that list is 2, then find the gear ratio of that list
	gearRatioSum := 0
	for i := 0; i < len(numbersList); i++ {
		firstCoord := numbersList[i].AdjacentAsteriskCoord
		if (numbersList[i].SkipCalculation) {
			continue
		}
		matches := []Number{}
		for j := 0; j < len(numbersList); j++ {
			if i == j {
				// Checking same element - continue instead
				continue
			}
			if numbersList[j].SkipCalculation {
				continue
			}
			secondCoord := numbersList[j].AdjacentAsteriskCoord
			if firstCoord.X == secondCoord.X && firstCoord.Y == secondCoord.Y {
				matches = append(matches, numbersList[j])
				numbersList[j].SkipCalculation = true
			}
		}
		
		// We only want one other match
		if len(matches) == 1 {
			gearRatioSum += numbersList[i].Value * matches[0].Value
		}
	}

	fmt.Printf("Part Two: %d", gearRatioSum)
}

// Return true if a symbol is found adjacent to a digit, as that means the overall number has a symbol adjacent
func checkForSymbolPartOne(input [][]rune, coord Point) bool {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			// Skip if trying the access would result in index out of range
			if coord.Y+y < 0 || coord.Y+y >= len(input) {
				continue
			}
			if coord.X+x < 0 || coord.X+x >= len(input[coord.Y+y]) {
				continue
			}

			if !unicode.IsDigit(input[coord.Y+y][coord.X+x]) && input[coord.Y+y][coord.X+x] != '.' {
				return true
			}
		}
	}

	return false
}

// Return true if a symbol is found adjacent to a digit, as that means the overall number has a symbol adjacent
// Returns the gear location if one was found. I know it's done kinda horribly, but again prioritizing dev speed
func checkForSymbolPartTwo(input [][]rune, coord Point) (bool, Symbol) {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			// Skip if trying the access would result in index out of range
			if coord.Y+y < 0 || coord.Y+y >= len(input) {
				continue
			}
			if coord.X+x < 0 || coord.X+x >= len(input[coord.Y+y]) {
				continue
			}

			if !unicode.IsDigit(input[coord.Y+y][coord.X+x]) && input[coord.Y+y][coord.X+x] != '.' {
				if input[coord.Y+y][coord.X+x] == '*' {
					return true, Symbol{IsAsterisk: true, Coord: Point{X: coord.X + x, Y: coord.Y + y}}
				}
				return true, Symbol{}
			}
		}
	}

	return false, Symbol{}
}

func extractNumbersFromRow(inputLine []rune, rowIndex int) []Number {
	numStr := ""
	numbers := []Number{}
	number := Number{}

	// Iterate through building the string until we reach the next non-digit on the other side
	for i := 0; i < len(inputLine); i++ {
		if !unicode.IsDigit(inputLine[i]) {
			// No digits have been added, so nothing to do here
			if len(number.Points) == 0 {
				continue
			}

			// Add current number to the list and reset so we can keep going until the end
			value, _ := strconv.Atoi(numStr)
			number.Value = value
			numbers = append(numbers, number)

			number = Number{}
			numStr = ""
			continue
		}
		numStr += string(inputLine[i])

		number.Points = append(number.Points, Point{X: i, Y: rowIndex})
	}

	// Append final number from the loop if it's not empty, as it wouldn't have been submitted if it went
	// to the edge of the row
	if len(number.Points) > 0 {
		value, _ := strconv.Atoi(numStr)
		number.Value = value
		numbers = append(numbers, number)
	}

	return numbers
}

type Point struct {
	X int
	Y int
}

type Number struct {
	Value int
	// The coord of each individual digit
	Points []Point
	// The coord of an adjacent gear - used for part two
	AdjacentAsteriskCoord Point
	SkipCalculation bool
}

type Symbol struct {
	IsAsterisk           bool
	Coord                Point
	AdjacentNumbersCount int
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		characters := []rune{}
		for _, character := range scanner.Text() {
			characters = append(characters, character)
		}
		grid = append(grid, characters)
	}

	return grid, scanner.Err()
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
