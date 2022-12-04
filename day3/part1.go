package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func Part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPriority = 0

	for scanner.Scan() {
		sack := scanner.Text()
		sackLength := len(sack)
		compartment1 := sack[0 : sackLength/2]
		compartment2 := sack[sackLength/2:]

		var commonItem rune

		for _, item := range compartment2 {
			if strings.ContainsRune(compartment1, item) {
				commonItem = item
				break
			}
		}

		totalPriority += getPriority(commonItem)
	}

	fmt.Println("The total part 1 priority is:", totalPriority)
}

func getPriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - int('a') + 1
	} else {
		return int(item) - int('A') + 1 + 26
	}
}
