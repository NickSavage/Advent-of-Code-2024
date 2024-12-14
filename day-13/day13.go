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

func FindCombinations(machine Machine) []Combination {
	results := []Combination{}
	maxA := machine.X/machine.A.X + 1
	maxB := machine.X/machine.B.X + 1
	max := 0
	if maxA > maxB {
		max = maxA
	} else {
		max = maxB
	}

	for tryA := range max {
		for tryB := range max {
			check := tryA*machine.A.X + tryB*machine.B.X
			// log.Printf("a %v * %v + b %v * %v = %v", tryA, machine.A.X, tryB, machine.B.X, check)

			if check > machine.X {
				break
			}
			if check == machine.X {
				checkY := tryA*machine.A.Y + tryB*machine.B.Y
				if checkY == machine.Y {
					log.Printf("found!")

					results = append(results, Combination{
						A: tryA,
						B: tryB,
					})

				}
			}

		}
	}

	return results
}

func PartOneCountTokens(input *Input) int {
	result := 0

	for i, machine := range input.Machines {
		log.Printf("i %v", i)
		combos := FindCombinations(machine)
		min := -1
		// found := []Combination{}
		for _, combo := range combos {
			score := combo.A*3 + combo.B*1
			log.Printf("combo %v score %v", combo, score)
			if min == -1 {
				min = score
			} else if score < min {
				min = score
			}
		}
		if min == -1 {
			min = 0
		}
		result += min

	}
	return result

}
func GCD(a, b int) int {
	// Make sure we're working with absolute values
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// GCD(a,b) = GCD(b,a mod b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Returns gcd and the coefficients x, y where ax + by = gcd
func ExtendedGCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := ExtendedGCD(b, a%b)

	x = y1
	y = x1 - (a/b)*y1

	return gcd, x, y
}
func PartTwoCountTokens(input *Input) int {
	result := 0
	log.Printf("part 2")
	for i, machine := range input.Machines {
		log.Printf("\nMachine %v", i)
		machine.X += 10000000000000
		machine.Y += 10000000000000

		// log.Printf("%v", machine.X/machine.)

		// if minTokens != -1 {
		// 	log.Printf("Machine %d minimum tokens: %d", i, minTokens)
		// 	result += minTokens
		// } else {
		// 	log.Printf("No solution found for machine %d", i)
		// }
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

	// input := parseInput("input")
	// partOne := PartOneCountTokens(input)
	// log.Printf("Part One: %v", partOne)

	testPartTwo := PartTwoCountTokens(testInput)
	if testPartTwo != 480 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 480)
	}
}
