package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	VelocityX int
	VelocityY int
}

type Square struct {
	X      int
	Y      int
	Robots []Robot
}

type Input struct {
	Squares [][]Square
	MaxX    int
	MaxY    int
}

func parseInput(path string, x, y int) *Input {

	content, err := os.ReadFile(path)
	if err != nil {
		return &Input{}
	}
	text := string(content)

	// Create an 11x7 grid of squares
	squares := make([][]Square, y)
	for i := range squares {
		squares[i] = make([]Square, x)
		// Initialize each square
		for j := range squares[i] {
			squares[i][j] = Square{
				X:      j,
				Y:      i,
				Robots: make([]Robot, 0),
			}
		}
	}

	// Parse the input text
	parts := strings.Fields(text)
	var currentX, currentY int

	for _, part := range parts {
		if strings.HasPrefix(part, "p=") {
			// Parse position
			pos := strings.TrimPrefix(part, "p=")
			coords := strings.Split(pos, ",")
			currentX, _ = strconv.Atoi(coords[0])
			currentY, _ = strconv.Atoi(coords[1])
		} else if strings.HasPrefix(part, "v=") {
			// Parse velocity and add robot to the corresponding square
			vel := strings.TrimPrefix(part, "v=")
			vels := strings.Split(vel, ",")
			vx, _ := strconv.Atoi(vels[0])
			vy, _ := strconv.Atoi(vels[1])

			// Add robot to the current position
			squares[currentY][currentX].Robots = append(squares[currentY][currentX].Robots, Robot{
				VelocityX: vx,
				VelocityY: vy,
			})
		}
	}

	return &Input{
		Squares: squares,
		MaxX:    x,
		MaxY:    y,
	}
}

func PartOneMoveRobots(input *Input, iterations int) int {

	for i := range iterations {
		log.Printf("Second %v", i)

		squares := make([][]Square, input.MaxY)
		for y := range squares {
			squares[y] = make([]Square, input.MaxX)
			for x := range input.MaxX {
				squares[y][x].X = x
				squares[y][x].Y = y

			}
		}

		for _, row := range input.Squares {
			for _, square := range row {
				for _, robot := range square.Robots {
					newX := square.X + robot.VelocityX
					if newX >= input.MaxX {
						newX -= input.MaxX
					} else if newX < 0 {
						newX += input.MaxX
					}
					newY := square.Y + robot.VelocityY
					if newY >= input.MaxY {
						newY -= input.MaxY
					} else if newY < 0 {
						newY += input.MaxY
					}
					squares[newY][newX].Robots = append(squares[newY][newX].Robots, robot)
				}

			}
		}
		input.Squares = squares

		for _, row := range input.Squares {
			for _, square := range row {
				if len(square.Robots) == 0 {
					print(".")
				} else {
					print(len(square.Robots))
				}
			}
			print("\n")
		}
	}
	// quadrant 1
	quad1 := 0
	quad1X := input.MaxX / 2
	quad1Y := input.MaxY / 2
	offsetX := quad1X + 1
	offsetY := quad1Y + 1
	for y := range quad1Y {
		for x := range quad1X {
			quad1 += len(input.Squares[y][x].Robots)
		}
	}
	quad2 := 0
	for y := range quad1Y {
		for x := range quad1X {
			quad2 += len(input.Squares[y][x+offsetX].Robots)
		}
	}

	quad3 := 0
	for y := range quad1Y {
		for x := range quad1X {
			quad3 += len(input.Squares[y+offsetY][x].Robots)
		}
	}
	quad4 := 0
	for y := range quad1Y {
		for x := range quad1X {
			quad4 += len(input.Squares[y+offsetY][x+offsetX].Robots)
		}
	}

	return quad1 * quad2 * quad3 * quad4
}

func main() {
	testInput := parseInput("sampleInput", 11, 7)
	testPartOne := PartOneMoveRobots(testInput, 100)
	if testPartOne != 12 {
		log.Fatalf("got wrong output for test part one, got %v want %v", testPartOne, 12)
	}

	input := parseInput("input", 101, 103)
	partOne := PartOneMoveRobots(input, 100)
	log.Printf("Part One: %v", partOne)

	input = parseInput("input", 101, 103)
	PartOneMoveRobots(input, 10000000)

	// testPartTwo := PartTwoCountTokens(testInput)
	// if testPartTwo != 480 {
	// 	log.Fatalf("got wrong output for test part two, got %v want %v", testPartTwo, 480)
	// }
}
