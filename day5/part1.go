package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Part1() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numStacks, stacks := buildStacks(scanner)

	scanner.Scan() // Skip the blank line

	fmt.Println("Starting with stacks:", stacks)

	for scanner.Scan() {
		processInstruction(scanner.Text(), stacks)
	}

	fmt.Println("There", numStacks, "stacks containing", stacks)

	topCrates := ""
	for _, stack := range stacks {
		topCrates += stack[0]
	}

	fmt.Println("The top crates are:", topCrates)
}

func buildStacks(scanner *bufio.Scanner) (int, [][]string) {
	scanner.Scan()
	lineLength := len(scanner.Text()) + 1
	numStacks := lineLength / 4
	stacks := make([][]string, numStacks)

	processStackLine(scanner.Text(), stacks)

	for scanner.Scan() {
		if string(scanner.Text()[1]) == "1" {
			break
		}
		processStackLine(scanner.Text(), stacks)
	}

	return numStacks, stacks
}

func processStackLine(line string, stacks [][]string) {
	for i := 0; i < len(line)-2; i += 4 {
		stackNum := i / 4
		crate := string(line[i+1])
		if crate != " " {
			stacks[stackNum] = append(stacks[stackNum], crate)
		}
	}
}

func processInstruction(instruction string, stacks [][]string) {
	fmt.Println(instruction)
	matcher := regexp.MustCompile(`\d+`)
	matches := matcher.FindAllString(instruction, -1)
	numCrates, _ := strconv.Atoi(matches[0])
	fromStack, _ := strconv.Atoi(matches[1])
	toStack, _ := strconv.Atoi(matches[2])

	fromStack--
	toStack--

	for i := 0; i < numCrates; i++ {

		stacks[toStack] = append([]string{stacks[fromStack][0]}, stacks[toStack]...)
		stacks[fromStack] = stacks[fromStack][1:]
	}
	fmt.Println(stacks)
}
