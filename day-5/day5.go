package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Rules   map[int][]int
	Updates [][]int
}

func parseInput(path string) *Input {

	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	result := &Input{
		Rules:   make(map[int][]int, 0),
		Updates: [][]int{},
	}
	sections := strings.Split(string(content), "\n\n")

	// rules

	rules := strings.Split(sections[0], "\n")
	for _, rule := range rules {
		ints := strings.Split(rule, "|")
		key, _ := strconv.Atoi(ints[0])
		value, _ := strconv.Atoi(ints[1])
		_, exists := result.Rules[key]
		if exists {
			result.Rules[key] = append(result.Rules[key], value)
		} else {
			result.Rules[key] = []int{value}
		}
	}

	// updates

	updates := strings.Split(sections[1], "\n")
	for _, update := range updates {
		ints := strings.Split(update, ",")
		array := []int{}
		for _, number := range ints {
			value, _ := strconv.Atoi(number)
			array = append(array, value)
		}
		result.Updates = append(result.Updates, array)
	}

	return result
}

func UpdateOrderCorrect(update []int, rules map[int][]int) bool {
	for i, page := range update {
		if i == len(update)-1 {
			return true
		}
		for j, otherPage := range update {
			// don't check pages before
			if j <= i {
				continue
			}
			found := false
			for _, rule := range rules[page] {
				// log.Printf("check %v %v rul %v", page, otherPage, rule)
				if rule == otherPage {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func FixOrder(update []int, rules map[int][]int) []int {
	result := make([]int, len(update))
	copy(result, update)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(result)-1; i++ {
			page1 := result[i]
			page2 := result[i+1]

			if shouldSwap(page1, page2, rules) {
				result[i], result[i+1] = result[i+1], result[i]
				changed = true
			}
		}
	}

	return result
}

func shouldSwap(page1, page2 int, rules map[int][]int) bool {
	if deps, exists := rules[page2]; exists {
		for _, dep := range deps {
			if dep == page1 {
				return true
			}
		}
	}
	return false
}

func Middle(update []int) int {
	return update[len(update)/2]
}

func PartOneCountMiddlePagesCorrectOrder(input *Input) int {
	result := 0
	for _, update := range input.Updates {
		if UpdateOrderCorrect(update, input.Rules) {
			result += Middle(update)
		}
	}
	return result
}

func PartTwoCountMiddlePagesIncorrectOrder(input *Input) int {
	result := 0
	for _, update := range input.Updates {
		if !UpdateOrderCorrect(update, input.Rules) {
			corrected := FixOrder(update, input.Rules)
			result += Middle(corrected)
		}
	}
	return result

}

func main() {
	testInput := parseInput("sampleInput")
	testPartOne := PartOneCountMiddlePagesCorrectOrder(testInput)
	if testPartOne != 143 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 143)
	}
	testPartTwo := PartTwoCountMiddlePagesIncorrectOrder(testInput)
	if testPartTwo != 123 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 123)
	}
	input := parseInput("input")
	PartOne := PartOneCountMiddlePagesCorrectOrder(input)
	log.Printf("Part One: %v", PartOne)
	PartTwo := PartTwoCountMiddlePagesIncorrectOrder(input)
	log.Printf("Part Two: %v", PartTwo)

}
