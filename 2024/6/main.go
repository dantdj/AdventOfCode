package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Printf("Part 1: %d\n", partOne())
	fmt.Printf("Part 2: %d\n", partTwo())
}

func partOne() int {
	grid, currentPoint := readInput("real-input.txt")
	// We've always visited the starting square
	visitedCount := 1
	// Always starting pointing up
	currentFacing := "^"
	for {
		if !withinBounds(len(grid[0]), len(grid), currentPoint.X, currentPoint.Y) {
			// If above is true, then we're off the grid, so we're done
			break
		}
		newX, newY := getNextLocation(currentFacing, currentPoint.X, currentPoint.Y)

		if grid[newY][newX].IsObstacle {
			currentFacing = getNextFacing(currentFacing)
			newX, newY = getNextLocation(currentFacing, currentPoint.X, currentPoint.Y)
		}

		currentPoint = Point{X: newX, Y: newY}
		if !grid[newY][newX].Visited {
			visitedCount += 1
			grid[newY][newX].Visited = true
		}
	}

	// for row := 0; row < 10; row++ {
	// 	for column := 0; column < 10; column++ {
	// 		if grid[row][column].IsObstacle {
	// 			fmt.Print("#", " ")
	// 		} else if grid[row][column].Visited {
	// 			fmt.Print("X", " ")
	// 		} else {
	// 			fmt.Print(".", " ")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	return visitedCount
}

func partTwo() int {
	grid, currentPoint := readInput("real-input.txt")
	startPoint := Point{X: currentPoint.X, Y: currentPoint.Y}
	numOfLoops := 0

	// Make a channel with the same size as the total grid,
	// even though we won't be running quite as many simulations as that
	resultsChan := make(chan bool, len(grid)*len(grid))
	var wg sync.WaitGroup

	for yIdx := range grid {
		for xIdx := range grid {
			if startPoint.X == xIdx && startPoint.Y == yIdx {
				// don't bother, we can't set the starting point as an obstacle
				continue
			}
			if grid[yIdx][xIdx].IsObstacle {
				// don't bother, this was an obstacle already so it won't
				// give us anything
				continue
			}

			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				resultsChan <- runRound(grid, x, y, startPoint)
			}(xIdx, yIdx)
		}
	}

	wg.Wait()
	close(resultsChan)
	for result := range resultsChan {
		if result {
			numOfLoops += 1
		}
	}

	// for row := 0; row < 10; row++ {
	// 	for column := 0; column < 10; column++ {
	// 		if grid[row][column].IsObstacle {
	// 			fmt.Print("#", " ")
	// 		} else if grid[row][column].Visited {
	// 			fmt.Print("X", " ")
	// 		} else {
	// 			fmt.Print(".", " ")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	return numOfLoops
}

// runs a round and returns if a loop was present for that round
func runRound(grid [][]Point, xIdx, yIdx int, startPoint Point) bool {
	newGrid := getGridWithObstacle(grid, xIdx, yIdx)
	fmt.Printf("Running for X: %d, Y: %d\n", xIdx, yIdx)
	loopFound := false
	outOfBounds := false
	currentPoint := startPoint
	// Always starting pointing up
	currentFacing := "^"

	visitedPoints := []PointWithFacing{}

	visitedPoints = append(visitedPoints, PointWithFacing{
		Point:  currentPoint,
		Facing: currentFacing,
	})
	for {
		for i := 1; i <= 50000; i++ {
			if !withinBounds(len(newGrid[0]), len(newGrid), currentPoint.X, currentPoint.Y) {
				// If above is true, then we're off the grid,
				// so definitely not in a loop
				outOfBounds = true
				break
			}
			newX, newY := getNextLocation(currentFacing, currentPoint.X, currentPoint.Y)

			if newGrid[newY][newX].IsObstacle {
				currentFacing = getNextFacing(currentFacing)
				newX, newY = getNextLocation(currentFacing, currentPoint.X, currentPoint.Y)
			}

			currentPoint = Point{X: newX, Y: newY}
			visitedPoints = append(visitedPoints, PointWithFacing{
				Point:  currentPoint,
				Facing: currentFacing,
			})
		}

		if outOfBounds {
			break
		}

		for idx1, point := range visitedPoints {
			samePointFound := false
			for idx2, otherPoint := range visitedPoints {
				if idx1 == idx2 {
					// don't check a point against itself
					continue
				}

				if point.Facing == otherPoint.Facing && point.Point.X == otherPoint.Point.X && point.Point.Y == otherPoint.Point.Y {
					samePointFound = true
					break
				}
			}

			if samePointFound {
				loopFound = true
				break
			}
		}

		if loopFound {
			return true
		}
	}

	return false
}

// Returns a copy of the grid with an obstacle inserted into the x and y coord
func getGridWithObstacle(grid [][]Point, x, y int) [][]Point {
	duplicate := make([][]Point, len(grid))
	for i := range grid {
		duplicate[i] = make([]Point, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	duplicate[y][x].IsObstacle = true

	return duplicate
}

// Given the current facing and location, returns the next index to move to
func getNextLocation(currentFacing string, currentX, currentY int) (int, int) {
	if currentFacing == "^" {
		return currentX, currentY - 1
	} else if currentFacing == ">" {
		return currentX + 1, currentY
	} else if currentFacing == "v" {
		return currentX, currentY + 1
	} else {
		// currentFacing is <
		return currentX - 1, currentY
	}
}

func getNextFacing(currentFacing string) string {
	if currentFacing == "^" {
		return ">"
	} else if currentFacing == ">" {
		return "v"
	} else if currentFacing == "v" {
		return "<"
	} else {
		// currentFacing is <
		return "^"
	}
}

func withinBounds(maxX, maxY, currentX, currentY int) bool {
	return currentX+1 < maxX && currentX-1 >= 0 && currentY+1 < maxY && currentY-1 >= 0
}

// Returns point grid and location of starting point
func readInput(filename string) ([][]Point, Point) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, Point{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	listOfSlices := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}

		listOfSlices = append(listOfSlices, chars)
	}

	pointGrid := [][]Point{}
	// We start off having visited the start point
	startPoint := Point{}

	for yIndex := range listOfSlices {
		pointRow := []Point{}
		for xIndex := range listOfSlices[yIndex] {
			point := Point{
				X:          xIndex,
				Y:          yIndex,
				Visited:    false,
				IsObstacle: false,
			}

			if listOfSlices[yIndex][xIndex] == "^" {
				point.Visited = true
				startPoint = point

			}

			if listOfSlices[yIndex][xIndex] == "#" {
				point.IsObstacle = true
			}

			pointRow = append(pointRow, point)
		}
		pointGrid = append(pointGrid, pointRow)
	}

	return pointGrid, startPoint
}

type Point struct {
	X          int
	Y          int
	Visited    bool
	IsObstacle bool
}

type PointWithFacing struct {
	Point  Point
	Facing string
}
