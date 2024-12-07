package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	rules, updates := readInput("real-input.txt")
	total := 0

	for _, update := range updates {
		// get rules related to this particular update
		updateRules := []PageOrderingRule{}
		for _, page := range update {
			for _, rule := range rules {
				if page == rule.First {
					updateRules = addUniqueRule(updateRules, rule)
				}
			}
		}

		orderMap := make(map[int][]int)
		for _, rule := range updateRules {
			orderMap[rule.First] = append(orderMap[rule.First], rule.Second)
		}
		updateCopy := []int{}
		updateCopy = append(updateCopy, update...)

		sort.Slice(update, func(i, j int) bool {
			for _, after := range orderMap[update[i]] {
				if update[j] == after {
					return true
				}
			}
			return false
		})

		valid := areSlicesEqual(update, updateCopy)

		if valid {
			total += getMiddleElement(update)
		}
	}

	return total
}

func partTwo() int {
	rules, updates := readInput("real-input.txt")
	total := 0

	for _, update := range updates {
		// get rules related to this particular update
		updateRules := []PageOrderingRule{}
		for _, page := range update {
			for _, rule := range rules {
				if page == rule.First {
					updateRules = addUniqueRule(updateRules, rule)
				}
			}
		}

		orderMap := make(map[int][]int)
		for _, rule := range updateRules {
			orderMap[rule.First] = append(orderMap[rule.First], rule.Second)
		}
		updateCopy := []int{}
		updateCopy = append(updateCopy, update...)

		sort.Slice(update, func(i, j int) bool {
			for _, after := range orderMap[update[i]] {
				if update[j] == after {
					return true
				}
			}
			return false
		})

		valid := areSlicesEqual(update, updateCopy)

		if !valid {
			total += getMiddleElement(update)
		}
	}

	return total
}

func addUniqueRule(rules []PageOrderingRule, rule PageOrderingRule) []PageOrderingRule {
	for _, p := range rules {
		if p == rule {
			// nothing to do, this rule already exists
			return rules
		}
	}

	return append(rules, rule)
}

func getMiddleElement(slice []int) int {
	return slice[len(slice)/2]
}

func areSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) { //if condition is not satisfied print false
		return false
	}
	for i, element := range slice1 { // use for loop to check equality
		if element != slice2[i] {
			return false
		}
	}
	return true
}

func readInput(filename string) ([]PageOrderingRule, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pageOrderingRules := []PageOrderingRule{}
	listOfSlices := [][]int{}

	// Kinda hacky, but less focused on reading input here
	sectionSplitReached := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sectionSplitReached = true
			continue
		}

		if !sectionSplitReached {
			rule := PageOrderingRule{}
			parts := strings.Split(line, "|")
			rule.First, _ = strconv.Atoi(parts[0])
			rule.Second, _ = strconv.Atoi(parts[1])
			pageOrderingRules = append(pageOrderingRules, rule)
		} else {
			nums := []int{}
			for _, numStr := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(numStr)
				nums = append(nums, num)
			}
			listOfSlices = append(listOfSlices, nums)
		}
	}

	return pageOrderingRules, listOfSlices
}

type PageOrderingRule struct {
	First  int
	Second int
}
