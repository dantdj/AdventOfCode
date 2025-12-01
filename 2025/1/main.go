package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dialValue := 50
	part1Answer := 0
	part2Answer := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		distanceStr := line[1:]
		distance, _ := strconv.Atoi(distanceStr)

		step := 0
		switch direction {
		case 'L':
			step = -1
		case 'R':
			step = 1
		default:
			panic("Invalid direction")
		}

		for range distance {
			if dialValue == 0 {
				part2Answer++
			}
			dialValue = (dialValue + step + 100) % 100
		}

		if dialValue == 0 {
			part1Answer += 1
		}
	}

	fmt.Println("Part 1: ", part1Answer)
	fmt.Println("Part 2: ", part2Answer)
}
