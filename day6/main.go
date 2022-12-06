package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	fmt.Println("Part 1: ", markerAfter(stream, 4))
	fmt.Println("Part 2: ", markerAfter(stream, 14))
}

func markerAfter(stream string, uniqueNeeded int) int {
	markerAt := 0
	for i := 0; i < len(stream)-uniqueNeeded; i++ {
		marker := stream[i : i+uniqueNeeded]

		if IsUnique(marker) {
			markerAt = i + uniqueNeeded
			break
		}
	}
	return markerAt
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
