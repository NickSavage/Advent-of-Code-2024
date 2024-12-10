package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Input struct {
	Trailheads []Position
	Map        map[Position]int
	MaxX       int
	MaxY       int
	PeaksFound map[Position]bool
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	rows := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(rows) == 0 {
		return &Input{}
	}

	trailheads := []Position{}
	topMap := make(map[Position]int)
	maxY := len(rows)
	maxX := len(strings.Split(rows[0], ""))

	for y, row := range rows {
		cols := strings.Split(row, "")
		for x, height := range cols {
			heightInt, _ := strconv.Atoi(height)
			pos := Position{
				X: x, // Using x for X coordinate
				Y: y, // Using y for Y coordinate
			}
			if heightInt == 0 {
				trailheads = append(trailheads, pos)
			}
			topMap[pos] = heightInt
		}
	}

	return &Input{
		Map:        topMap,
		Trailheads: trailheads,
		MaxX:       maxX,
		MaxY:       maxY,
	}
}

func CheckBounds(input *Input, pos Position) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.X < input.MaxX && pos.Y < input.MaxY
}
func FindPaths(input *Input, pos Position, currentHeight int) int {
	if !CheckBounds(input, pos) {
		return 0
	}

	if input.Map[pos] != currentHeight {
		return 0
	}

	if currentHeight == 9 {
		if !input.PeaksFound[pos] {
			log.Printf("scoring %v %v", pos, input.Map[pos])
			input.PeaksFound[pos] = true
			return 1
		}
	}

	result := 0
	directions := []Position{
		{X: pos.X + 1, Y: pos.Y}, // right
		{X: pos.X, Y: pos.Y + 1}, // down
		{X: pos.X - 1, Y: pos.Y}, // left
		{X: pos.X, Y: pos.Y - 1}, // up
	}

	for _, nextPos := range directions {
		if CheckBounds(input, nextPos) && input.Map[nextPos] == currentHeight+1 {
			result += FindPaths(input, nextPos, currentHeight+1)
		}
	}

	return result
}
func PartOneScoreTrailheads(input *Input) int {
	score := 0
	for _, trailhead := range input.Trailheads {
		input.PeaksFound = make(map[Position]bool, 0)
		paths := FindPaths(input, trailhead, 0)
		score += paths
	}
	return score
}
func FindPathsPartTwo(input *Input, pos Position, currentHeight int, visited map[Position]bool) int {
	if !CheckBounds(input, pos) || visited[pos] {
		return 0
	}

	if input.Map[pos] != currentHeight {
		return 0
	}

	if currentHeight == 9 {
		return 1
	}

	visited[pos] = true
	defer func() { visited[pos] = false }()

	result := 0
	directions := []Position{
		{X: pos.X + 1, Y: pos.Y}, // right
		{X: pos.X, Y: pos.Y + 1}, // down
		{X: pos.X - 1, Y: pos.Y}, // left
		{X: pos.X, Y: pos.Y - 1}, // up
	}

	for _, nextPos := range directions {
		if CheckBounds(input, nextPos) && !visited[nextPos] && input.Map[nextPos] == currentHeight+1 {
			result += FindPathsPartTwo(input, nextPos, currentHeight+1, visited)
		}
	}

	return result
}
func PartTwoScoreTrailheads(input *Input) int {
	score := 0
	for i, trailhead := range input.Trailheads {
		visited := make(map[Position]bool)
		paths := FindPathsPartTwo(input, trailhead, 0, visited)
		log.Printf("Trailhead %d at position {%d,%d} found %d paths", i, trailhead.X, trailhead.Y, paths)
		score += paths
	}
	return score
}

func main() {
	testInput := parseInput("sampleInput")
	log.Printf("Found %d trailheads in test input", len(testInput.Trailheads))
	testPartOne := PartOneScoreTrailheads(testInput)
	if testPartOne != 36 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 36)
	}

	input := parseInput("input")
	partOne := PartOneScoreTrailheads(input)
	log.Printf("Part One: %v", partOne)

	testPartTwo := PartTwoScoreTrailheads(testInput)
	if testPartTwo != 81 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 81)
	}

	partTwo := PartTwoScoreTrailheads(input)
	log.Printf("Part Two: %v", partTwo)
}
