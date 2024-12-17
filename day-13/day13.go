package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Combination struct {
	A int
	B int
}

type Movement struct {
	Tokens int
	X      int
	Y      int
}
type Machine struct {
	A Movement
	B Movement
	X int
	Y int
}
type Input struct {
	Machines []Machine
}

func ParseMachineData(input string) (Machine, error) {
	var machine Machine
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "Button A:") {
			movement, err := parseMovement(strings.TrimPrefix(line, "Button A:"))
			if err != nil {
				return Machine{}, err
			}
			movement.Tokens = 3
			machine.A = movement
		} else if strings.HasPrefix(line, "Button B:") {
			movement, err := parseMovement(strings.TrimPrefix(line, "Button B:"))
			if err != nil {
				return Machine{}, err
			}
			movement.Tokens = 1
			machine.B = movement
		} else if strings.HasPrefix(line, "Prize:") {
			err := parsePrize(strings.TrimPrefix(line, "Prize:"), &machine)
			if err != nil {
				return Machine{}, err
			}
		}
	}

	return machine, nil
}

func parseMovement(s string) (Movement, error) {
	var movement Movement
	parts := strings.Split(strings.TrimSpace(s), ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "X+") {
			val, err := strconv.Atoi(strings.TrimPrefix(part, "X+"))
			if err != nil {
				return Movement{}, err
			}
			movement.X = val
		} else if strings.HasPrefix(part, "Y+") {
			val, err := strconv.Atoi(strings.TrimPrefix(part, "Y+"))
			if err != nil {
				return Movement{}, err
			}
			movement.Y = val
		}
	}

	return movement, nil
}

func parsePrize(s string, machine *Machine) error {
	parts := strings.Split(strings.TrimSpace(s), ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "X=") {
			val, err := strconv.Atoi(strings.TrimPrefix(part, "X="))
			if err != nil {
				return err
			}
			machine.X = val
		} else if strings.HasPrefix(part, "Y=") {
			val, err := strconv.Atoi(strings.TrimPrefix(part, "Y="))
			if err != nil {
				return err
			}
			machine.Y = val
		}
	}

	return nil
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	machines := strings.Split(string(content), "\n\n")
	result := Input{
		Machines: []Machine{},
	}
	for _, rawMachine := range machines {
		machine, err := ParseMachineData(rawMachine)
		if err != nil {
			log.Fatal(err)
		}

		result.Machines = append(result.Machines, machine)

	}
	return &result
}

func Determinant(ax, ay, bx, by int) int {
	return ax*by - ay*bx
}

func FindSolution(machine Machine) int {
	det := Determinant(machine.A.X, machine.A.Y, machine.B.X, machine.B.Y)
	log.Printf("det %v", det)
	if det == 0 {
		return 0
	}
	nx := (machine.X * machine.B.Y) - (machine.Y * machine.B.X)
	log.Printf("x %v * by %v - y %v * ay %v", machine.X, machine.B.Y, machine.Y, machine.A.Y)
	x := nx / det
	
	ny := (machine.Y * machine.A.X) - (machine.X * machine.A.Y)
	y := ny / det
	log.Printf("nx %v, ny  %v", nx, ny)
	
	// Verify solution
	targetX := x*machine.A.X + y*machine.B.X
	targetY := x*machine.A.Y + y*machine.B.Y
	
	if targetX == machine.X && targetY == machine.Y && x >= 0 && y >= 0 {
		score := x*3 + y*1
		log.Printf("x %v y %v score %v", x, y, score)
		return score
	} else {
		log.Printf("result not good")
		return 0
	}
}

func PartOneCountTokens(input *Input) int {
	result := 0

	for _, machine := range input.Machines {
		result += FindSolution(machine)
	}
	return result

}
func PartTwoCountTokens(input *Input) int {
	result := 0
	log.Printf("part 2")
	for i, machine := range input.Machines {
		log.Printf("\nMachine %v", i)
		machine.X += 10000000000000
		machine.Y += 10000000000000

		result += FindSolution(machine)

	}
	return result
}

func main() {
	testInput := parseInput("sampleInput")
	log.Printf("%v", testInput)
	testPartOne := PartOneCountTokens(testInput)
	if testPartOne != 480 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 480)
	}

	input := parseInput("input")
	partOne := PartOneCountTokens(input)
	log.Printf("Part One: %v", partOne)

	testPartTwo := PartTwoCountTokens(testInput)
	if testPartTwo !=875318608908 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 875318608908)
	}
	input = parseInput("input")
	partTwo := PartTwoCountTokens(input)
	log.Printf("Part Two: %v", partTwo)
}
