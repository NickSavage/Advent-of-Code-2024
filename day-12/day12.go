package main

import (
	"log"
	"os"
	"strings"
)

type Position struct {
	Y int
	X int
}
type Plot struct {
	X         int
	Y         int
	Letter    string
	Perimeter int
	Visited   bool
}

type Input struct {
	Plots   [][]Plot
	Visited [][]bool
	MaxX    int
	MaxY    int
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	rows := strings.Split(string(content), "\n")
	result := Input{
		Plots: [][]Plot{},
	}
	for y, row := range rows {
		cols := strings.Split(row, "")
		new := []Plot{}
		for x, col := range cols {
			new = append(new, Plot{Y: y, X: x, Letter: col})
		}
		result.Plots = append(result.Plots, new)
	}
	rowCount := len(result.Plots)
	colCount := len(result.Plots[0])
	slice := make([][]bool, rowCount)
	for i := range slice {
		slice[i] = make([]bool, colCount)
	}
	result.Visited = slice
	result.MaxX = colCount
	result.MaxY = rowCount
	return &result
}

func GetPlotAtPosition(input *Input, pos Position) string {
	if pos.X < 0 || pos.Y < 0 {
		return ""
	}
	if pos.X > input.MaxX-1 || pos.Y > input.MaxY-1 {
		return ""
	}
	return input.Plots[pos.Y][pos.X].Letter
}

func FindPlots(input *Input, letter string, position Position) []Position {
	check := GetPlotAtPosition(input, position)
	if check == "" || check != letter {
		return []Position{}
	}
	if input.Plots[position.Y][position.X].Visited {
		return []Position{}
	}
	input.Plots[position.Y][position.X].Visited = true
	results := []Position{
		position,
	}

	up := FindPlots(input, letter, Position{Y: position.Y - 1, X: position.X})
	for _, p := range up {
		results = append(results, p)
	}
	down := FindPlots(input, letter, Position{Y: position.Y + 1, X: position.X})
	for _, p := range down {
		results = append(results, p)
	}
	left := FindPlots(input, letter, Position{Y: position.Y, X: position.X - 1})
	for _, p := range left {
		results = append(results, p)
	}
	right := FindPlots(input, letter, Position{Y: position.Y, X: position.X + 1})
	for _, p := range right {
		results = append(results, p)
	}
	return results
}

func CalculatePerimeter(input *Input, pos Position) int {
	letter := GetPlotAtPosition(input, pos)
	nextPos := pos
	perimeter := 0

	nextPos.X -= 1

	if GetPlotAtPosition(input, nextPos) != letter {
		perimeter += 1
	}
	nextPos.X += 1
	nextPos.Y -= 1
	if GetPlotAtPosition(input, nextPos) != letter {
		perimeter += 1
	}
	nextPos.Y += 1
	nextPos.X += 1

	if GetPlotAtPosition(input, nextPos) != letter {
		perimeter += 1
	}
	nextPos.X -= 1
	nextPos.Y += 1
	if GetPlotAtPosition(input, nextPos) != letter {
		perimeter += 1
	}

	return perimeter
}

func PartOneCostFences(input *Input) int {
	result := 0
	for y, rows := range input.Plots {
		for x, _ := range rows {
			plot := input.Plots[y][x]
			plot.Perimeter = CalculatePerimeter(input, Position{Y: y, X: x})
			input.Plots[y][x] = plot
		}
	}
	for y, rows := range input.Plots {
		for x, _ := range rows {
			if input.Plots[y][x].Visited {
				continue
			}
			region := FindPlots(input, input.Plots[y][x].Letter, Position{X: x, Y: y})
			area := 0
			perimeter := 0
			for _, plot := range region {
				area += 1
				perimeter += input.Plots[plot.Y][plot.X].Perimeter
			}
			result += area * perimeter
		}
	}
	return result
}

func main() {
	testInput := parseInput("sampleInput")
	testPartOne := PartOneCostFences(testInput)
	if testPartOne != 1930 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 1930)
	}

	input := parseInput("input")
	partOne := PartOneCostFences(input)
	log.Printf("Part One: %v", partOne)

	// input = parseInput("input")
	// partTwo := PartOneCostFences(input)
	// log.Printf("Part Two: %v", partTwo)
}
