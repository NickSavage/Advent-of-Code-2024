package main

import (
	"log"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

type State struct {
	Position  Position
	Direction Direction
}

type Input struct {
	Map               [][]string
	Position          Position
	Direction         Direction
	DistinctPositions int
	MaxX              int
	MaxY              int
	OffScreen         bool
}

type Direction = int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func parseInput(path string) *Input {

	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	rows := strings.Split(string(content), "\n")
	mapRows := [][]string{}
	for _, row := range rows {
		new := strings.Split(row, "")
		mapRows = append(mapRows, new)
	}
	result := &Input{
		Map:               mapRows,
		DistinctPositions: 1,
	}
	return result
}

func FindPosition(positions [][]string) (Position, Direction) {
	for y, row := range positions {
		for x, value := range row {
			if value == "v" {
				return Position{X: x, Y: y}, Down
			}
			if value == "^" {
				return Position{X: x, Y: y}, Up
			}
			if value == "<" {
				return Position{X: x, Y: y}, Left
			}
			if value == ">" {
				return Position{X: x, Y: y}, Right
			}
		}
	}
	return Position{X: -1, Y: -1}, Up
}

func MoveGuard(input *Input) {
	pos := input.Position
	newPos := Position{X: pos.X, Y: pos.Y}
	if input.Direction == Up {
		newPos.Y -= 1
	}
	if input.Direction == Down {
		newPos.Y += 1
	}
	if input.Direction == Left {
		newPos.X -= 1
	}
	if input.Direction == Right {
		newPos.X += 1
	}
	if newPos.X < 0 || newPos.Y < 0 {
		input.OffScreen = true
		// input.DistinctPositions += 1
		return
	}
	if newPos.X == input.MaxX || newPos.Y == input.MaxY {
		input.OffScreen = true
		// input.DistinctPositions += 1
		return
	}
	if input.Map[newPos.Y][newPos.X] == "#" {
		if input.Direction == Left {
			input.Direction = 0
		} else {
			input.Direction += 1
		}
	} else if input.Map[newPos.Y][newPos.X] == "X" || input.Map[newPos.Y][newPos.X] == "^" {
		input.Position = newPos

	} else {
		input.DistinctPositions += 1
		input.Position = newPos
		input.Map[newPos.Y][newPos.X] = "X"
	}

}

func PartOneCountDistinctPositions(input *Input) int {
	input.MaxX = len(input.Map[0])
	input.MaxY = len(input.Map)

	guardPos, guardDir := FindPosition(input.Map)
	log.Printf("guard %v, %v", guardPos, guardDir)
	input.Position = guardPos
	input.Direction = guardDir
	// input.Map[guardPos.Y][guardPos.X] = "X"

	for {
		MoveGuard(input)
		log.Printf("new position %v, %v, %v", input.Position, input.Direction, input.DistinctPositions)

		if input.OffScreen {
			break
		}
	}

	return input.DistinctPositions

}
func copyMap(original [][]string) [][]string {
	newMap := make([][]string, len(original))
	for i := range original {
		newMap[i] = make([]string, len(original[i]))
		copy(newMap[i], original[i])
	}
	return newMap
}
func PartTwoCheckPlacementObstacle(input *Input) int {
	placements := 0
	originalMap := copyMap(input.Map)
	guardPos, guardDir := FindPosition(originalMap)

	// Skip the guard's starting position
	originalMap[guardPos.Y][guardPos.X] = "."

	for y := 0; y < len(originalMap); y++ {
		for x := 0; x < len(originalMap[0]); x++ {
			if originalMap[y][x] == "#" {
				continue
			}

			newMap := copyMap(originalMap)
			newMap[y][x] = "#"

			testInput := &Input{
				Map:       newMap,
				Position:  guardPos,
				Direction: guardDir,
				MaxX:      len(newMap[0]),
				MaxY:      len(newMap),
				OffScreen: false,
			}

			visited := make(map[State]bool)
			hasLoop := false

			for !testInput.OffScreen {
				currentState := State{
					Position:  testInput.Position,
					Direction: testInput.Direction,
				}

				if visited[currentState] {
					hasLoop = true
					break
				}

				visited[currentState] = true
				MoveGuard(testInput)
			}

			if hasLoop {
				placements++
			}
		}
	}
	return placements
}
func main() {
	testInput := parseInput("sampleInput")
	input := parseInput("input")

	testPartOne := PartOneCountDistinctPositions(testInput)
	if testPartOne != 41 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 41)
	}
	testPartTwo := PartTwoCheckPlacementObstacle(testInput)
	if testPartTwo != 6 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 6)
	}

	partOne := PartOneCountDistinctPositions(input)
	log.Printf("Part One: %v", partOne)
	partTwo := PartTwoCheckPlacementObstacle(input)
	log.Printf("Part Two: %v", partTwo)
}
