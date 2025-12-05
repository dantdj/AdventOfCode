package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	start := time.Now()
	fmt.Printf("Part One: %d (%s)\n", partOne(grid), time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d (%s)\n", partTwo(grid), time.Since(start))
}

func partOne(grid [][]string) int {
	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "@" {
				if CheckNumAdjacent(grid, i, j) < 4 {
					answer++
				}
			}
		}
	}

	return answer
}

func partTwo(grid [][]string) int {
	// Stuff going on in this function will modify the grid in main,
	// as Go slices are semi-pass-by-reference. (They are a reference
	// to the underlying memory location, even though the slice itself is passed by value.)
	// This would be tricky if I needed to use it later on in main.
	// Fortunately, I don't need to! I could just create a copy of the grid
	// in main, and pass that to the function if I wanted to though.
	answer := 0
	for {
		removedPaperThisRound := false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "@" {
					if CheckNumAdjacent(grid, i, j) < 4 {
						grid[i][j] = "."
						answer++
						removedPaperThisRound = true
					}
				}
			}
		}
		if !removedPaperThisRound {
			break
		}
	}

	return answer
}

func CheckNumAdjacent(grid [][]string, r, c int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	adjacent := 0
	for i := 0; i < 8; i++ {
		nr, nc := r+dr[i], c+dc[i]

		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == "@" {
				adjacent++
			}
		}
	}
	return adjacent
}
