package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}
type Direction = int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func TurnRight(d Direction) Direction {
	return (d + 1) % 4
}

func TurnLeft(d Direction) Direction {
	return (d + 4 - 1) % 4
}

type Input struct {
	Map   map[Position]string
	MaxX  int
	MaxY  int
	Start Position
	End   Position
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	text := strings.Split(string(content), "\n\n")
	mapString := text[0]
	rows := strings.Split(mapString, "\n")
	input := &Input{
		Map: make(map[Position]string, 0),
	}
	for y, row := range rows {
		split := strings.Split(row, "")
		values := make([]string, 0)
		for x, value := range split {
			position := Position{X: x, Y: y}
			if value == "S" {
				input.Start = position
			}
			if value == "E" {
				input.End = position
			}
			input.Map[position] = value
			values = append(values, value)
		}
	}
	return input
}
func PrintMap(input *Input) {
	if input == nil || len(input.Map) == 0 {
		fmt.Println("Empty map")
		return
	}

	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for pos := range input.Map {
		minX = min(minX, pos.X)
		minY = min(minY, pos.Y)
		maxX = max(maxX, pos.X)
		maxY = max(maxY, pos.Y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pos := Position{X: x, Y: y}
			if val, exists := input.Map[pos]; exists {
				fmt.Print(val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Printf("\nMap Dimensions: %dx%d\n", maxX-minX+1, maxY-minY+1)
	fmt.Printf("Start Position: (%d,%d)\n", input.Start.X, input.Start.Y)
	fmt.Printf("End Position: (%d,%d)\n", input.End.X, input.End.Y)
}

func GetNextPosition(direction Direction, position Position) Position {
	next := position
	if direction == Up {
		next.Y--
	} else if direction == Right {
		next.X++
	} else if direction == Down {
		next.Y++
	} else if direction == Left {
		next.X--
	} else {
		log.Fatalf("next position is invalid at %v", next)
	}
	return next
}

// Create a struct to track state
type State struct {
	Position  Position
	Direction Direction
}

func FindPath(input *Input, direction Direction, score int, position Position, visited map[State]int) int {
	// Create current state
	state := State{Position: position, Direction: direction}

	// Base cases
	if score >= math.MaxInt32 {
		return math.MaxInt32 // Path too long
	}
	if input.Map[position] == "#" {
		return math.MaxInt32 // Hit a wall
	}
	if val, exists := visited[state]; exists && val <= score {
		return math.MaxInt32 // We've been here before with a better score
	}
	if input.Map[position] == "E" {
		return score
	}

	// Mark this state as visited
	visited[state] = score

	// Try all possible moves
	minScore := math.MaxInt32

	// Try going straight
	next := GetNextPosition(direction, position)
	straightScore := FindPath(input, direction, score+1, next, visited)
	minScore = min(minScore, straightScore)

	// Try turning right
	right := TurnRight(direction)
	nextRight := GetNextPosition(right, position)
	rightScore := FindPath(input, right, score+1001, nextRight, visited)
	minScore = min(minScore, rightScore)

	// Try turning left
	left := TurnLeft(direction)
	nextLeft := GetNextPosition(left, position)
	leftScore := FindPath(input, left, score+1001, nextLeft, visited)
	minScore = min(minScore, leftScore)

	return minScore
}

func PartOneFindPaths(input *Input) int {
	// PrintMap(input)
	visited := make(map[State]int)
	result := FindPath(input, Right, 0, input.Start, visited)
	log.Printf("result %v", len(visited))
	if result == math.MaxInt32 {
		return -1 // No path found
	}
	return result
}

func main() {
	testInput := parseInput("sampleInput")
	testPartOne := PartOneFindPaths(testInput)
	if testPartOne != 7036 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 7036)
	}
	testInput = parseInput("sampleInput2")
	testPartOne = PartOneFindPaths(testInput)
	if testPartOne != 11048 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 11048)
	}

	input := parseInput("input")
	partOne := PartOneFindPaths(input)
	log.Printf("Part One: %v", partOne)
	// testPartTwo := PartTwoCountTokens(testInput)
	// if testPartTwo != 480 {
	// log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 480)
	// }
}
