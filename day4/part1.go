package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
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

		if (range1Start <= range2Start && range1End >= range2End) || (range1Start >= range2Start && range1End <= range2End) {

			overlaps++
		}
	}

	fmt.Println("There are", overlaps, "overlaps")
}

func getRange(rangeString string) (int, int) {
	rangePair := strings.Split(rangeString, "-")
	one, _ := strconv.Atoi(rangePair[0])
	two, _ := strconv.Atoi(rangePair[1])

	return one, two
}
