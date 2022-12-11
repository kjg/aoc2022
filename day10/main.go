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

	commandRegex := regexp.MustCompile(`^(\w+)\s?(.*)$`)

	cycle := 0
	xRegister := 1

	for _, instruction := range instructions {
		cycle = tick(cycle, xRegister)
		parsedInstruction := commandRegex.FindStringSubmatch(instruction)
		cmd := parsedInstruction[1]

		switch cmd {
		case "noop":
		case "addx":
			cycle = tick(cycle, xRegister)
			value, _ := strconv.Atoi(parsedInstruction[2])

			xRegister = xRegister + value
		}
	}
	fmt.Println("strengthSum", strengthSum)
}

func tick(cycle int, xRegister int) int {
	cycle = cycle + 1
	fmt.Println("starting cycle", cycle, "with register value", xRegister)
	if shouldDoInspection(cycle) {
		strength := cycle * xRegister
		strengthSum += strength
		fmt.Println("strenth", strength)
	}
	return cycle
}

func shouldDoInspection(cycle int) bool {
	return (cycle-20)%40 == 0
}
