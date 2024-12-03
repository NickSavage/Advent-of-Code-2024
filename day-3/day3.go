package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Input struct {
	Memory string
	Pairs  [][]int
	Result int
	Active bool
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	return &Input{
		Memory: string(content),
		Pairs:  make([][]int, 0),
		Active: true,
	}
}

func PartOneParseMemory(input *Input) int {

	input.Result = 0
	pattern := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)

	matches := pattern.FindAllString(input.Memory, -1)
	for _, match := range matches {
		parsed := strings.ReplaceAll(match, "mul(", "")
		parsed = strings.ReplaceAll(parsed, ")", "")
		numbers := strings.Split(parsed, ",")

		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])
		result := []int{
			first,
			second,
		}
		input.Pairs = append(input.Pairs, result)
		input.Result += first * second
	}
	return input.Result

}

func PartTwoParseMemory(input *Input) int {
	input.Result = 0

	pattern := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)|do\(\)|don't\(\)`)

	matches := pattern.FindAllString(input.Memory, -1)
	for _, match := range matches {
		log.Printf("match %v", match)
		if strings.Contains(match, "don't()") {
			input.Active = false
			continue
		} else if strings.Contains(match, "do()") {
			input.Active = true
			continue
		}
		if !input.Active {
			continue
		}

		parsed := strings.ReplaceAll(match, "mul(", "")
		parsed = strings.ReplaceAll(parsed, ")", "")
		numbers := strings.Split(parsed, ",")

		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])
		result := []int{
			first,
			second,
		}
		log.Printf("maths %v", result)
		input.Pairs = append(input.Pairs, result)
		input.Result += first * second
	}
	return input.Result
}

func main() {
	testInput := parseInput("sampleInput")
	partOneTest := PartOneParseMemory(testInput)
	if partOneTest != 161 {
		log.Fatalf("got wwrong part one test got %v want %v", partOneTest, 161)
	}
	log.Printf("Part One Test: %v", partOneTest)

	partTwoTest := PartTwoParseMemory(testInput)
	if partTwoTest != 48 {
		log.Fatalf("got wwrong part one test got %v want %v", partTwoTest, 48)
	}
	log.Printf("Part Two Test: %v", partTwoTest)

	input := parseInput("input")
	partOne := PartOneParseMemory(input)
	log.Printf("Part One: %v", partOne)
	partTwo := PartTwoParseMemory(input)
	log.Printf("Part Two: %v", partTwo)
}
