package main

import (
	"log"
	"os"
	"strings"
)

type Input struct {
	Letters [][]string
	MaxX    int
	MaxY    int
}

func parseInput(path string) *Input {

	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	results := [][]string{}
	rows := strings.Split(string(content), "\n")
	for _, row := range rows {
		letters := strings.Split(row, "")
		results = append(results, letters)

	}
	return &Input{
		Letters: results,
		MaxX:    len(results[0]),
		MaxY:    len(results),
	}

}

func FindWords(input *Input, x, y int) int {
	results := 0
	// right
	if input.MaxX-x <= 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y][x+1] + input.Letters[y][x+2] + input.Letters[y][x+3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}

	// down right

	if input.MaxX-x <= 3 || input.MaxY-y <= 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y+1][x+1] + input.Letters[y+2][x+2] + input.Letters[y+3][x+3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	// down

	if input.MaxY-y <= 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y+1][x] + input.Letters[y+2][x] + input.Letters[y+3][x]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}

	// down left

	if x < 3 || input.MaxY-y <= 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y+1][x-1] + input.Letters[y+2][x-2] + input.Letters[y+3][x-3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	// left

	if x < 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y][x-1] + input.Letters[y][x-2] + input.Letters[y][x-3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	// up left

	if x < 3 || y < 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y-1][x-1] + input.Letters[y-2][x-2] + input.Letters[y-3][x-3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	// up

	if y < 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y-1][x] + input.Letters[y-2][x] + input.Letters[y-3][x]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	// up right

	if input.MaxX-x <= 3 || y < 3 {

	} else {
		word := input.Letters[y][x] + input.Letters[y-1][x+1] + input.Letters[y-2][x+2] + input.Letters[y-3][x+3]
		log.Printf("word %v, %v", word, word == "XMAS")
		if word == "XMAS" {
			results += 1
		}

	}
	return results
}

func FindWordCross(input *Input, x, y int) int {
	results := 0
	if x == 0 || y == 0 {
		return 0
	}
	if x == input.MaxX-1 || y == input.MaxY-1 {
		return 0
	}

	if input.Letters[y][x] != "A" {
		return 0
	}

	if input.Letters[y-1][x-1] == "S" && input.Letters[y+1][x+1] == "M" {
		if input.Letters[y-1][x+1] == "S" && input.Letters[y+1][x-1] == "M" {
			return 1
		}
		if input.Letters[y-1][x+1] == "M" && input.Letters[y+1][x-1] == "S" {
			return 1
		}
	}
	if input.Letters[y-1][x-1] == "M" && input.Letters[y+1][x+1] == "S" {
		if input.Letters[y-1][x+1] == "M" && input.Letters[y+1][x-1] == "S" {
			return 1
		}
		if input.Letters[y-1][x+1] == "S" && input.Letters[y+1][x-1] == "M" {
			return 1
		}
	}

	return results
}

func PartOneFindXmas(input *Input) int {
	result := 0
	for y, row := range input.Letters {
		for x, _ := range row {
			log.Printf("row %v", row)
			result += FindWords(input, x, y)
		}
	}
	return result
}

func PartTwoFindXmas(input *Input) int {
	result := 0
	for y, row := range input.Letters {
		for x, _ := range row {
			log.Printf("row %v", row)
			result += FindWordCross(input, x, y)
		}
	}
	return result

}

func main() {
	testInput := parseInput("sampleInput")
	testPartOne := PartOneFindXmas(testInput)
	if testPartOne != 18 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 18)
	}
	testPartTwo := PartTwoFindXmas(testInput)
	if testPartTwo != 9 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 9)
	}

	input := parseInput("input")
	partOne := PartOneFindXmas(input)
	log.Printf("Part One: %v", partOne)
	partTwo := PartTwoFindXmas(input)
	log.Printf("Part Two: %v", partTwo)

}
