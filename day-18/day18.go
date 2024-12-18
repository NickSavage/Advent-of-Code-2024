package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

type State struct {
	Position  Position
	Direction Direction
}

type Input struct {
	Map    [][]string
	MaxX   int
	MaxY   int
	Player Position
	Bytes  []Position

	MaxInstruction int
}

func parseInput(path string, x, y int) *Input {
	// Read file content
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	// Initialize result
	result := &Input{
		Map:   make([][]string, y),
		Bytes: make([]Position, 0),
	}

	for i := 0; i < y; i++ {
		result.Map[i] = make([]string, x)
		for j := 0; j < x; j++ {
			result.Map[i][j] = "."
		}
	}

	// Split into lines and parse each coordinate pair
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			log.Fatalf("invalid coordinate pair: %v", line)
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("invalid x coordinate: %v", coords[0])
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("invalid y coordinate: %v", coords[1])
		}

		result.Bytes = append(result.Bytes, Position{X: x, Y: y})
	}

	result.MaxX = x - 1
	result.MaxY = y - 1

	return result
}

// write a function that will print the map
func printMap(input *Input) {
	for _, row := range input.Map {
		fmt.Println(strings.Join(row, ""))
	}
}

func GetNextPosition(direction Direction, position Position, input *Input) Position {
	// we need to check for out of bounds here

	if direction == Up && position.Y == 0 {
		return position
	}
	if direction == Right && position.X == input.MaxX {
		return position
	}
	if direction == Down && position.Y == input.MaxY {
		return position
	}
	if direction == Left && position.X == 0 {
		return position
	}
	switch direction {
	case Up:
		return Position{X: position.X, Y: position.Y - 1}
	case Right:
		return Position{X: position.X + 1, Y: position.Y}
	case Down:
		return Position{X: position.X, Y: position.Y + 1}
	case Left:
		return Position{X: position.X - 1, Y: position.Y}
	}
	return position
}

func FindPath(input *Input, direction Direction, score int, visited map[State]int, position Position) int {
	state := State{Position: position, Direction: direction}
	// we ened to check for out of bounds here
	if position.X < 0 || position.X > input.MaxX || position.Y < 0 || position.Y > input.MaxY {
		return 0
	}

	current := input.Map[position.Y][position.X]
	if current == "#" {
		return 0
	}
	if val, exists := visited[state]; exists && val <= score {
		return 0
	}
	// found the end
	if position.X == input.MaxX && position.Y == input.MaxY {
		return score
	}
	visited[state] = score
	up := GetNextPosition(Up, position, input)
	upScore := FindPath(input, Up, score+1, visited, up)
	right := GetNextPosition(Right, position, input)
	rightScore := FindPath(input, Right, score+1, visited, right)
	down := GetNextPosition(Down, position, input)
	downScore := FindPath(input, Down, score+1, visited, down)
	left := GetNextPosition(Left, position, input)
	leftScore := FindPath(input, Left, score+1, visited, left)

	scores := []int{upScore, rightScore, downScore, leftScore}
	validScores := []int{}
	for _, s := range scores {
		if s > 0 {
			validScores = append(validScores, s)
		}
	}
	if len(validScores) == 0 {
		return 0
	}
	minScore := validScores[0]
	for _, s := range validScores {
		if s < minScore {
			minScore = s
		}
	}
	return minScore
}

func PartOneFindRoute(input *Input, steps int) int {
	for i := range steps {
		instruction := input.Bytes[i]
		input.Map[instruction.Y][instruction.X] = "#"
		input.MaxInstruction = i
	}
	visited := make(map[State]int)
	result := FindPath(input, Up, 0, visited, Position{X: 0, Y: 0})

	return result
}

func PartTwoFindRoute(input *Input) string {
	for i := range len(input.Bytes) {
		log.Printf("i: %v", i)
		instruction := input.Bytes[i]
		input.Map[instruction.Y][instruction.X] = "#"
		input.MaxInstruction = i
		visited := make(map[State]int)
		result := FindPath(input, Up, 0, visited, Position{X: 0, Y: 0})
		if result == 0 {
			return fmt.Sprintf("%v,%v", instruction.X, instruction.Y)
		}
	}
	return ""
}

func main() {
	sampleInput := parseInput("sampleInput", 7, 7)
	printMap(sampleInput)
	testPartOne := PartOneFindRoute(sampleInput, 12)
	if testPartOne != 22 {
		log.Fatalf("got wrong answer for part one, got %v want %v", testPartOne, 22)
	}
	input := parseInput("input", 71, 71)
	partOne := PartOneFindRoute(input, 1024)
	log.Printf("Part One: %v", partOne)

	sampleInput = parseInput("sampleInput", 7, 7)
	testPartTwo := PartTwoFindRoute(sampleInput)
	if testPartTwo != "6,1" {
		log.Fatalf("got wrong answer for part two, got %v want %v", testPartTwo, "6,1")
	}
	input = parseInput("input", 71, 71)
	partTwo := PartTwoFindRoute(input)
	log.Printf("Part Two: %v", partTwo)
}
