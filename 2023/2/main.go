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
	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14

	input, _ := readInput("real-input")
	result := 0
	for _, game := range input {
		possibleGame := true
		for _, round := range game.Rounds {
			if round.RedCubes > maxRedCubes || round.GreenCubes > maxGreenCubes || round.BlueCubes > maxBlueCubes {
				possibleGame = false
				break
			}
		}

		if possibleGame {
			result += game.GameId
		}
	}

	fmt.Printf("Part One: %d", result)
}

func partTwo() {
	input, _ := readInput("real-input")
	result := 0
	
	for _, game := range input {
		requiredRedCubes := 0
		requiredGreenCubes := 0
		requiredBlueCubes := 0
		
		for _, round := range game.Rounds {
			if round.RedCubes > requiredRedCubes {
				requiredRedCubes = round.RedCubes
			}
			
			if round.GreenCubes > requiredGreenCubes {
				requiredGreenCubes = round.GreenCubes
			}
			
			if round.BlueCubes > requiredBlueCubes {
				requiredBlueCubes = round.BlueCubes
			}
		}
		
		result += requiredRedCubes * requiredGreenCubes * requiredBlueCubes
	}
	
	fmt.Printf("Part Two: %d", result)
}

func readInput(filename string) ([]GameRecord, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	var records []GameRecord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameRecord := GameRecord{}
		// Split by colon to get the full game ID, then split the first string there by space to get just the ID number itself
		splitString := strings.Split(scanner.Text(), ":")
		details := strings.Split(splitString[0], " ")
		gameId, _ := strconv.Atoi(details[1])
		gameRecord.GameId = gameId

		// Split the second string in the previous section by semi-colon to get each individual round in the game
		rounds := strings.Split(splitString[1], ";")

		// Split by comma to get each individual block, trim to remove leading and trailing spaces,
		// then split by spaces to get the number of blocks and the block colour
		for _, roundStr := range rounds {
			round := GameRound{}
			roundStr := strings.TrimSpace(roundStr)
			blocks := strings.Split(roundStr, ", ")
			for _, block := range blocks {
				blockDetails := strings.Split(block, " ")
				blockCount, _ := strconv.Atoi(blockDetails[0])
				switch blockDetails[1] {
				case "blue":
					round.BlueCubes = blockCount
				case "red":
					round.RedCubes = blockCount
				case "green":
					round.GreenCubes = blockCount
				default:
					fmt.Printf("Couldn't match block: %s", blockDetails[1])
				}
			}

			gameRecord.Rounds = append(gameRecord.Rounds, round)
		}

		records = append(records, gameRecord)
	}

	return records, scanner.Err()
}

type GameRecord struct {
	GameId int
	Rounds []GameRound
}

type GameRound struct {
	BlueCubes  int
	RedCubes   int
	GreenCubes int
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
