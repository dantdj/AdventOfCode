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
	readInput("test-input")
	
	// Run through recursively - first, replace the first unknown character
	// with a #, then skip one (groups are always separated by a space), and
	// if that next one is unknown, then also try replacing that, and so on
	
}

func partTwo() {
	
}

func readInput(filename string) ([]Record, error){
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	records := []Record{}
	for scanner.Scan() {
		record := Record{}
		splitRecord := strings.Split(scanner.Text(), " ")
		
		for _, character := range splitRecord[0] {
			record.Conditions = append(record.Conditions, string(character))
		}
		
		splitGroup := strings.Split(splitRecord[1], ",")
		for _, groupNumber := range splitGroup {
			value, _ := strconv.Atoi(groupNumber)
			record.DamagedGroups = append(record.DamagedGroups, value)
		}

		records = append(records, record)
	}
	
	return records, nil
}

type Record struct {
	Conditions []string
	DamagedGroups []int
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}