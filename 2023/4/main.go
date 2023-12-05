package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	// Populate initial card stacks, and calculate the win states so we know how many new cards
	// to create if they win. This saves time on calculating the answer every loop, as it'll always remain
	// the same
	cardStacks := []CardStack{}
	for _, card := range input {
		matches := 0
		for _, myNumber := range card.MyNumbers {
			for _, winningNumber := range card.WinningNumbers {
				if myNumber == winningNumber {
					matches++
				}
			}
		}

		card.NumberOfMatches = matches
		cardStacks = append(cardStacks, CardStack{CardNumber: card.CardNumber, Cards: []ScratchCard{card}})
	}

	numberOfCards := 0

	for _, cardStack := range cardStacks {
		for _, card := range cardStack.Cards {
			numberOfCards += 1
			
			for i := card.CardNumber; i < card.NumberOfMatches+card.CardNumber; i++ {
				// Add a new copy of the original card on the particular stack - this ensures
				// that the pre-computed match count is persisted
				cardStacks[i].Cards = append(cardStacks[i].Cards, cardStacks[i].Cards[0])
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
	findNumRegex := `\d+`
	re := regexp.MustCompile(findNumRegex)
	for scanner.Scan() {
		// Divide into card number and card state
		cardSplit := strings.Split(scanner.Text(), ": ")

		// Find card number
		match := re.FindString(cardSplit[0])
		cardNumber, _ := strconv.Atoi(match)

		// Split into my numbers and winning numbers
		numberSplit := strings.Split(cardSplit[1], " | ")

		myNumbers := re.FindAllString(numberSplit[0], -1)
		winningNumbers := re.FindAllString(numberSplit[1], -1)

		scratchcard := ScratchCard{CardNumber: cardNumber}

		for _, numberStr := range myNumbers {
			value, _ := strconv.Atoi(numberStr)
			scratchcard.MyNumbers = append(scratchcard.MyNumbers, value)
		}

		for _, numberStr := range winningNumbers {
			value, _ := strconv.Atoi(numberStr)
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
	NumberOfMatches int
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
