package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	First  []int
	Second []int
}

func FindSmallest(list []int) (int, []int) {
	smallest := list[0]
	for _, item := range list {
		if item < smallest {
			smallest = item
		}
	}

	found := false
	results := []int{}
	for _, item := range list {
		if !found && item == smallest {
			found = true
		} else {
			results = append(results, item)
		}
	}

	return smallest, results
}

func parseInput(path string) (Input, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Input{}, err
	}
	first := []int{}
	second := []int{}
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		split := strings.Split(string(line), "   ")
		firstInt, err := strconv.Atoi(split[0])
		if err != nil {
			return Input{}, fmt.Errorf("failed on line %v, err %v", i, err)
		}
		secondInt, err := strconv.Atoi(split[1])
		if err != nil {
			return Input{}, fmt.Errorf("failed on line %v, err %v", i, err)
		}
		first = append(first, firstInt)
		second = append(second, secondInt)
	}
	return Input{
		First:  first,
		Second: second,
	}, nil
}

func iterateInput(input *Input) int {
	result := 0
	// log.Printf("first %v", input.First)
	// log.Printf("second %v", input.Second)

	for range len(input.First) {
		first, list := FindSmallest(input.First)
		input.First = list
		second, list := FindSmallest(input.Second)
		input.Second = list

		if first > second {
			result += first - second
		} else {
			result += second - first
		}
		// log.Printf("first %v second %v", first, second)
	}
	return result

}

func calculateSimilarityScore(input Input) int {
	result := 0
	secondCounts := make(map[int]int, 0)
	for _, second := range input.Second {
		_, exists := secondCounts[second]
		if exists {
			secondCounts[second] += second
		} else {
			secondCounts[second] = second
		}
	}
	for _, first := range input.First {
		if _, exists := secondCounts[first]; exists {
			result += secondCounts[first]
		}
	}
	return result
}

func main() {
	input, err := parseInput("sampleInput")
	if err != nil {
		log.Fatal(err)
	}
	// tests
	test, testList := FindSmallest(input.First)
	if test != 1 {
		log.Fatalf("wrong result returned from FindIndexSmallest, got %v want %v", test, 1)
	}
	if len(testList) != 5 {
		log.Fatalf("wrong length of returned list, got %v want %v", len(testList), 5)
	}
	testSimilarity := calculateSimilarityScore(input)
	if testSimilarity != 31 {
		log.Fatalf("wrong result returned, got %v want %v", testSimilarity, 31)
	}
	result := iterateInput(&input)
	if result != 11 {
		log.Fatalf("wrong result returned, got %v want %v", result, 11)
	}
	// log.Printf("Test: %v", result)

	// actual
	input, err = parseInput("input")
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("%v", input)
	similarity := calculateSimilarityScore(input)
	result = iterateInput(&input)
	log.Printf("Part 1: %v", result)
	log.Printf("Part 2: %v", similarity)
}
