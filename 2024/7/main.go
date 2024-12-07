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
	equations := readInput("real-input.txt")

	total := 0

	for _, equation := range equations {
		ans := calculate(equation.CorrectAnswer, equation.Operands[0], equation.Operands[1], equation.Operands[2:]...)

		validSum := ans == equation.CorrectAnswer
		if validSum {
			total += equation.CorrectAnswer
		}
	}
	return total
}

func calculate(ans, first, second int, rest ...int) int {
	operators := []string{"+", "*"}
	sumTotal := 0
	productTotal := 0
	for _, operator := range operators {
		if operator == "+" {
			sumTotal = first + second

			if len(rest) > 0 {
				sumTotal = calculate(ans, sumTotal, rest[0], rest[1:]...)
			}
		} else if operator == "*" {
			productTotal = first * second

			if len(rest) > 0 {
				productTotal = calculate(ans, productTotal, rest[0], rest[1:]...)
			}
		} else {
			panic("shouldn't be getting here")
		}
	}

	if sumTotal == ans || productTotal == ans {
		return ans
	}

	return 0
}

func partTwo() int {
	equations := readInput("real-input.txt")

	total := 0

	for _, equation := range equations {
		ans := calculatePartTwo(equation.CorrectAnswer, equation.Operands[0], equation.Operands[1], equation.Operands[2:]...)

		validSum := ans == equation.CorrectAnswer
		if validSum {
			total += equation.CorrectAnswer
		}
	}
	return total
}

func calculatePartTwo(ans, first, second int, rest ...int) int {
	operators := []string{"+", "*", "||"}
	sumTotal := 0
	productTotal := 0
	concatTotal := 0
	for _, operator := range operators {
		if operator == "+" {
			sumTotal = first + second

			if len(rest) > 0 {
				sumTotal = calculatePartTwo(ans, sumTotal, rest[0], rest[1:]...)
			}
		} else if operator == "*" {
			productTotal = first * second

			if len(rest) > 0 {
				productTotal = calculatePartTwo(ans, productTotal, rest[0], rest[1:]...)
			}
		} else if operator == "||" {
			concatTotal = concatenateInts(first, second)
			if len(rest) > 0 {
				concatTotal = calculatePartTwo(ans, concatTotal, rest[0], rest[1:]...)
			}
		} else {
			panic("shouldn't be getting here")
		}
	}

	if sumTotal == ans || productTotal == ans || concatTotal == ans {
		return ans
	}

	return 0
}

func readInput(filename string) []Equation {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	listOfEquations := []Equation{}

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		e := Equation{}
		ans, _ := strconv.Atoi(lineParts[0])
		e.CorrectAnswer = ans

		operands := strings.Fields(lineParts[1])
		nums := []int{}
		for _, operand := range operands {
			num, _ := strconv.Atoi(operand)
			nums = append(nums, num)
		}

		e.Operands = nums
		listOfEquations = append(listOfEquations, e)
	}

	return listOfEquations
}

// given e.g 12 and 34, combines into 1234
func concatenateInts(first, second int) int {
	str1 := strconv.Itoa(first)
	str2 := strconv.Itoa(second)

	combined := str1 + str2

	combinedInt, _ := strconv.Atoi(combined)

	return combinedInt
}

type Equation struct {
	CorrectAnswer int
	Operands      []int
}
