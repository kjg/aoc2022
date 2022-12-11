package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var strengthSum = 0

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	instructions := strings.Split(stream, "\n")
	instructions = instructions[0 : len(instructions)-1]
	screen := make([]string, 6)

	commandRegex := regexp.MustCompile(`^(\w+)\s?(.*)$`)

	cycle := 0
	xRegister := 1

	for _, instruction := range instructions {
		cycle = tick(cycle, xRegister, screen)
		parsedInstruction := commandRegex.FindStringSubmatch(instruction)
		cmd := parsedInstruction[1]

		switch cmd {
		case "noop":
		case "addx":
			cycle = tick(cycle, xRegister, screen)
			value, _ := strconv.Atoi(parsedInstruction[2])

			xRegister = xRegister + value
		}
	}
	fmt.Println("strengthSum", strengthSum)
	for _, line := range screen {
		fmt.Println(line)
	}
}

func tick(cycle int, xRegister int, screen []string) int {
	cycle = cycle + 1
	fmt.Println("starting cycle", cycle, "with register value", xRegister)
	if shouldDoInspection(cycle) {
		strength := cycle * xRegister
		strengthSum += strength
		fmt.Println("strenth", strength)
	}

	drawScreen(cycle, xRegister, screen)

	return cycle
}

func shouldDoInspection(cycle int) bool {
	return (cycle-20)%40 == 0
}

func drawScreen(cycle int, spritePosition int, screen []string) {
	colNumber := (cycle - 1) % 40
	rowNumber := (cycle - 1) / 40

	pixel := "."
	fmt.Println("cycle", cycle, "colNum", colNumber, "sprite", spritePosition)

	if colNumber >= spritePosition-1 && colNumber <= spritePosition+1 {
		pixel = "#"
	}
	screen[rowNumber] = screen[rowNumber] + pixel
}
