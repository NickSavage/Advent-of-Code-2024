package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Stones []int
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	stones := strings.Split(strings.TrimSpace(string(content)), " ")
	result := Input{
		Stones: []int{},
	}
	for _, stone := range stones {
		conv, err := strconv.Atoi(stone)
		if err != nil {
			log.Fatalf("failed to convert %v", stone)
		}
		result.Stones = append(result.Stones, conv)
	}
	return &result
}

func CountStones(stone int, iterations int) int {
	new := IterateStone(stone)

	if iterations == 1 {
		return len(new)
	}
	result := 0
	for _, stone := range new {
		result += CountStones(stone, iterations-1)
	}

	return result
}

func IterateStone(stone int) []int {
	new := []int{}
	stoneStr := strconv.Itoa(stone) // Convert int to string
	if stone == 0 {
		new = append(new, 1)
	} else if len(stoneStr)%2 == 0 {
		mid := len(stoneStr) / 2
		firstHalf, _ := strconv.Atoi(stoneStr[:mid])
		secondHalf, _ := strconv.Atoi(stoneStr[mid:])
		new = append(new, firstHalf)
		new = append(new, secondHalf)
	} else {
		new = append(new, stone*2024)

	}
	return new

}

func PartOneCountStones(input *Input, iterations int) int {
	results := 0
	for _, original := range input.Stones {
		counts := make(map[int]int)
		counts[original] = 1

		// Iterate for the specified number of times
		for i := 0; i < iterations; i++ {
			newCounts := make(map[int]int)

			// Process each stone in the current counts
			for stone, count := range counts {
				if count > 0 {
					// Get new stones from iterating current stone
					new := IterateStone(stone)
					// Add new stones to the new counts map
					for _, newStone := range new {
						newCounts[newStone] += count
					}
				}
			}

			// Replace old counts with new counts
			counts = newCounts
		}

		// Sum up all stones in the final counts
		for _, count := range counts {
			results += count
		}
	}
	return results
}

func main() {
	testInput := parseInput("sampleInput")
	testPartOne := PartOneCountStones(testInput, 6)
	if testPartOne != 22 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 22)
	}
	testInput = parseInput("sampleInput")
	testPartOne = PartOneCountStones(testInput, 25)
	if testPartOne != 55312 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 55312)
	}

	input := parseInput("input")
	partOne := PartOneCountStones(input, 25)
	log.Printf("Part One: %v", partOne)

	input = parseInput("input")
	partTwo := PartOneCountStones(input, 75)
	log.Printf("Part Two: %v", partTwo)
}
