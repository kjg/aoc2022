package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	rows := strings.Split(stream, "\n")
	forest := rows[:len(rows)-1]

	Part1(forest)
	Part2(forest)
}

func Part1(forest []string) {
	visibleCount := 0
	visibleCount += len(forest[0]) * 2
	visibleCount += (len(forest) * 2) - 4

	for row := 1; row < len(forest)-1; row++ {
		for col := 1; col < len(forest[row])-1; col++ {
			visible := false
			currentHeight := treeHeight(forest[row][col])

			// check above
			visibleFromAbove := true
			for i := row - 1; i >= 0; i-- {

				if treeHeight(forest[i][col]) >= currentHeight {
					visibleFromAbove = false
					break
				}
			}
			visible = visibleFromAbove

			if !visible {
				// check below
				visibleFromBelow := true
				for i := row + 1; i < len(forest); i++ {
					if treeHeight(forest[i][col]) >= currentHeight {
						visibleFromBelow = false
						break
					}
				}
				visible = visibleFromBelow
			}

			if !visible {
				// check left
				visibleFromLeft := true
				for i := col - 1; i >= 0; i-- {
					if treeHeight(forest[row][i]) >= currentHeight {
						visibleFromLeft = false
						break
					}

				}
				visible = visibleFromLeft
			}

			if !visible {
				// check right
				visibleFromRight := true
				for i := col + 1; i < len(forest[row]); i++ {
					if treeHeight(forest[row][i]) >= currentHeight {
						visibleFromRight = false
						break
					}
				}
				visible = visibleFromRight
			}

			if visible {
				visibleCount++
			}
		}
	}

	fmt.Println("Trees visible from outside", visibleCount)
}

func Part2(forest []string) {
	bestScenicScore := 0

	for row := 1; row < len(forest)-1; row++ {
		for col := 1; col < len(forest[row])-1; col++ {
			currentHeight := treeHeight(forest[row][col])

			// check above
			visibilityAbove := row
			for i := row - 1; i >= 0; i-- {
				if treeHeight(forest[i][col]) >= currentHeight {
					visibilityAbove = row - i
					break
				}
			}
			fmt.Println(row, col, "visibilityAbove", visibilityAbove)

			// check below
			visibilityBelow := len(forest) - row - 1
			for i := row + 1; i < len(forest); i++ {
				if treeHeight(forest[i][col]) >= currentHeight {
					visibilityBelow = i - row
					break
				}
			}
			fmt.Println(row, col, "visibilityBelow", visibilityBelow)

			// check left
			visibilityLeft := col
			for i := col - 1; i >= 0; i-- {
				if treeHeight(forest[row][i]) >= currentHeight {
					visibilityLeft = col - i
					break
				}
			}
			fmt.Println(row, col, "visibilityLeft", visibilityLeft)

			// check right
			visibilityRight := len(forest[row]) - col - 1
			for i := col + 1; i < len(forest[row]); i++ {
				if treeHeight(forest[row][i]) >= currentHeight {
					visibilityRight = i - col
					break
				}
			}
			fmt.Println(row, col, "visibilityRight", visibilityRight)

			scenicScore := visibilityAbove * visibilityBelow * visibilityLeft * visibilityRight
			fmt.Println(row, col, "scenicScore", scenicScore)

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	fmt.Println("Best scenic score", bestScenicScore)
}

func treeHeight(char byte) int {
	height, _ := strconv.Atoi(string(char))
	return height
}
