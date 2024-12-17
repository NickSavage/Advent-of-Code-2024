package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Memory  []int
	IP      int
	Opcodes map[int]func(*Input)
	RegA    int
	RegB    int
	RegC    int
	Output  string
	Halt    bool
}

func parseInput(path string) *Input {

	input, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	result := &Input{
		IP:      0,
		Opcodes: make(map[int]func(*Input)),
		Memory:  make([]int, 0),
	}

	// Split the input into lines
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	// Parse registers
	for _, line := range lines {
		if strings.HasPrefix(line, "Register A:") {
			result.RegA, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register A:")))
		}
		if strings.HasPrefix(line, "Register B:") {
			result.RegB, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register B:")))
		}
		if strings.HasPrefix(line, "Register C:") {
			result.RegC, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register C:")))
		}
		if strings.HasPrefix(line, "Program:") {
			// Parse program memory
			programStr := strings.TrimSpace(strings.TrimPrefix(line, "Program:"))
			numbers := strings.Split(programStr, ",")
			for _, num := range numbers {
				n, _ := strconv.Atoi(strings.TrimSpace(num))
				result.Memory = append(result.Memory, n)
			}
		}
	}

	result.Opcodes[0] = adv
	result.Opcodes[1] = bxl
	result.Opcodes[2] = bst
	result.Opcodes[3] = jnz
	result.Opcodes[4] = bxc
	result.Opcodes[5] = out
	result.Opcodes[6] = bdv
	result.Opcodes[7] = cdv

	return result
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func adv(input *Input) {

	operand := input.Memory[input.IP+1]
	denom := -1
	if operand == 0 {
		denom = 0
	} else if operand == 1 {
		denom = 2
	} else if operand == 2 {
		denom = 4
	} else if operand == 3 {
		denom = 8
	} else if operand == 4 {
		denom = IntPow(2, input.RegA)
	} else if operand == 5 {
		denom = IntPow(2, input.RegB)
	} else if operand == 6 {
		denom = IntPow(2, input.RegC)
	} else {
		log.Fatalf("adv invalid input %v", operand)
	}
	if denom == -1 {
		log.Fatalf("adv invalid denom %v", operand)
	}

	result := input.RegA / denom
	input.RegA = result
}

func bxl(input *Input) {
	operand := input.Memory[input.IP+1]
	input.RegB = input.RegB ^ operand // bitwise XOR of RegB and literal operand
}

func bst(input *Input) {
	operand := input.Memory[input.IP+1]
	combo := -1
	// Use the same combo value determination as in other functions
	if operand == 0 {
		combo = 0
	} else if operand == 1 {
		combo = 2
	} else if operand == 2 {
		combo = 4
	} else if operand == 3 {
		combo = 8
	} else if operand == 4 {
		combo = input.RegA
	} else if operand == 5 {
		combo = input.RegB
	} else if operand == 6 {
		combo = input.RegC
	} else {
		log.Fatalf("bst invalid input %v", operand)
	}
	input.RegB = combo % 8 // store lowest 3 bits in RegB
}

func jnz(input *Input) {
	if input.RegA == 0 {
		return
	}
	operand := input.Memory[input.IP+1]
	input.IP = operand - 2
}

func bxc(input *Input) {
	input.RegB = input.RegB ^ input.RegC // bitwise XOR of RegB and RegC
}
func out(input *Input) {

	operand := input.Memory[input.IP+1]
	combo := -1
	if operand == 0 {
		combo = 0
	} else if operand == 1 {
		combo = 2
	} else if operand == 2 {
		combo = 4
	} else if operand == 3 {
		combo = 8
	} else if operand == 4 {
		combo = input.RegA
	} else if operand == 5 {
		combo = input.RegB
	} else if operand == 6 {
		combo = input.RegC
	} else {
		log.Fatalf("out invalid input %v", operand)
	}
	if input.Output == "" {
		input.Output = fmt.Sprintf("%v", combo%8)
	} else {
		input.Output += fmt.Sprintf(",%v", combo%8)

	}

}
func bdv(input *Input) {
	operand := input.Memory[input.IP+1]
	denom := -1
	if operand == 0 {
		denom = 0
	} else if operand == 1 {
		denom = 2
	} else if operand == 2 {
		denom = 4
	} else if operand == 3 {
		denom = 8
	} else if operand == 4 {
		denom = IntPow(2, input.RegA)
	} else if operand == 5 {
		denom = IntPow(2, input.RegB)
	} else if operand == 6 {
		denom = IntPow(2, input.RegC)
	} else {
		log.Fatalf("adv invalid input %v", operand)
	}
	if denom == -1 {
		log.Fatalf("adv invalid denom %v", operand)
	}

	result := input.RegA / denom
	input.RegB = result

}
func cdv(input *Input) {
	operand := input.Memory[input.IP+1]
	denom := -1
	if operand == 0 {
		denom = 0
	} else if operand == 1 {
		denom = 2
	} else if operand == 2 {
		denom = 4
	} else if operand == 3 {
		denom = 8
	} else if operand == 4 {
		denom = IntPow(2, input.RegA)
	} else if operand == 5 {
		denom = IntPow(2, input.RegB)
	} else if operand == 6 {
		denom = IntPow(2, input.RegC)
	} else {
		log.Fatalf("adv invalid input %v", operand)
	}
	if denom == -1 {
		log.Fatalf("adv invalid denom %v", operand)
	}

	result := input.RegA / denom
	input.RegC = result

}

func PartOneRunProgram(input *Input) string {
	input.Halt = false
	for {
		if input.IP >= len(input.Memory) {
			input.Halt = true
			// log.Printf("moved past end, halting")
			break
		}
		opcode := input.Memory[input.IP]
		var function func(*Input)
		var exists bool
		if function, exists = input.Opcodes[opcode]; !exists {
			log.Fatalf("tried to run opcode %v, no function exists", opcode)
		}
		// log.Printf("run %v", opcode)
		function(input)

		input.IP += 2

		if input.Halt {
			break
		}
	}
	return input.Output
}

func PartTwoRunProgram(input *Input) int {
	a := 117440
	expectedOutput := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input.Memory)), ","), "[]")
	for {
		log.Printf("check %v", a)

		input.RegA = a
		input.RegB = 0
		input.RegC = 0
		input.IP = 0

		input.Halt = false
		for {
			if input.IP >= len(input.Memory) {
				input.Halt = true
				// log.Printf("moved past end, halting")
				break
			}
			opcode := input.Memory[input.IP]
			var function func(*Input)
			var exists bool
			if function, exists = input.Opcodes[opcode]; !exists {
				log.Fatalf("tried to run opcode %v, no function exists", opcode)
			}
			// log.Printf("run %v", opcode)
			function(input)
			if len(input.Output) > len(expectedOutput) {
				break
			}

			input.IP += 2

			if input.Halt {
				break
			}
		}
		if input.Output == expectedOutput {
			log.Printf("found %v", a)
			break
		}
		a++

	}
	return a
}

func main() {
	testInput := parseInput("sampleInput")
	log.Printf("%v", testInput)
	testPartOne := PartOneRunProgram(testInput)
	log.Printf("test %v", testPartOne)
	if testPartOne != "4,6,3,5,6,3,5,2,1,0" {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 7036)
	}
	// testInput = parseInput("sampleInput2")
	// testPartOne = PartOneFindPaths(testInput)
	// if testPartOne != 11048 {
	// 	log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 11048)
	// }

	input := parseInput("input")
	partOne := PartOneRunProgram(input)
	log.Printf("Part One: %v", partOne)

	// testInput = parseInput("sampleInput")
	// testPartTwo := PartTwoFindOptimalPaths(testInput)
	// if testPartTwo != 45 {
	// 	log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 45)
	// }
	testInput = parseInput("sampleInput2")
	testPartTwo := PartTwoRunProgram(testInput)
	if testPartTwo != 117440 {
		log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 117400)
	}

	partTwo := PartTwoRunProgram(input)
	log.Printf("Part Two: %v", partTwo)

}
