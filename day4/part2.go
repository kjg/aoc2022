package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2() {
	var overlaps = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pairAssignment := scanner.Text()
		pairAssignments := strings.Split(pairAssignment, ",")

		range1Start, range1End := getRange(pairAssignments[0])
		range2Start, range2End := getRange(pairAssignments[1])

		if (range2Start <= range1End && range2End >= range1Start) || (range1Start <= range2End && range1End >= range2Start) {
			overlaps++
		}
	}

	fmt.Println("There are", overlaps, "partial overlaps")
}
