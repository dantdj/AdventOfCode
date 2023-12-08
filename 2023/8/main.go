package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	measureTime(partOne)
	measureTime(partTwo)
}

func partOne() {
	tree, _ := readInput("real-input")
	currentNodeId := "AAA"
	destinationNodeId := "ZZZ"
	stepsTaken := 0
	for !tree.FoundDestination {
		for _, direction := range tree.Directions {
			if currentNodeId == destinationNodeId {
				tree.FoundDestination = true
				break
			}
			switch string(direction) {
			case "L":
				currentNodeId = tree.Nodes[currentNodeId].Left
			case "R":
				currentNodeId = tree.Nodes[currentNodeId].Right
			}
			stepsTaken++
		}
	}

	fmt.Printf("Part One: %d", stepsTaken)
}

// Part two essentially has multiple loops inside the input, each one starting at a node ID that ends with A.
// Each loop will take a fixed amount of time to go around - as such, we can mark how many steps it takes to complete
// each loop (I called it depth here, but really it's more like a circumference), add each one to an array
// and then find the lowest common multiple of them. E.g for ones of size 2, 3, and 4, to find the earliest point where _all_
// loops will coincide on their closing point (a node ID ending in Z), you find the lowest number that _each_ number
// can multiply into without a remainder. For 2,3,4, that number is 16. You will have done 16 steps to find the point where they
// all have an ID that ends in Z.
func partTwo() {
	tree, _ := readInput("real-input")

	startingNodes := []Node{}
	for k, v := range tree.Nodes {
		if strings.HasSuffix(k, "A") {
			startingNodes = append(startingNodes, v)
		}
	}

	currentNodeIds := map[string]string{}
	for _, node := range startingNodes {
		currentNodeIds[node.ID] = node.ID
	}

	stepsTaken := 0
	endDepths := []int64{}

	foundAllEnds := false 
	for !foundAllEnds {
		for _, direction := range tree.Directions {
			numAtEnd := 0
			
			for _, id := range currentNodeIds {
				node := tree.Nodes[id]
				
				if node.EndDepth > 0 {
					// Skip processing, we're done
					numAtEnd++
					continue
				}
				
				if strings.HasSuffix(node.ID, "Z") {
					// Found the end depth for this particular node
					node.EndDepth = stepsTaken
					endDepths = append(endDepths, int64(stepsTaken))
					tree.Nodes[id] = node
					numAtEnd++
					continue
				}
				
				// Remove current node from the list, as we'll replace it with the next node
				// down in a sec
				delete(currentNodeIds, id)
				switch string(direction) {
				case "L":
					currentNodeIds[node.Left] = node.Left
					break
				case "R":
					currentNodeIds[node.Right] = node.Right
					break
				}
			}
			
			if numAtEnd == len(currentNodeIds) {
				foundAllEnds = true
				break
			}
			
			stepsTaken++
		}
	}
	
	fmt.Printf("Part Two: %v", findLowestCommonMultiple(endDepths))
}

// I'll be honest, I just googled this one, it's not how I'd usually write things
func findLowestCommonMultiple(arr []int64) *big.Int {
	if len(arr) < 2 {
		panic("Array must contain at least two elements")
	}

	result := big.NewInt(arr[0])

	for i := 1; i < len(arr); i++ {
		gcd := new(big.Int).GCD(nil, nil, result, big.NewInt(arr[i]))
		result = result.Mul(result, big.NewInt(arr[i])).Div(result, gcd)
	}

	return result
}

func readInput(filename string) (Tree, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return Tree{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tree := Tree{Nodes: map[string]Node{}}
	findTreeNodesRegex := `\w{3}`
	re := regexp.MustCompile(findTreeNodesRegex)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if !strings.Contains(scanner.Text(), "=") {
			// String didn't contain an equals sign, and it's not empty
			// so it has to be the directions
			tree.Directions = scanner.Text()
			continue
		}

		// Just gonna assume this works out and finds the right things
		matchedStrings := re.FindAllString(scanner.Text(), -1)
		node := Node{ID: matchedStrings[0], Left: matchedStrings[1], Right: matchedStrings[2]}
		tree.Nodes[matchedStrings[0]] = node

	}

	return tree, nil
}

type Tree struct {
	Directions string
	// Map of Node ID to Node
	Nodes            map[string]Node
	FoundDestination bool
}

type Node struct {
	ID string
	// The ID of the Node to the left
	Left string
	// The ID of the Node to the right
	Right      string
	EndDepth int
}

func measureTime(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)
	fmt.Printf(" - %dµs\n", duration.Microseconds())
}
