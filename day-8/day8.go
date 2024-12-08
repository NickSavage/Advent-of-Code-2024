package main

import (
	"log"
	"os"
	"strings"
)

type Location struct {
	Y    int
	X    int
	Char string
}

type Input struct {
	Map               []Location
	Antenna           map[string][]Location
	Antinodes         []Location
	ResonantHarmonics bool

	MaxX int
	MaxY int
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}

	rows := strings.Split(string(content), "\n")
	locations := []Location{}
	antenna := make(map[string][]Location)

	for y, row := range rows {
		chars := strings.Split(row, "")
		for x, char := range chars {
			location := Location{
				Y:    y,
				X:    x,
				Char: char,
			}
			locations = append(locations, location)
			if char == "." {
				continue
			}
			if _, exists := antenna[char]; !exists {
				antenna[char] = []Location{}
			}
			antenna[char] = append(antenna[char], location)
		}

	}

	input := &Input{
		Map:       locations,
		Antenna:   antenna,
		MaxX:      len(rows),
		MaxY:      len(strings.Split(rows[0], "")),
		Antinodes: []Location{},
	}
	return input
}

func AntinodeExists(new Location, antinodes []Location) bool {
	for _, antinode := range antinodes {
		if antinode.X == new.X && antinode.Y == new.Y {
			return true
		}
	}
	return false
}

func AntinodeInBounds(antinode Location, input *Input) bool {
	if antinode.X < 0 || antinode.Y < 0 {
		return false
	}
	if antinode.X >= input.MaxX || antinode.Y >= input.MaxY {
		return false
	}
	return true
}

func CreateAntinodes(input *Input, one Location, two Location) {
	if one.X == two.X && one.Y == two.Y {
		return
	}

	if input.ResonantHarmonics {
		// Add the antennas themselves as antinodes
		if !AntinodeExists(one, input.Antinodes) {
			input.Antinodes = append(input.Antinodes, one)
		}
		if !AntinodeExists(two, input.Antinodes) {
			input.Antinodes = append(input.Antinodes, two)
		}

		diffX := two.X - one.X
		diffY := two.Y - one.Y

		for i := -input.MaxX; i <= input.MaxX; i++ {
			point := Location{
				X:    one.X + (i * diffX),
				Y:    one.Y + (i * diffY),
				Char: "#",
			}

			if AntinodeInBounds(point, input) && !AntinodeExists(point, input.Antinodes) {
				input.Antinodes = append(input.Antinodes, point)
			}
		}
	} else {
		diffX := two.X - one.X
		diffY := two.Y - one.Y

		three := Location{
			X:    one.X - diffX,
			Y:    one.Y - diffY,
			Char: "#",
		}
		four := Location{
			X:    two.X + diffX,
			Y:    two.Y + diffY,
			Char: "#",
		}

		if AntinodeInBounds(three, input) && !AntinodeExists(three, input.Antinodes) {
			input.Antinodes = append(input.Antinodes, three)
		}
		if AntinodeInBounds(four, input) && !AntinodeExists(four, input.Antinodes) {
			input.Antinodes = append(input.Antinodes, four)
		}
	}
}

func PartOneComputeUniqueAntinodes(input *Input) int {
	for freq, antennas := range input.Antenna {
		log.Printf("Processing frequency: %s", freq)
		for i, antenna1 := range antennas {
			for j, antenna2 := range antennas {
				if i >= j { // Skip same antenna and avoid double-counting pairs
					continue
				}
				CreateAntinodes(input, antenna1, antenna2)

				log.Printf("\nChecking pair: %v and %v", antenna1, antenna2)
			}
		}
	}

	// visualizeAntinodes(input)
	return len(input.Antinodes)
}

func PartTwoComputeAntinodes(input *Input) int {
	input.ResonantHarmonics = true
	return PartOneComputeUniqueAntinodes(input)

}
func visualizeAntinodes(input *Input) {
	// Create a 2D slice to represent the map
	grid := make([][]string, input.MaxY) // Note: switched to MaxY first
	for i := range grid {
		grid[i] = make([]string, input.MaxX) // Note: MaxX second
		// Fill with dots initially
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Mark all antinodes with #
	for _, antinode := range input.Antinodes {
		// Note: Using Y for rows and X for columns
		grid[antinode.Y][antinode.X] = "#"
	}

	// Print the original antennas as well, for reference
	for _, antennaList := range input.Antenna {
		for _, antenna := range antennaList {
			grid[antenna.Y][antenna.X] = antenna.Char
		}
	}

	// Print the map
	log.Println("Antinode Map:")
	for i := 0; i < input.MaxY; i++ {
		row := strings.Join(grid[i], "")
		log.Println(row)
	}
	log.Printf("Total antinodes: %d\n", len(input.Antinodes))
}
func main() {
	testInput := parseInput("sampleInput")
	// log.Printf("test %v", testInput)
	input := parseInput("input")

	testPartOne := PartOneComputeUniqueAntinodes(testInput)
	visualizeAntinodes(testInput)
	if testPartOne != 14 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 14)
	}
	partOne := PartOneComputeUniqueAntinodes(input)
	log.Printf("Part One: %v", partOne)
	testInput = parseInput("sampleInput")
	testPartTwo := PartTwoComputeAntinodes(testInput)
	visualizeAntinodes(testInput) // Add this line
	if testPartTwo != 34 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 34)
	}

	partTwo := PartTwoComputeAntinodes(input)
	log.Printf("Part Two: %v", partTwo)
}
