package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// read input
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")

	minMaxRanges := make([][]int, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, "-")
		min, _ := strconv.Atoi(splitLine[0])
		max, _ := strconv.Atoi(splitLine[1])

		minMaxRanges = append(minMaxRanges, []int{min, max})
	}

	var ingredients []int
	for _, line := range strings.Split(parts[1], "\n") {
		if line == "" {
			continue
		}
		val, err := strconv.Atoi(line)
		if err == nil {
			ingredients = append(ingredients, val)
		}
	}

	start := time.Now()
	fmt.Printf("Part One: %d (%s)\n", partOne(minMaxRanges, ingredients), time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d (%s)\n", partTwo(minMaxRanges), time.Since(start))
}

func partOne(ranges [][]int, ingredients []int) int {
	answer := 0
	for _, ingredient := range ingredients {
		for _, rng := range ranges {
			if ingredient >= rng[0] && ingredient <= rng[1] {
				answer++
				break
			}
		}
	}
	return answer
}

func partTwo(ranges [][]int) int {
	answer := 0

	// Sort range list by starting numbers to make
	// merging easier
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Merge ranges by seeing if they overlap,
	// and if they do, merge them into a single range
	mergedRanges := make([][]int, 0)
	for _, rng := range ranges {
		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, rng)
		} else {
			last := mergedRanges[len(mergedRanges)-1]
			if rng[0] <= last[1] {
				mergedRanges[len(mergedRanges)-1] = []int{last[0], max(rng[1], last[1])}
			} else {
				mergedRanges = append(mergedRanges, rng)
			}
		}
	}

	for _, rng := range mergedRanges {
		// The +1 ensures we treat the range as inclusive
		// It's a bit of a hack maybe, but it works
		ingredientCount := (rng[1] - rng[0]) + 1
		answer += ingredientCount
	}
	return answer
}
