package main

import (
	"fmt"
	"os"
)

func Part1() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)
	markerAt := 0
	for i := 0; i < len(stream)-4; i++ {
		marker := stream[i : i+4]

		if IsUnique(marker) {
			markerAt = i + 4
			break
		}
	}

	fmt.Println("Part 1: ", markerAt)
}

func IsUnique(s string) bool {
	uniqMap := make(map[rune]bool)
	for _, c := range s {
		if uniqMap[c] {
			return false
		}
		uniqMap[c] = true
	}
	return true
}
