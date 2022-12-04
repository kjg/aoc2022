package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2() {

	var totalPriority = 0

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	// init slice of slices
	var groups [][]string
	// group lines into groups of 3
	for i := 0; i < len(lines)-1; i += 3 {
		groups = append(groups, lines[i:i+3])
	}

	for _, group := range groups {
		var commonItem rune

		for _, item := range group[2] {
			if strings.ContainsRune(group[0], item) && strings.ContainsRune(group[1], item) {
				commonItem = item
				break
			}
		}

		totalPriority += getPriority(commonItem)
	}

	fmt.Println("The total part 2 priority is:", totalPriority)
}
