package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}

func partOne() {
	// In the queue for BFS, store a _pointer_ to the node. Then we can very easily mark it as discovered
	input, _ := readInput("real-input")

	startPoint := findStartPoint(input)

	// Find the surrounding pipes, as we don't know the shape of the beginning pipe
	// They become the beginning of the left and right paths
	pathStarts := []*Node{}

	if len(input) > startPoint.Y + 1 && startPoint.Y-1 >= 0 {
		if _, found := getNextNode(input[startPoint.Y-1][startPoint.X], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y-1][startPoint.X])
		}
		if _, found := getNextNode(input[startPoint.Y+1][startPoint.X], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y+1][startPoint.X])
		}
	}
	
	if len(input[0]) > startPoint.X + 1 && startPoint.X-1 >= 0 {
		if _, found := getNextNode(input[startPoint.Y][startPoint.X-1], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y][startPoint.X-1])
		}
		if _, found := getNextNode(input[startPoint.Y][startPoint.X+1], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y][startPoint.X+1])
		}
	}

	input[startPoint.Y][startPoint.X].Discovered = true
	pathStarts[0].Discovered = true
	pathStarts[1].Discovered = true

	queue := make([]*Node, 0)
	queue = append(queue, pathStarts[0], pathStarts[1])
	stepsTaken := 0
	for len(queue) > 0 {
		current := queue[0]
		current.Discovered = true

		next, found := getNextNode(*current, input)
		if found {
			queue = append(queue, next)
		}

		// Remove the element we just processed
		queue = queue[1:]
		stepsTaken++
	}

	// Divide by two because we were counting every step, so the current steps taken is the whole loop
	fmt.Printf("Part 1: %d", stepsTaken / 2)
}

func partTwo() {
	// In the queue for BFS, store a _pointer_ to the node. Then we can very easily mark it as discovered
	input, _ := readInput("test-input2")

	startPoint := findStartPoint(input)

	// Find the surrounding pipes, as we don't know the shape of the beginning pipe
	// They become the beginning of the left and right paths
	pathStarts := []*Node{}

	if len(input) > startPoint.Y + 1 && startPoint.Y-1 >= 0 {
		if _, found := getNextNode(input[startPoint.Y-1][startPoint.X], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y-1][startPoint.X])
		}
		if _, found := getNextNode(input[startPoint.Y+1][startPoint.X], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y+1][startPoint.X])
		}
	}

	if len(input[0]) > startPoint.X + 1 && startPoint.X-1 >= 0 {
		if _, found := getNextNode(input[startPoint.Y][startPoint.X-1], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y][startPoint.X-1])
		}
		if _, found := getNextNode(input[startPoint.Y][startPoint.X+1], input); found {
			pathStarts = append(pathStarts, &input[startPoint.Y][startPoint.X+1])
		}
	}

	input[startPoint.Y][startPoint.X].Discovered = true
	pathStarts[0].Discovered = true
	pathStarts[1].Discovered = true

	queue := make([]*Node, 0)
	queue = append(queue, pathStarts[0], pathStarts[1])
	// This will mark all of the required things as discovered
	for len(queue) > 0 {
		current := queue[0]
		current.Discovered = true

		next, found := getNextNode(*current, input)
		if found {
			queue = append(queue, next)
		}

		// Remove the element we just processed
		queue = queue[1:]
	}

	// Hard-coded replacement for my input - this isn't great,
	// but I want to solve first
	input[startPoint.Y][startPoint.X].Value = "F"

	print2DArray(input)

	spotsInShape := 0
	for y, row := range input {
		for x, _ := range row {
			// Cast a ray to the right (i.e check the rest of the row)
			// If number of | is even, we're outside the shape
			// If number of | is odd, we're outisde the shape
			numberOfPipes := 0
			for i := x; i < len(row); i++ {
				node := input[y][i]
				if node.Discovered {
					if isBadWallCharacter(node.Value) {
						// If it's seeing a wall that's not the |, it can't
						// be in the shape
						continue
					}

					if node.Value == "|" {
						numberOfPipes++
					}
				}
			}
			// Only count a tile if it has an odd number of pipes
			if numberOfPipes % 2 != 0 {
				spotsInShape++
			}
		}
	}

	fmt.Printf("Part 2: %d", spotsInShape)
}

func isBadWallCharacter(value string) bool {
	return value == "-"
}

// Finds the next node to go to based on the current pipe
// If one of the two potential nodes you can go to has already been discovered,
// we must have discovered it just now, so go the other route.
func getNextNode(currentNode Node, input [][]Node) (*Node, bool) {
	var newNode *Node
	foundNode := true
	
	switch currentNode.Value {
	case "|":
		if !input[currentNode.Position.Y-1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y-1][currentNode.Position.X]
		}
		if !input[currentNode.Position.Y+1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y+1][currentNode.Position.X]
		}
		break
	case "-":
		if !input[currentNode.Position.Y][currentNode.Position.X-1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X-1]
		} else if !input[currentNode.Position.Y][currentNode.Position.X+1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X+1]
		}
		break
	case "L":
		if !input[currentNode.Position.Y-1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y-1][currentNode.Position.X]
		} else if !input[currentNode.Position.Y][currentNode.Position.X+1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X+1]
		}
		break
	case "J":
		if !input[currentNode.Position.Y-1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y-1][currentNode.Position.X]
		} else if !input[currentNode.Position.Y][currentNode.Position.X-1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X-1]
		}
		break
	case "7":
		if !input[currentNode.Position.Y+1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y+1][currentNode.Position.X]
		} else if !input[currentNode.Position.Y][currentNode.Position.X-1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X-1]
		}
		break
	case "F":
		if !input[currentNode.Position.Y+1][currentNode.Position.X].Discovered {
			newNode = &input[currentNode.Position.Y+1][currentNode.Position.X]
		} else if !input[currentNode.Position.Y][currentNode.Position.X+1].Discovered {
			newNode = &input[currentNode.Position.Y][currentNode.Position.X+1]
		}
		break
	case ".":
		foundNode = false
	}

	if newNode == nil {
		foundNode = false
	}
	return newNode, foundNode
}

// Returns the coordinate of the S character, which indicates the start
func findStartPoint(input [][]Node) Point {
	startPoint := Point{}
	for y, row := range input {
		for x, node := range row {
			if node.Value == "S" {
				startPoint.X = x
				startPoint.Y = y
			}
		}
	}

	return startPoint
}

func readInput(filename string) ([][]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	var pipes [][]Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nodes := []Node{}
		for _, character := range scanner.Text() {
			nodes = append(nodes, Node{Value: string(character)})
		}
		pipes = append(pipes, nodes)
	}
	
	for y, row := range pipes {
		for x, _ := range row {
			pipes[y][x].Position = Point{X: x, Y: y}
		}
	}
	
	return pipes, scanner.Err()
}

type Node struct {
	Value      string
	Position   Point
	Discovered bool
}

type Point struct {
	X int
	Y int
}


func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}

func print2DArray(arr [][]Node) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%s ", arr[i][j].Value)
		}
		fmt.Println() // Move to the next line after printing each row
	}
}