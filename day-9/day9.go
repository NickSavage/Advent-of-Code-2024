package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	DiskMap []string
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	return &Input{
		DiskMap: strings.Split(string(content), ""),
	}
}

func PartOneMoveFilesComputeChecksum(input *Input) int {
	fileAlternate := true
	fileID := -1
	diskString := []string{}
	for _, char := range input.DiskMap {
		number, _ := strconv.Atoi(char)
		if fileAlternate {
			fileID += 1

			for range number {
				diskString = append(diskString, strconv.Itoa(fileID))
			}
			fileAlternate = false
		} else {
			for range number {
				diskString = append(diskString, ".")
			}
			fileAlternate = true
		}

	}
	for i := len(diskString) - 1; i >= 0; i-- {
		for n, char := range diskString {
			if i == n {
				break
			}
			if char == "." {
				diskString[n] = diskString[i]
				diskString[i] = "."
				break
			}
		}
	}
	result := 0
	for i, char := range diskString {
		operand, _ := strconv.Atoi(char)
		result += i * operand
	}
	return result

}
func PartTwoMoveFilesContinguousSpace(input *Input) int {
	fileAlternate := true
	fileID := -1
	diskString := []string{}

	// Create initial disk layout
	for _, char := range input.DiskMap {
		number, _ := strconv.Atoi(char)
		if fileAlternate {
			fileID += 1
			for range number {
				diskString = append(diskString, strconv.Itoa(fileID))
			}
			fileAlternate = false
		} else {
			for range number {
				diskString = append(diskString, ".")
			}
			fileAlternate = true
		}
	}

	// Process files in decreasing file ID order
	for currentFileID := fileID; currentFileID >= 0; currentFileID-- {
		// Find the file size and position
		fileSize := 0
		fileStart := -1
		for i := 0; i < len(diskString); i++ {
			if diskString[i] == strconv.Itoa(currentFileID) {
				if fileStart == -1 {
					fileStart = i
				}
				fileSize++
			}
		}

		if fileSize == 0 {
			continue
		}

		// Find leftmost suitable free space
		currentFreeSize := 0
		bestFreeStart := -1

		for i := 0; i < fileStart; i++ {
			if diskString[i] == "." {
				if currentFreeSize == 0 {
					bestFreeStart = i
				}
				currentFreeSize++
				if currentFreeSize == fileSize {
					// Move the file
					for j := 0; j < fileSize; j++ {
						diskString[bestFreeStart+j] = strconv.Itoa(currentFileID)
						diskString[fileStart+j] = "."
					}
					break
				}
			} else {
				currentFreeSize = 0
				bestFreeStart = -1
			}
		}
	}

	// Calculate checksum
	result := 0
	for i, char := range diskString {
		if char != "." {
			operand, _ := strconv.Atoi(char)
			result += i * operand
		}
	}
	return result
}

func main() {

	testInput := parseInput("sampleInput")
	// log.Printf("test %v", testInput)
	input := parseInput("input")

	testPartOne := PartOneMoveFilesComputeChecksum(testInput)
	if testPartOne != 1928 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 1928)
	}

	partOne := PartOneMoveFilesComputeChecksum(input)
	log.Printf("Part One: %v", partOne)

	testPartTwo := PartTwoMoveFilesContinguousSpace(testInput)
	if testPartTwo != 2858 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 2858)
	}
	partTwo := PartTwoMoveFilesContinguousSpace(input)
	log.Printf("Part Two: %v", partTwo)
}
