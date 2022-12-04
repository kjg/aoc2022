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
		var itemPriority int

		for _, item := range compartment2 {
			if strings.ContainsRune(compartment1, item) {
				commonItem = item
				break
			}
		}
		if unicode.IsLower(commonItem) {
			itemPriority = int(commonItem) - int('a') + 1
		} else {
			itemPriority = int(commonItem) - int('A') + 1 + 26
		}

		totalPriority += itemPriority
	}

	fmt.Println("The total priority is:", totalPriority)
}
