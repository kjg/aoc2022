package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Part2() {

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
		process9001Instruction(scanner.Text(), stacks)
	}

	fmt.Println("There", numStacks, "stacks containing", stacks)

	topCrates := ""
	for _, stack := range stacks {
		topCrates += stack[0]
	}

	fmt.Println("The top crates are:", topCrates)
}

func process9001Instruction(instruction string, stacks [][]string) {
	fmt.Println(instruction)
	matcher := regexp.MustCompile(`\d+`)
	matches := matcher.FindAllString(instruction, -1)
	numCrates, _ := strconv.Atoi(matches[0])
	fromStack, _ := strconv.Atoi(matches[1])
	toStack, _ := strconv.Atoi(matches[2])

	fromStack--
	toStack--

	for i := numCrates - 1; i >= 0; i-- {
		stacks[toStack] = append([]string{stacks[fromStack][i]}, stacks[toStack]...)
	}

	stacks[fromStack] = stacks[fromStack][numCrates:]

	fmt.Println(stacks)
}
