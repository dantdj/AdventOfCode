package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	almanac, _ := readInputPartOne("real-input")
	lowestLocation := math.MaxInt

	for _, seed := range almanac.Seeds {
		currentValue := seed
		for _, mappings := range almanac.Mappings {

			result := currentValue
			// Quick explanation of this - the mappings don't overlap, so we can safely go through
			// all of the mappings if we need to. As such, we go through all the mappings, and try and
			// find one where the currentValue maps into a new value. If it doesn't at all, then we
			// return the same number because nothing maps. If result is different from the current value,
			// we know we mapped, so we can just break out
			for _, mappingDetail := range mappings.MappingDetails {
				if result != currentValue {
					// Short-circuit - we mapped to a brand new value already
					break
				}
				result = calculateDestination(currentValue, mappingDetail)
			}

			currentValue = result

		}

		if currentValue < lowestLocation {
			lowestLocation = currentValue
		}
	}

	fmt.Printf("Part One: %d", lowestLocation)
}

func partTwo() {
	almanac, _ := readInputPartTwo("real-input")
	lowestLocation := math.MaxInt

	for _, pairing := range almanac.SeedPairings {
		for i := pairing[0]; i < pairing[0] + pairing[1]; i++ {
			currentValue := i

			for _, mappings := range almanac.Mappings {

				result := currentValue
				for _, mappingDetail := range mappings.MappingDetails {
					if result != currentValue {
						// Short-circuit - we mapped to a brand new value already
						break
					}
					result = calculateDestination(currentValue, mappingDetail)
				}

				currentValue = result
			}

			if currentValue < lowestLocation {
				lowestLocation = currentValue
			}
		}
	}

	fmt.Printf("Part Two: %d", lowestLocation)
}

func readInputPartOne(filename string) (Almanac, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return Almanac{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	findNumRegex := `\d+`
	re := regexp.MustCompile(findNumRegex)
	almanac := Almanac{}
	currentMapInput := [][]int{}
	for scanner.Scan() {
		// Parse seeds
		if strings.Contains(scanner.Text(), "seeds") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			almanac.Seeds = convertStrSliceToIntSlice(numStrs)
			// Skip here - now that the seeds are taken care of, we know that
			// the only other lines of numbers are map values
			continue
		}

		// If we have a blank line, we've reached the end of the current map data,
		// so send it off to be calculated and add it to the almanac, then clear the map tracker
		// and continue to skip the rest of the loop
		if scanner.Text() == "" {
			// The first blank line we hit is right after the seeds, where we have
			// no input, so skip if empty
			if len(currentMapInput) == 0 {
				continue
			}
			almanac.Mappings = append(almanac.Mappings, Mapping{
				MappingDetails: getMappingDetails(currentMapInput),
			})
			currentMapInput = [][]int{}
			continue
		}

		// Otherwise, append the list of ints from the current line into the
		// current map input
		numStrs := re.FindAllString(scanner.Text(), -1)
		// Lines with only strings (like the map labels) will return nothing,
		// so if we get that, skip, as it's not a line we're interested in
		if numStrs == nil {
			continue
		}

		currentMapInput = append(currentMapInput, convertStrSliceToIntSlice(numStrs))
	}

	// Finish by appending the final mapping - this doesn't happen
	// because the loop finishes before we can hit another empty line to run the calculation
	// as normal
	almanac.Mappings = append(almanac.Mappings, Mapping{MappingDetails: getMappingDetails(currentMapInput)})

	return almanac, err
}

func readInputPartTwo(filename string) (Almanac2, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return Almanac2{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	findNumRegex := `\d+`
	re := regexp.MustCompile(findNumRegex)
	almanac := Almanac2{}
	currentMapInput := [][]int{}
	for scanner.Scan() {
		// Parse seeds
		if strings.Contains(scanner.Text(), "seeds") {
			numStrs := re.FindAllString(scanner.Text(), -1)
			seedListing := convertStrSliceToIntSlice(numStrs)
			seedRangePairings := splitIntoPairs(seedListing)

			for _, pairing := range seedRangePairings {
				// pairing[0] is the start of the range, pairing[1] is the number of elements
				fmt.Printf("Adding for initial count %d and range number %d\n", pairing[0], pairing[1])

				almanac.SeedPairings = append(almanac.SeedPairings, pairing)
			}
			// Skip here - now that the seeds are taken care of, we know that
			// the only other lines of numbers are map values
			continue
		}

		// If we have a blank line, we've reached the end of the current map data,
		// so send it off to be calculated and add it to the almanac, then clear the map tracker
		// and continue to skip the rest of the loop
		if scanner.Text() == "" {
			// The first blank line we hit is right after the seeds, where we have
			// no input, so skip if empty
			if len(currentMapInput) == 0 {
				continue
			}
			almanac.Mappings = append(almanac.Mappings, Mapping{
				MappingDetails: getMappingDetails(currentMapInput),
			})
			currentMapInput = [][]int{}
			continue
		}

		// Otherwise, append the list of ints from the current line into the
		// current map input
		numStrs := re.FindAllString(scanner.Text(), -1)
		// Lines with only strings (like the map labels) will return nothing,
		// so if we get that, skip, as it's not a line we're interested in
		if numStrs == nil {
			continue
		}

		currentMapInput = append(currentMapInput, convertStrSliceToIntSlice(numStrs))
	}

	// Finish by appending the final mapping - this doesn't happen
	// because the loop finishes before we can hit another empty line to run the calculation
	// as normal
	almanac.Mappings = append(almanac.Mappings, Mapping{MappingDetails: getMappingDetails(currentMapInput)})

	return almanac, err
}

func getMappingDetails(input [][]int) []MappingDetail {
	details := []MappingDetail{}
	for _, rangeInfo := range input {
		detail := MappingDetail{
			// This is the difference between the destination start and the source start
			SumToAdd:    rangeInfo[0] - rangeInfo[1],
			SourceStart: rangeInfo[1],
			RangeLength: rangeInfo[2],
		}
		details = append(details, detail)
	}

	return details
}

func calculateDestination(num int, detail MappingDetail) int {
	if num >= detail.SourceStart && num < (detail.SourceStart+detail.RangeLength) {
		return num + detail.SumToAdd
	}

	// Number isn't in our map, so just return the number
	return num
}

func splitIntoPairs(nums []int) [][]int {
	var pairs [][]int

	// Jump forward 2 on each iteration so we don't multi-process things
	for i := 0; i < len(nums); i += 2 {
		pair := []int{nums[i], nums[i+1]}
		pairs = append(pairs, pair)
	}

	return pairs
}

func convertStrSliceToIntSlice(input []string) []int {
	nums := make([]int, len(input))
	for i, x := range input {
		nums[i], _ = strconv.Atoi(x)
	}
	return nums
}

type Almanac struct {
	Seeds []int

	Mappings       []Mapping
	MappingDetails []MappingDetail
}

type Almanac2 struct {
	SeedPairings [][]int

	Mappings       []Mapping
}

type Mapping struct {
	MappingDetails []MappingDetail
}

type MappingDetail struct {
	SourceStart int
	// This will be calculated as the difference between the source start
	// and the destination start
	SumToAdd    int
	RangeLength int
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %fs\n", duration.Seconds())
}
