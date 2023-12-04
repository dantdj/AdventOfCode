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
	totalScore := 0
	for _, card := range input {
		matches := 0
		for _, myNumber := range card.MyNumbers {
			for _, winningNumber := range card.WinningNumbers {
				if myNumber == winningNumber {
					matches++
				}
			}
		}

		cardScore := 0
		for i := 1; i <= matches; i++ {
			if i == 1 {
				// First match is worth one
				cardScore += 1
				continue
			}
			cardScore *= 2
		}
		totalScore += cardScore
	}

	fmt.Printf("Part One: %d", totalScore)
}

func partTwo() {
	input, _ := readInput("real-input")

	// Populate initial card stacks
	cardStacks := []CardStack{}
	for _, card := range input {
		cardStacks = append(cardStacks, CardStack{CardNumber: card.CardNumber, Cards: []ScratchCard{card}})
	}

	numberOfCards := 0

	for _, cardStack := range cardStacks {
		for _, card := range cardStack.Cards {
			numberOfCards += 1
			matches := 0
			for _, myNumber := range card.MyNumbers {
				for _, winningNumber := range card.WinningNumbers {
					if myNumber == winningNumber {
						matches++
					}
				}
			}
			
			for i := card.CardNumber; i < matches+card.CardNumber; i++ {
				cardStacks[i].Cards = append(cardStacks[i].Cards, input[i])
			} 
		}
	}
	
	fmt.Printf("Part Two: %d", numberOfCards)
}

func readInput(filename string) ([]ScratchCard, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	scratchcards := []ScratchCard{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Divide into card number and card state
		cardSplit := strings.Split(scanner.Text(), ": ")
		cardNumberStrs := strings.Split(cardSplit[0], " ")
		cardNumber, _ := strconv.Atoi(cardNumberStrs[len(cardNumberStrs) - 1])

		// Split into my numbers and winning numbers
		numberSplit := strings.Split(cardSplit[1], " | ")

		// Split into individual values
		myNumbers := strings.Split(numberSplit[0], " ")
		winningNumbers := strings.Split(numberSplit[1], " ")

		scratchcard := ScratchCard{CardNumber: cardNumber}

		for _, numberStr := range myNumbers {
			value, _ := strconv.Atoi(numberStr)
			if value == 0 {
				// We probably tried to parse an empty string, skip
				// This can be left over from the string split
				continue
			}
			scratchcard.MyNumbers = append(scratchcard.MyNumbers, value)
		}

		for _, numberStr := range winningNumbers {
			value, _ := strconv.Atoi(numberStr)
			if value == 0 {
				// We probably tried to parse an empty string, skip
				// This can be left over from the string split
				continue
			}
			scratchcard.WinningNumbers = append(scratchcard.WinningNumbers, value)
		}

		scratchcards = append(scratchcards, scratchcard)
	}

	return scratchcards, scanner.Err()
}

type ScratchCard struct {
	CardNumber     int
	MyNumbers      []int
	WinningNumbers []int
}

type CardStack struct {
	CardNumber int
	Cards      []ScratchCard
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
