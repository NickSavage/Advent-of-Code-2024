package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result   int
	Operands []int
}

type Input struct {
	Equations []Equation
}

func parseInput(path string) *Input {
	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}

	rows := strings.Split(string(content), "\n")
	equations := []Equation{}

	for _, row := range rows {
		if len(row) == 0 {
			continue // Skip empty lines
		}

		equation := Equation{}
		parts := strings.Split(row, ":")

		// Parse result
		result, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatalf("fell over parsing result: %v", err)
		}

		operands := []int{}
		operandParse := strings.Fields(parts[1])
		for _, operand := range operandParse {
			conv, err := strconv.Atoi(strings.TrimSpace(operand))
			if err != nil {
				log.Fatalf("fell over parsing operand: %v", err)
			}
			operands = append(operands, conv)
		}

		equation.Result = result
		equation.Operands = operands
		equations = append(equations, equation)
	}

	input := &Input{
		Equations: equations,
	}
	return input
}

func PartOneComputeOperands(input *Input) int {
	final := 0
	total := len(input.Equations)
	for x, equation := range input.Equations {
		log.Printf("equation %v/%v", x, total)
		numberOfOperations := len(equation.Operands) - 1
		possibleCombinations := 1 << numberOfOperations // This is same as 2^numberOfOperations
		for combo := range possibleCombinations {
			result := equation.Operands[0]
			binary := fmt.Sprintf("%0*b", numberOfOperations, combo)
			operations := strings.Split(binary, "")
			for i, operation := range operations {
				if operation == "0" {
					result += equation.Operands[i+1]
				} else {
					result = result * equation.Operands[i+1]
				}
				if result > equation.Result {
					// already too high
					continue
				}
			}
			if result == equation.Result {
				final += result
				break
			}
		}

	}
	return final
}

func PartTwoComputeOperands(input *Input) int {

	final := 0
	for _, equation := range input.Equations {
		intermediate := []int{equation.Operands[0]}
		for operand := range len(equation.Operands) - 1 {
			new := []int{}
			// try each operation

			for _, n := range intermediate {
				for i := range 3 {
					var result int
					if i == 0 {
						result = n + equation.Operands[operand+1]

					} else if i == 1 {
						result = n * equation.Operands[operand+1]

					} else {
						combined := fmt.Sprintf("%v%v", n, equation.Operands[operand+1])
						conv, _ := strconv.Atoi(combined)
						result = conv
					}
					if result <= equation.Result {
						new = append(new, result)
					}
				}
			}
			intermediate = new
		}
		for _, i := range intermediate {
			if i == equation.Result {
				final += i
				break
			}
		}
	}
	return final

}

func main() {
	testInput := parseInput("sampleInput")
	log.Printf("test %v", testInput)
	input := parseInput("input")

	testPartOne := PartOneComputeOperands(testInput)
	if testPartOne != 3749 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 3749)
	}
	testPartTwo := PartTwoComputeOperands(testInput)
	if testPartTwo != 11387 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 11387)
	}

	partOne := PartOneComputeOperands(input)
	log.Printf("Part One: %v", partOne)
	partTwo := PartTwoComputeOperands(input)
	log.Printf("Part Two: %v", partTwo)
}
