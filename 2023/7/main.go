package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var partOneCardValues = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4,
	"3": 3, "2": 2,
}

var partTwoCardValues = map[string]int{
	"A": 14, "K": 13, "Q": 12, "T": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4,
	"3": 3, "2": 2, "J": 0,
}

const (
	HighCard     = iota
	OnePair      = iota
	TwoPair      = iota
	ThreeOfAKind = iota
	FullHouse    = iota
	FourOfAKind  = iota
	FiveOfAKind  = iota
)

func main() {
	measureTime(partOne)
	
	// Incorrect answers
	// 246198085 - too high
	// 246750515 - too high
	measureTime(partTwo)
}

func partOne() {
	input, _ := readInput("real-input")

	// Add type of hand to each card - I could do this at the input parsing step too, but there's not a huge number
	// of elements so I'm gonna prefer the separation
	for i, hand := range input {
		input[i].Type = findHandTypePartOne(hand)
	}

	// Sort the cards using our custom function
	sort.Slice(input, func(i, j int) bool {
		// This is a _less_ function, so i has to be worse than j
		return hand1WorseThanHand2PartOne(input[i], input[j])
	})

	// Iterate through all the cards, and do bid * rank
	// to get final result. Should be fine to do a direct iteration
	// because we sort the list above
	winnings := 0
	for i, card := range input {
		// Ranks start at 1, so add one to 1
		winnings += card.Bid * (i + 1)
	}

	fmt.Printf("Part One: %d", winnings)
}

func partTwo() {
	input, _ := readInput("real-input")

	// Add type of hand to each card - I could do this at the input parsing step too, but there's not a huge number
	// of elements so I'm gonna prefer the separation
	for i, hand := range input {
		input[i].Type = findHandTypePartTwo(hand)
	}

	sort.Slice(input, func(i, j int) bool {
		// This is a _less_ function, so i has to be worse than j
		return hand1WorseThanHand2PartTwo(input[i], input[j])
	})

	winnings := 0
	for i, card := range input {
		// Ranks start at 1, so add one to 1
		winnings += card.Bid * (i + 1)
	}

	fmt.Printf("Part Two: %d", winnings)
}

// Used for disambiguating cards that are the same type
func hand1WorseThanHand2PartOne(hand1, hand2 CardHand) bool {
	if hand1.Type > hand2.Type {
		// Hand 1 immediately wins if its type is higher
		return false
	} else if hand2.Type > hand1.Type {
		// Hand 2 immmediately wins if its type is higher, so hand 1 doesn't beat hand 2
		return true
	}

	// If hand 1 didn't immediately lose, we have two cards the same type, so
	// let's go through the cards...
	// Both hands are the same length, so I can just go on the length of hand 1
	for i, _ := range hand1.Cards {
		if partOneCardValues[hand1.Cards[i].Value] < partOneCardValues[hand2.Cards[i].Value] {
			// If the card value of a card in hand 1 is lower
			// than hand 2's, it loses
			return true
		} else if partOneCardValues[hand1.Cards[i].Value] == partOneCardValues[hand2.Cards[i].Value] {
			// Skip this, we can't use it to tell if hand1 is worse or better
			continue
		} else {
			return false
		}
	}

	// We shouldn't be able to reach here, so log a message if we do...
	fmt.Println("What's happened here?")
	return true
}

// Used for disambiguating cards that are the same type
func hand1WorseThanHand2PartTwo(hand1, hand2 CardHand) bool {
	if hand1.Type == hand2.Type{
		//fmt.Println("thing of interest")
	}
	if hand1.Type > hand2.Type {
		// Hand 1 immediately wins if its type is higher
		return false
	} else if hand2.Type > hand1.Type {
		// Hand 2 immmediately wins if its type is higher, so hand 1 doesn't beat hand 2
		return true
	}

	// If hand 1 didn't immediately lose, we have two cards the same type, so
	// let's go through the cards...
	// Both hands are the same length, so I can just go on the length of hand 1
	for i, _ := range hand1.Cards {
		if partTwoCardValues[hand1.Cards[i].Value] < partTwoCardValues[hand2.Cards[i].Value] {
			// If the card value of a card in hand 1 is lower
			// than hand 2's, it loses
			return true
		} else if partTwoCardValues[hand1.Cards[i].Value] == partTwoCardValues[hand2.Cards[i].Value] {
			// Skip this, we can't use it to tell if hand1 is worse or better
			continue
		} else {
			return false
		}
	}

	// We shouldn't be able to reach here, so log a message if we do...
	fmt.Println("What's happened here?")
	return true
}

// Determines the type (e.g five of a kind, two pair, etc) of a hand
func findHandTypePartOne(hand CardHand) int {
	// At most, each hand can have 2 pairs of matching cards. Therefore, for each card, create two new "stacks"
	// that contain any cards that had the same value as another card.
	// If both are populated, check if it's a full house or two pair. If one is length 3 and the other length 2 it's full house
	// if not then it's two pair
	// If only one is populated, check the length. If it's 5, five of a kind, if it's 4, four of a kind, 3, three of a kind, 2, one pair, etc
	// We can just iterate through at that point because we've caught the special cases already
	// If none get matched, then it's high card

	labels := map[int][]int{}

	// Build the value stacks
	for _, card := range hand.Cards {
		value := partOneCardValues[card.Value]
		if _, found := labels[value]; found {
			labels[value] = append(labels[value], value)
		} else {
			labels[value] = []int{value}
		}
	}

	stacksWithMultipleNumbers := []int{}
	for k, v := range labels {
		if len(v) > 1 {
			stacksWithMultipleNumbers = append(stacksWithMultipleNumbers, k)
		}
	}

	// Ideally improve this, so many conditions...
	if len(stacksWithMultipleNumbers) == 2 {
		// If they're equal, we have two pairs
		if len(labels[stacksWithMultipleNumbers[0]]) == len(labels[stacksWithMultipleNumbers[1]]) {
			return TwoPair
		}

		// Check for full house
		if len(labels[stacksWithMultipleNumbers[0]]) == 3 && len(labels[stacksWithMultipleNumbers[1]]) == 2 ||
			len(labels[stacksWithMultipleNumbers[0]]) == 2 && len(labels[stacksWithMultipleNumbers[1]]) == 3 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}

	if len(stacksWithMultipleNumbers) == 1 {
		switch len(labels[stacksWithMultipleNumbers[0]]) {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			return ThreeOfAKind
		case 2:
			return OnePair
		}
	}

	return HighCard
}

// Determines the type (e.g five of a kind, two pair, etc) of a hand
func findHandTypePartTwo(hand CardHand) int {
	// If a card contains a joker, iterate through all of the possible card values
	// and find the result that gives the best hand type
	// Naively, I think the best hand types are when there's more same numbers, so
	// I'm going to say try replacing the joker with each other card type and see what type it returns
	// and take the highest type

	// Build the value stacks
	jokerValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	bestType := HighCard
	
	for _, jokerValue := range jokerValues {
		labels := map[int][]int{}
		for _, card := range hand.Cards {
			value := 0
			// Replace the joker with the value we're currently trying
			if card.Value == "J" {
				value = partTwoCardValues[jokerValue]
			} else {
				value = partTwoCardValues[card.Value]
			}
			
			if _, found := labels[value]; found {
				labels[value] = append(labels[value], value)
			} else {
				labels[value] = []int{value}
			}
		}

		stacksWithMultipleNumbers := []int{}
		for k, v := range labels {
			if len(v) > 1 {
				stacksWithMultipleNumbers = append(stacksWithMultipleNumbers, k)
			}
		}

		if len(stacksWithMultipleNumbers) == 2 {
			if len(labels[stacksWithMultipleNumbers[0]]) == len(labels[stacksWithMultipleNumbers[1]]) {
				bestType = updateBestTypeIfHigher(bestType, TwoPair)
				continue
			}

			if len(labels[stacksWithMultipleNumbers[0]]) == 3 && len(labels[stacksWithMultipleNumbers[1]]) == 2 ||
				len(labels[stacksWithMultipleNumbers[0]]) == 2 && len(labels[stacksWithMultipleNumbers[1]]) == 3 {
				bestType = updateBestTypeIfHigher(bestType, FullHouse)
				continue
			} else {
				bestType = updateBestTypeIfHigher(bestType, ThreeOfAKind)
				continue
			}
		}

		if len(stacksWithMultipleNumbers) == 1 {
			switch len(labels[stacksWithMultipleNumbers[0]]) {
			case 5:
				bestType = updateBestTypeIfHigher(bestType, FiveOfAKind)
				continue
			case 4:
				bestType = updateBestTypeIfHigher(bestType, FourOfAKind)
				continue
			case 3:
				bestType = updateBestTypeIfHigher(bestType, ThreeOfAKind)
				continue
			case 2:
				bestType = updateBestTypeIfHigher(bestType, OnePair)
				continue
			}
		}
	}

	return bestType
}

func updateBestTypeIfHigher(current, new int) int {
	if current < new {
		return new
	}

	return current
}

func readInput(filename string) ([]CardHand, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return []CardHand{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []CardHand{}
	for scanner.Scan() {
		hand := CardHand{}
		state := strings.Split(scanner.Text(), " ")
		for _, character := range state[0] {
			hand.Cards = append(hand.Cards, Card{Value: string(character)})
		}

		bid, _ := strconv.Atoi(state[1])
		hand.Bid = bid

		hands = append(hands, hand)
	}
	return hands, nil
}

type CardHand struct {
	Cards []Card
	Bid   int
	// The type of hand it is, starting at 0 for 5 of a kind,
	// through to 6 for high card.
	Type int
}

type Card struct {
	Value string
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
