package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Reports [][]int
}

func parseInput(path string) Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return Input{}
	}

	result := Input{
		Reports: make([][]int, 0),
	}

	reports := strings.Split(string(content), "\n")
	for _, report := range reports {
		data := []int{}
		levels := strings.Split(string(report), " ")
		for _, level := range levels {
			converted, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("something has gone wrong with converting ints: %v", err)
			}
			data = append(data, converted)

		}
		result.Reports = append(result.Reports, data)
	}
	return result
}

type Direction int

const (
	Null       Direction = 0
	Increasing Direction = 1
	Decreasing Direction = 2
	Same       Direction = 3
)

func CheckIfReportIsSafe(report []int) bool {
	increasingOrDecreasing := Null
	last := 0

	for i, level := range report {
		if i == 0 {
			last = level
			if report[1]-report[0] > 0 {
				increasingOrDecreasing = Increasing
			} else if report[1]-report[0] < 0 {
				increasingOrDecreasing = Decreasing
			} else {
				increasingOrDecreasing = Same
			}
			continue
		}
		var direction Direction
		if level > last {
			direction = Increasing
		} else if last > level {
			direction = Decreasing
		} else {
			return false
		}
		if increasingOrDecreasing != Null && increasingOrDecreasing != direction {
			return false
		}
		difference := level - last
		if difference > 3 || difference < -3 {
			return false
		}
		last = level
	}
	return true
}

func PartOneSumSafeLevels(input Input) int {
	result := 0
	for _, report := range input.Reports {
		if CheckIfReportIsSafe(report) {
			result += 1
		} else {
		}
	}
	return result

}

func CheckIfReportIsSafeProblemDampener(report []int, dampenerActivated bool) bool {
	if CheckIfReportIsSafe(report) {
		return true
	}

	if dampenerActivated {
		return false
	}

	for i := 0; i < len(report); i++ {
		newSlice := make([]int, 0)
		newSlice = append(newSlice, report[:i]...)
		newSlice = append(newSlice, report[i+1:]...)

		if CheckIfReportIsSafe(newSlice) {
			return true
		}
	}

	return false
}
func PartTwoSumSafeLevels(input Input) int {
	result := 0
	for _, report := range input.Reports {
		if CheckIfReportIsSafeProblemDampener(report, false) {
			result += 1
		}
	}
	return result

}

func main() {
	testInput := parseInput("sampleInput")
	log.Printf("input %v", testInput)
	testResult := PartOneSumSafeLevels(testInput)
	log.Printf("test part one: %v", testResult)
	testResult = PartTwoSumSafeLevels(testInput)
	log.Printf("test part two: %v", testResult)

	input := parseInput("input")
	result := PartOneSumSafeLevels(input)
	log.Printf("part one: %v", result)
	result = PartTwoSumSafeLevels(input)
	log.Printf("part two: %v", result)
}
