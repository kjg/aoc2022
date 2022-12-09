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

	fmt.Println(visibleCount)
}

func treeHeight(char byte) int {
	height, _ := strconv.Atoi(string(char))
	return height
}
