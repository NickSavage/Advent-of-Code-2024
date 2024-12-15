package main

import (
	"log"
	"os"
	"strings"
)

type Input struct {
	Map          [][]string
	Instructions []string
	RobotX       int
	RobotY       int
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	text := strings.Split(string(content), "\n\n")
	mapString := text[0]
	rows := strings.Split(mapString, "\n")
	input := &Input{
		Map:          make([][]string, 0),
		Instructions: []string{},
	}
	for y, row := range rows {
		split := strings.Split(row, "")
		values := make([]string, 0)
		for x, value := range split {
			if value == "@" {
				input.RobotX = x
				input.RobotY = y
			}
			values = append(values, value)
		}
		input.Map = append(input.Map, values)
	}
	instructionsString := strings.Split(text[1], "")
	for _, instruction := range instructionsString {
		if instruction == "\n" {
			continue
		}
		input.Instructions = append(input.Instructions, instruction)
	}
	return input
}

func IsBox(text string) bool {
	if text == "O" {
		return true
	}
	if text == "[" || text == "]" {
		return true
	}
	return false
}

// returns true if the x,y is empty after moving
func AttemptMoveBox(input *Input, x, y int, direction string) bool {
	text := input.Map[y][x]
	isBox := IsBox(text)
	if !isBox {
		return true
	}
	// log.Printf("attempting to move boxwes at %v/%v", x, y)
	newX := x
	newY := y
	if direction == "left" {
		newX--
	}
	if direction == "right" {
		newX++
	}
	if direction == "up" {
		newY--
	}
	if direction == "down" {
		newY++
	}

	if IsBox(input.Map[newY][newX]) {
		AttemptMoveBox(input, newX, newY, direction)
	}
	if text == "O" {
		if input.Map[newY][newX] == "." {
			input.Map[newY][newX] = "O"
			input.Map[y][x] = "."
			return true
		}
	} else if text == "]" {
		if direction == "left" {
			if input.Map[newY][newX-1] == "." {
				input.Map[newY][newX-1] = "["
				input.Map[newY][newX] = "]"
				return true
			}
		}
		if direction == "right" {
			PrintBoard(input)
			log.Fatalf("pushed on a place you shouldn't have %v/%v", newX, newY)
		}
		if direction == "up" {
			if input.Map[newY][newX] == "." {
				if input.Map[newY][newX-1] == "." {
					input.Map[newY][newX-1] = "["
					input.Map[newY][newX] = "]"
					return true
				}
			}
		}
		if direction == "down" {
			if input.Map[newY][newX] == "." {
				if input.Map[newY][newX] == "." {
					input.Map[newY][newX-1] = "["
					input.Map[newY][newX-1] = "]"
					return true
				}
			}
		}

	} else if text == "[" {
		if direction == "left" {
			PrintBoard(input)
			log.Fatalf("pushed on a place you shouldn't have %v/%v", newX, newY)
		}
		if direction == "right" {
		}
		if direction == "up" {
			if input.Map[newY][newX] == "." {
				if input.Map[newY][newX+1] == "." {
					input.Map[newY][newX+1] = "]"
					input.Map[newY][newX] = "["
					return true
				}
			}
		}
		if direction == "down" {
			if input.Map[newY][newX] == "." {
				if input.Map[newY][newX+1] == "." {
					input.Map[newY][newX+1] = "]"
					input.Map[newY][newX+1] = "["
					return true
				}
			}
		}

	}
	return false

}

func PrintBoard(input *Input) {
	for _, row := range input.Map {
		for _, value := range row {
			print(value)
		}
		print("\n")
	}

}

func CountBoard(input *Input) int {
	result := 0
	for x, row := range input.Map {
		for y, value := range row {
			if value == "O" {
				result += 100*x + y
			}
		}
	}
	return result
}

func PartOneMoveRobots(input *Input) int {
	// log.Printf("map %v", input.Map)
	// log.Printf("instructions %v", input.Instructions)
	// PrintBoard(input)
	for _, instruction := range input.Instructions {
		// log.Printf("instruction %v - %v", i, instruction)
		// log.Printf("robot %v/%v", input.RobotX, input.RobotY)

		newX := input.RobotX
		newY := input.RobotY
		direction := ""
		if instruction == "<" {
			newX--
			direction = "left"
		}
		if instruction == ">" {
			newX++
			direction = "right"
		}
		if instruction == "^" {
			newY--
			direction = "up"
		}
		if instruction == "v" {
			newY++
			direction = "down"
		}
		if direction == "" {
			log.Fatalf("direction is nil at %v/%v", newX, newY)
		}
		AttemptMoveBox(input, newX, newY, direction)
		if input.Map[newY][newX] == "." {
			// log.Printf("moving robot to %v/%v", newX, newY)
			input.Map[newY][newX] = "@"
			input.Map[input.RobotY][input.RobotX] = "."
			input.RobotY = newY
			input.RobotX = newX
		}
		// PrintBoard(input)
	}
	result := CountBoard(input)
	return result
}

func main() {
	testInput := parseInput("sampleInput2")
	testPartOne := PartOneMoveRobots(testInput)
	if testPartOne != 2028 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 2028)
	}
	testInput = parseInput("sampleInput")
	testPartOne = PartOneMoveRobots(testInput)
	if testPartOne != 10092 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 10092)
	}
	input := parseInput("input")
	partOne := PartOneMoveRobots(input)
	log.Printf("Part One: %v", partOne)
	// testPartTwo := PartTwoCountTokens(testInput)
	// if testPartTwo != 480 {
	// log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 480)
	// }
}
