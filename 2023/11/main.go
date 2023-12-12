package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}

func partOne() {
	galaxies, _ := readInput("real-input", false)
	
	// Get pairs of galaxies - only count each pair once
	// This might be easier if each character in the array
	// is a `Galaxy` type, which stores a list of its paired elements.
	// Then, when adding a pair, check the list of the other element involved
	// in the pair. If it's already present, then skip adding so we avoid double-counting
	galaxyCoordMap := map[Point][]Point{}
	for _, galaxy := range galaxies {
		galaxyCoordMap[galaxy] = []Point{}
		
		for _, otherGalaxy := range galaxies {
			if galaxy.X == otherGalaxy.X && galaxy.Y == otherGalaxy.Y {
				continue
			}
			
			if _, existingValue := galaxyCoordMap[otherGalaxy]; existingValue {
				continue
			}
			galaxyCoordMap[galaxy] = append(galaxyCoordMap[galaxy], otherGalaxy)
		}
	}
	
	total := 0
	
	for k, v := range galaxyCoordMap {
		for _, pair := range v {
			total += getDistanceBetweenTwoPoints(k, pair)
		}
	}
	
	fmt.Printf("Part 1: %d", total)
}

func partTwo() {
	galaxies, _ := readInput("test-input", true)

	// Get pairs of galaxies - only count each pair once
	// This might be easier if each character in the array
	// is a `Galaxy` type, which stores a list of its paired elements.
	// Then, when adding a pair, check the list of the other element involved
	// in the pair. If it's already present, then skip adding so we avoid double-counting
	galaxyCoordMap := map[Point][]Point{}
	for _, galaxy := range galaxies {
		galaxyCoordMap[galaxy] = []Point{}

		for _, otherGalaxy := range galaxies {
			if galaxy.X == otherGalaxy.X && galaxy.Y == otherGalaxy.Y {
				continue
			}

			if _, existingValue := galaxyCoordMap[otherGalaxy]; existingValue {
				continue
			}
			galaxyCoordMap[galaxy] = append(galaxyCoordMap[galaxy], otherGalaxy)
		}
	}

	total := 0

	for k, v := range galaxyCoordMap {
		for _, pair := range v {
			total += getDistanceBetweenTwoPoints(k, pair)
		}
	}

	fmt.Printf("Part 2: %d", total)
}

func expandUniverse(universe [][]Node) [][]Node {
	indexesToAddColumnsAt := []int{}
	indexesToAddRowsAt := []int{}
	
	for i, row := range universe {
		isEmptyRow := true
		for _, value := range row {
			if value.IsGalaxy {
				isEmptyRow = false
				break
			}
		}

		if isEmptyRow {
			indexesToAddRowsAt = append(indexesToAddRowsAt, i)
		}
	}

	if len(universe) > 0 {
		for j := range universe[0] {
			isEmptyColumn := true
			for i := range universe {
				if universe[i][j].IsGalaxy {
					isEmptyColumn = false
					break
				}
			}

			if isEmptyColumn {
				indexesToAddColumnsAt = append(indexesToAddColumnsAt, j)
			}
		}
	}
	
	for i, index := range indexesToAddRowsAt {
		// +i because each new row added means we're offset one more from our original discovery
		universe = insertEmptyRow(universe, index+i)
	}
	
	for i, index := range indexesToAddColumnsAt {
		// +i because each new row added means we're offset one more from our original discovery
		universe = insertEmptyColumn(universe, index+i)
	}

	return universe
}

func expandUniverse2(universe [][]Node) [][]Node {
	indexesToAddColumnsAt := []int{}
	indexesToAddRowsAt := []int{}

	for i, row := range universe {
		isEmptyRow := true
		for _, value := range row {
			if value.IsGalaxy {
				isEmptyRow = false
				break
			}
		}

		if isEmptyRow {
			indexesToAddRowsAt = append(indexesToAddRowsAt, i)
		}
	}

	if len(universe) > 0 {
		for j := range universe[0] {
			isEmptyColumn := true
			for i := range universe {
				if universe[i][j].IsGalaxy {
					isEmptyColumn = false
					break
				}
			}

			if isEmptyColumn {
				indexesToAddColumnsAt = append(indexesToAddColumnsAt, j)
			}
		}
	}
	
	
	extraRowsOrColumns := 1000
	for i, index := range indexesToAddRowsAt {
		// +i*offset because we've created that number of extra rows compared to our discovery
		// i*offset is how many extra rows we've already created
		for j:= 0; j < extraRowsOrColumns; j++ {
			universe = insertEmptyRow(universe, index+(i*extraRowsOrColumns)+i)
		}
	}

	for i, index := range indexesToAddColumnsAt {
		// +i*offset because we've created that number of extra columns compared to our discovery
		// Our original indexes were based on the beginning grid, so we need to increase them to take that into account
		for j := 0; j < extraRowsOrColumns; j++ {
			universe = insertEmptyColumn(universe, index+(i*extraRowsOrColumns)+i)
		}
	}

	return universe
}

func insertEmptyColumn(matrix [][]Node, columnIndex int) [][]Node {
	if columnIndex < 0 || columnIndex > len(matrix[0]) {
		fmt.Println("Column index out of bounds.")
		return matrix
	}
	
	newMatrix := make([][]Node, len(matrix))
	for i := range newMatrix {
		newRow := make([]Node, len(matrix[0])+1)

		copy(newRow[:columnIndex], matrix[i][:columnIndex])
		newRow[columnIndex] = Node{}
		copy(newRow[columnIndex+1:], matrix[i][columnIndex:])
		
		newMatrix[i] = newRow
	}

	return newMatrix
}

func insertEmptyRow(matrix [][]Node, rowIndex int) [][]Node {
	if rowIndex < 0 || rowIndex > len(matrix) {
		fmt.Println("Row index out of bounds.")
		return matrix
	}
	
	row := []Node{}
	for range matrix[0] {
		row = append(row, Node{})
	}

	newMatrix := make([][]Node, len(matrix)+1)

	copy(newMatrix[:rowIndex], matrix[:rowIndex])
	newMatrix[rowIndex] = row
	copy(newMatrix[rowIndex+1:], matrix[rowIndex:])

	return newMatrix
}

// Gets the Manhattan Distance between two points in the provided 2D array
func getDistanceBetweenTwoPoints(point1, point2 Point) int {
	return int(math.Abs(float64(point1.X) - float64(point2.X)) + math.Abs(float64(point1.Y) - float64(point2.Y)))
}

type Node struct {
	IsGalaxy bool
	Pairs []Point
}

type Point struct {
	X int
	Y int
}

func readInput(filename string, part2 bool) ([]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	universe := [][]Node{}
	for scanner.Scan() {
		mapValues := strings.Split(scanner.Text(), "")
		row := []Node{}
		for _, value := range mapValues {
			if value == "#" {
				row = append(row, Node{IsGalaxy: true})
			} else {
				row = append(row, Node{IsGalaxy: false})
			}
		}
		universe = append(universe, row)
	}
	
	//universe = expandUniverse(universe)
	if part2 {
		universe = expandUniverse2(universe)
	} else {
		universe = expandUniverse(universe)
	}
	
	galaxies := []Point{}
	
	for y, row := range universe {
		for x, _ := range row {
			if universe[y][x].IsGalaxy {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}

	return galaxies, nil
}


func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}