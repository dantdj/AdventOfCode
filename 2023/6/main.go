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
	races, _ := readInputPartOne("real-input")
	
	// Start at 1, because we're multiplying numbers in here, and
	// starting at 0 would mean we never get anywhere
	result := 1
	
	for _, race := range races {
		distances := []int{}
		winningDistances := 0
		currentSpeed := 0
		for i := 0; i < race.Duration; i++ {
			// Append the distance travelled in the remainder of the time in the race
			distances = append(distances, currentSpeed * (race.Duration - i))
			currentSpeed++
		}
		
		for _, distance := range distances {
			if distance > race.RecordDistance {
				winningDistances++
			}
		}
		
		result *= winningDistances
	}
	
	fmt.Printf("Part One: %d", result)
}

func partTwo() {
	race, _ := readInputPartTwo("real-input")
	
	distances := []int{}
	winningDistances := 0
	currentSpeed := 0
	for i := 0; i < race.Duration; i++ {
		// Append the distance travelled in the remainder of the time in the race
		distances = append(distances, currentSpeed * (race.Duration - i))
		currentSpeed++
	}

	for _, distance := range distances {
		if distance > race.RecordDistance {
			winningDistances++
		}
	}

	result := winningDistances
	
	fmt.Printf("Part Two: %d", result)
}

func readInputPartOne(filename string) ([]Race, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return []Race{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	findNumRegex := `\d+`
	re := regexp.MustCompile(findNumRegex)
	
	// This is going to be a naive parsing that goes through the entire input,
	// then runs through and creates the pairing after. Could be a performance impact
	var (
		times []int
		distances []int
	)
	
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Time") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			times = convertStrSliceToIntSlice(numStrs)
		}
		
		if strings.Contains(scanner.Text(), "Distance") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			distances = convertStrSliceToIntSlice(numStrs)
		}
	}
	
	races := []Race{}
	
	for i := 0; i < len(times); i++ {
		races = append(races, Race{Duration: times[i], RecordDistance: distances[i]})
	}
	
	return races, nil
}


func readInputPartTwo(filename string) (Race, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return Race{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	findNumRegex := `\d+`
	re := regexp.MustCompile(findNumRegex)

	// This is going to be a naive parsing that goes through the entire input,
	// then runs through and creates the pairing after. Could be a performance impact

	race := Race{}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Time") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			finalStr := ""
			for _, num := range numStrs {
				finalStr += num
			}
			race.Duration, _ = strconv.Atoi(finalStr) 
		}

		if strings.Contains(scanner.Text(), "Distance") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			finalStr := ""
			for _, num := range numStrs {
				finalStr += num
			}
			race.RecordDistance, _ = strconv.Atoi(finalStr) 
		}
	}

	return race, nil
}

type Race struct {
	Duration int
	RecordDistance int
}

func convertStrSliceToIntSlice(input []string) []int {
	nums := make([]int, len(input))
	for i, x := range input {
		nums[i], _ = strconv.Atoi(x)
	}
	return nums
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
